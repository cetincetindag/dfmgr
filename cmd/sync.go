package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cetincetindag/dfmgr/pkg/config"
	"github.com/cetincetindag/dfmgr/pkg/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	overwriteExisting bool
	autoOrganize      bool
)

var syncCmd = &cobra.Command{
	Use:   "sync [file_paths...]",
	Short: "Sync files to your dotfiles repository",
	Long: `Add configuration files to your dotfiles repository.
Supports globbing patterns such as ~/.config/nvim/**/*.lua.
Can automatically organize files into appropriate categories.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runSyncCommand(args); err != nil {
			utils.Error("Failed to sync: %s", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)
	
	syncCmd.Flags().BoolVarP(&overwriteExisting, "force", "f", false, "Overwrite existing files")
	syncCmd.Flags().BoolVarP(&autoOrganize, "organize", "o", false, "Automatically organize files by category")
}

func runSyncCommand(paths []string) error {
	if len(paths) == 0 {
		return fmt.Errorf("no files specified")
	}
	
	localPath := config.CurrentConfig.LocalPath
	if !utils.IsGitRepo(localPath) {
		return fmt.Errorf("no dotfiles repository found at %s", localPath)
	}
	
	home := os.Getenv("HOME")
	
	// Expand any glob patterns
	expandedPaths := []string{}
	for _, pattern := range paths {
		if !strings.HasPrefix(pattern, "/") {
			pattern = filepath.Join(home, pattern)
		}
		
		matches, err := filepath.Glob(pattern)
		if err != nil {
			utils.Warning("Invalid pattern %s: %s", pattern, err)
			continue
		}
		
		if len(matches) == 0 {
			utils.Warning("No matches found for pattern: %s", pattern)
			continue
		}
		
		expandedPaths = append(expandedPaths, matches...)
	}
	
	if len(expandedPaths) == 0 {
		return fmt.Errorf("no files matched the specified patterns")
	}
	
	successCount := 0
	for _, path := range expandedPaths {
		relPath, err := filepath.Rel(home, path)
		if err != nil || strings.HasPrefix(relPath, "..") {
			utils.Warning("Skipping file outside of home directory: %s", path)
			continue
		}
		
		info, err := os.Stat(path)
		if err != nil {
			utils.Warning("Failed to stat file: %s", err)
			continue
		}
		
		// Determine the target directory
		targetDir := localPath
		if autoOrganize {
			fileInfo, found := config.GetConfigFileInfo(relPath)
			if found {
				if config.CurrentConfig.MultiOS {
					osFolder := config.GetOSFolder()
					targetDir = filepath.Join(localPath, osFolder, fileInfo.Category)
				} else {
					targetDir = filepath.Join(localPath, fileInfo.Category)
				}
			} else {
				basename := filepath.Base(relPath)
				category := promptCategory(basename)
				if config.CurrentConfig.MultiOS {
					osFolder := config.GetOSFolder()
					targetDir = filepath.Join(localPath, osFolder, category)
				} else {
					targetDir = filepath.Join(localPath, category)
				}
			}
		} else if config.CurrentConfig.MultiOS {
			osFolder := config.GetOSFolder()
			targetDir = filepath.Join(localPath, osFolder)
		}
		
		if err := utils.EnsureDirExists(targetDir); err != nil {
			utils.Warning("Failed to create directory: %s", err)
			continue
		}
		
		// Handle directories differently
		if info.IsDir() {
			if err := syncDirectory(path, targetDir, relPath); err != nil {
				utils.Warning("Failed to sync directory %s: %s", relPath, err)
				continue
			}
		} else {
			if err := syncFile(path, targetDir, relPath); err != nil {
				utils.Warning("Failed to sync file %s: %s", relPath, err)
				continue
			}
		}
		
		successCount++
	}
	
	if successCount > 0 {
		utils.Success("Successfully synced %d files/directories to your dotfiles repository", successCount)
		utils.Info("Remember to run 'dfmgr apply' to create symlinks for the new files")
	} else {
		return fmt.Errorf("failed to sync any files")
	}
	
	return nil
}

func syncFile(sourcePath, targetDir, relPath string) error {
	targetPath := filepath.Join(targetDir, filepath.Base(relPath))
	
	// Check if file already exists
	if _, err := os.Stat(targetPath); err == nil && !overwriteExisting {
		prompt := promptui.Prompt{
			Label:     fmt.Sprintf("File %s already exists in repository. Overwrite", filepath.Base(relPath)),
			IsConfirm: true,
		}
		
		result, err := prompt.Run()
		if err != nil || strings.ToLower(result) != "y" {
			utils.Info("Skipping file: %s", relPath)
			return nil
		}
	}
	
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}
	
	err = os.WriteFile(targetPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}
	
	utils.Success("Added file: %s", relPath)
	return nil
}

func syncDirectory(sourcePath, targetDir, relPath string) error {
	targetPath := filepath.Join(targetDir, filepath.Base(relPath))
	
	// Check if directory already exists
	if _, err := os.Stat(targetPath); err == nil && !overwriteExisting {
		prompt := promptui.Prompt{
			Label:     fmt.Sprintf("Directory %s already exists in repository. Merge", filepath.Base(relPath)),
			IsConfirm: true,
		}
		
		result, err := prompt.Run()
		if err != nil || strings.ToLower(result) != "y" {
			utils.Info("Skipping directory: %s", relPath)
			return nil
		}
	}
	
	if err := utils.EnsureDirExists(targetPath); err != nil {
		return fmt.Errorf("failed to create target directory: %w", err)
	}
	
	entries, err := os.ReadDir(sourcePath)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}
	
	for _, entry := range entries {
		entryPath := filepath.Join(sourcePath, entry.Name())
		entryRelPath := filepath.Join(relPath, entry.Name())
		
		info, err := entry.Info()
		if err != nil {
			utils.Warning("Failed to get info for %s: %s", entryRelPath, err)
			continue
		}
		
		if info.IsDir() {
			if err := syncDirectory(entryPath, targetPath, entryRelPath); err != nil {
				utils.Warning("Failed to sync subdirectory %s: %s", entryRelPath, err)
			}
		} else {
			if err := syncFile(entryPath, targetPath, entryRelPath); err != nil {
				utils.Warning("Failed to sync file %s: %s", entryRelPath, err)
			}
		}
	}
	
	utils.Success("Added directory: %s", relPath)
	return nil
}

func promptCategory(filename string) string {
	categories := config.ListCategories()
	
	prompt := promptui.Select{
		Label: fmt.Sprintf("Select category for %s", filename),
		Items: categories,
	}
	
	_, category, err := prompt.Run()
	if err != nil {
		return "Misc"
	}
	
	return category
} 
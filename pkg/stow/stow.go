package stow

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cetincetindag/dfmgr/pkg/config"
	"github.com/cetincetindag/dfmgr/pkg/utils"
)

func IsStowInstalled() bool {
	return utils.IsCommandAvailable("stow")
}

func StowPackages(sourcePath, targetPath string, packages []string) error {
	if !IsStowInstalled() {
		return fmt.Errorf("GNU stow is not installed")
	}

	utils.Info("Symlinking packages using stow: %s", strings.Join(packages, ", "))
	
	args := append([]string{
		"--verbose=1",
		"--target", targetPath,
		"--dir", sourcePath,
	}, packages...)
	
	cmd := exec.Command("stow", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func UnstowPackages(sourcePath, targetPath string, packages []string) error {
	if !IsStowInstalled() {
		return fmt.Errorf("GNU stow is not installed")
	}

	utils.Info("Removing symlinks using stow: %s", strings.Join(packages, ", "))
	
	args := append([]string{
		"--verbose=1",
		"--target", targetPath,
		"--dir", sourcePath,
		"--delete",
	}, packages...)
	
	cmd := exec.Command("stow", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func BackupAndRemoveConflicts(sourcePath, targetPath, backupDir string, packages []string) error {
	if err := utils.EnsureDirExists(backupDir); err != nil {
		return err
	}

	home := os.Getenv("HOME")
	for _, pkg := range packages {
		pkgPath := filepath.Join(sourcePath, pkg)
		
		err := filepath.Walk(pkgPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			
			if info.IsDir() {
				return nil
			}
			
			relPath, err := filepath.Rel(pkgPath, path)
			if err != nil {
				return err
			}
			
			targetFilePath := filepath.Join(targetPath, relPath)
			
			if _, err := os.Stat(targetFilePath); err == nil {
				relToHome, err := filepath.Rel(home, targetFilePath)
				if err != nil {
					relToHome = targetFilePath
				}
				
				utils.Warning("Found existing file: %s", relToHome)
				
				backupSubDir := filepath.Join(backupDir, pkg)
				backupPath, err := utils.BackupFile(targetFilePath, backupSubDir)
				if err != nil {
					return err
				}
				
				if backupPath != "" {
					utils.Info("Backed up to: %s", backupPath)
					if err := os.Remove(targetFilePath); err != nil {
						return err
					}
					utils.Info("Removed: %s", targetFilePath)
				}
			}
			
			return nil
		})
		
		if err != nil {
			return err
		}
	}
	
	return nil
}

func ApplyDotfiles(interactive bool) error {
	localPath := config.CurrentConfig.LocalPath
	home := os.Getenv("HOME")
	
	if !utils.IsGitRepo(localPath) {
		return fmt.Errorf("no dotfiles repository found at %s", localPath)
	}
	
	packages := []string{}
	
	entries, err := os.ReadDir(localPath)
	if err != nil {
		return err
	}
	
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() != ".git" {
			if config.CurrentConfig.MultiOS {
				osFolder := config.GetOSFolder()
				if osFolder != "" && entry.Name() == osFolder {
					osDir := filepath.Join(localPath, osFolder)
					osEntries, err := os.ReadDir(osDir)
					if err != nil {
						return err
					}
					
					for _, osEntry := range osEntries {
						if osEntry.IsDir() {
							packages = append(packages, filepath.Join(osFolder, osEntry.Name()))
						}
					}
				}
			} else {
				packages = append(packages, entry.Name())
			}
		}
	}
	
	if interactive && len(packages) > 0 {
		selectedPackages := []string{}
		
		utils.Info("Available packages:")
		for i, pkg := range packages {
			fmt.Printf("[%d] %s\n", i+1, pkg)
		}
		
		fmt.Print("Enter package numbers to apply (comma separated, or 'all' for all packages): ")
		var input string
		fmt.Scanln(&input)
		
		if strings.ToLower(input) == "all" {
			selectedPackages = packages
		} else {
			indices := strings.Split(input, ",")
			for _, idx := range indices {
				idx = strings.TrimSpace(idx)
				if i, err := fmt.Sscanf(idx, "%d", new(int)); err == nil && i > 0 && i <= len(packages) {
					selectedPackages = append(selectedPackages, packages[i-1])
				}
			}
		}
		
		packages = selectedPackages
	}
	
	if len(packages) == 0 {
		utils.Warning("No packages to apply")
		return nil
	}
	
	backupDir := filepath.Join(home, ".dfmgr_backup")
	if err := BackupAndRemoveConflicts(localPath, home, backupDir, packages); err != nil {
		return err
	}
	
	return StowPackages(localPath, home, packages)
} 
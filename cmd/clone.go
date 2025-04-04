package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cetincetindag/dfmgr/pkg/config"
	"github.com/cetincetindag/dfmgr/pkg/git"
	"github.com/cetincetindag/dfmgr/pkg/stow"
	"github.com/cetincetindag/dfmgr/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	selectiveFlag bool
)

var cloneCmd = &cobra.Command{
	Use:   "clone [username]",
	Short: "Clone a dotfiles repository",
	Long:  `Clone a dotfiles repository from GitHub and apply the configurations to your system.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		if err := runCloneCommand(username); err != nil {
			utils.Error("Failed to clone: %s", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(cloneCmd)
	cloneCmd.Flags().BoolVarP(&selectiveFlag, "selective", "s", false, "Selectively apply dotfiles")
}

func runCloneCommand(username string) error {
	if !utils.IsValidGitHubUsername(username) {
		return fmt.Errorf("invalid GitHub username format")
	}

	repo := "dotfiles"
	destPath := config.CurrentConfig.LocalPath

	if utils.IsGitRepo(destPath) {
		return fmt.Errorf("destination directory already contains a Git repository")
	}

	utils.Info("Cloning %s's dotfiles repository", username)

	if err := git.CloneRepo(username, repo, destPath); err != nil {
		return fmt.Errorf("failed to clone repository: %w", err)
	}

	utils.Success("Successfully cloned repository to %s", destPath)

	configPath := filepath.Join(destPath, ".dfmgr.json")
	if _, err := os.Stat(configPath); err == nil {
		utils.Info("Found dfmgr configuration file in repository")
		config.CurrentConfig.GithubUsername = username
		config.CurrentConfig.DotfilesRepo = repo
		config.CurrentConfig.LocalPath = destPath
		if err := config.SaveConfig(); err != nil {
			utils.Warning("Failed to save configuration: %s", err)
		}
	} else {
		utils.Warning("No dfmgr configuration file found in repository")
		config.CurrentConfig.GithubUsername = username
		config.CurrentConfig.DotfilesRepo = repo
		config.CurrentConfig.LocalPath = destPath
		if err := config.SaveConfig(); err != nil {
			utils.Warning("Failed to save configuration: %s", err)
		}
	}

	if err := stow.ApplyDotfiles(selectiveFlag); err != nil {
		return fmt.Errorf("failed to apply dotfiles: %w", err)
	}

	utils.Success("Successfully applied dotfiles from %s/%s", username, repo)
	return nil
} 
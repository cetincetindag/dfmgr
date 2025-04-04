package cmd

import (
	"fmt"
	"os"

	"github.com/cetincetindag/dfmgr/pkg/config"
	"github.com/cetincetindag/dfmgr/pkg/git"
	"github.com/cetincetindag/dfmgr/pkg/stow"
	"github.com/cetincetindag/dfmgr/pkg/utils"
	"github.com/spf13/cobra"
)

var forkCmd = &cobra.Command{
	Use:   "fork [username]",
	Short: "Fork a dotfiles repository",
	Long:  `Fork someone else's dotfiles repository and make it your own.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		username := args[0]
		if err := runForkCommand(username); err != nil {
			utils.Error("Failed to fork: %s", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(forkCmd)
}

func runForkCommand(username string) error {
	if !utils.IsValidGitHubUsername(username) {
		return fmt.Errorf("invalid GitHub username format")
	}

	repo := "dotfiles"
	destPath := config.CurrentConfig.LocalPath

	if utils.IsGitRepo(destPath) {
		return fmt.Errorf("destination directory already contains a Git repository")
	}

	utils.Info("Forking %s's dotfiles repository", username)

	if err := git.ForkRepo(username, repo); err != nil {
		return fmt.Errorf("failed to fork repository: %w", err)
	}

	utils.Success("Successfully forked repository from %s/%s", username, repo)

	myUsername := config.CurrentConfig.GithubUsername
	if myUsername == "" {
		utils.Warning("GitHub username not set in configuration. Please run 'dfmgr init' first.")
		return nil
	}

	if err := git.CloneRepo(myUsername, repo, destPath); err != nil {
		return fmt.Errorf("failed to clone forked repository: %w", err)
	}

	config.CurrentConfig.DotfilesRepo = repo
	config.CurrentConfig.LocalPath = destPath
	if err := config.SaveConfig(); err != nil {
		utils.Warning("Failed to save configuration: %s", err)
	}

	if err := stow.ApplyDotfiles(selectiveFlag); err != nil {
		return fmt.Errorf("failed to apply dotfiles: %w", err)
	}

	utils.Success("Successfully forked and applied dotfiles from %s/%s", username, repo)
	utils.Info("You can now customize the dotfiles and push your changes.")
	return nil
} 
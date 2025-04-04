package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/cetincetindag/dfmgr/pkg/config"
	"github.com/cetincetindag/dfmgr/pkg/git"
	"github.com/cetincetindag/dfmgr/pkg/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize dfmgr",
	Long:  `Set up dfmgr with your GitHub account and create your dotfiles repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runInitCommand(); err != nil {
			utils.Error("Failed to initialize: %s", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInitCommand() error {
	utils.Info("Starting dfmgr setup process...")

	if !utils.IsCommandAvailable("git") {
		return fmt.Errorf("git is not installed")
	}

	if !utils.IsCommandAvailable("gh") {
		utils.Warning("GitHub CLI (gh) is not installed. Some features may not work properly.")
		prompt := promptui.Prompt{
			Label:     "Continue without GitHub CLI",
			IsConfirm: true,
		}
		if _, err := prompt.Run(); err != nil {
			return fmt.Errorf("aborted setup")
		}
	}

	if !utils.IsCommandAvailable("stow") {
		utils.Warning("GNU stow is not installed. Required for symlinking dotfiles.")
		prompt := promptui.Prompt{
			Label:     "Continue without GNU stow",
			IsConfirm: true,
		}
		if _, err := prompt.Run(); err != nil {
			return fmt.Errorf("aborted setup")
		}
	}

	promptGithubUsername := promptui.Prompt{
		Label: "GitHub Username",
		Validate: func(input string) error {
			if input == "" {
				return fmt.Errorf("username cannot be empty")
			}
			fmt.Printf("Debug - Username entered: '%s'\n", input)
			return nil
		},
	}

	githubUsername, err := promptGithubUsername.Run()
	if err != nil {
		return fmt.Errorf("failed to get GitHub username: %w", err)
	}

	config.CurrentConfig.GithubUsername = githubUsername

	promptMultiOS := promptui.Prompt{
		Label:     "Use separate folders for different operating systems",
		IsConfirm: true,
	}

	multiOS, err := promptMultiOS.Run()
	if err == nil && multiOS == "y" {
		config.CurrentConfig.MultiOS = true
		utils.Info("Using separate folders for different operating systems")
	} else {
		config.CurrentConfig.MultiOS = false
		utils.Info("Using single directory structure")
	}

	promptRepoName := promptui.Prompt{
		Label:   "Repository Name",
		Default: "dotfiles",
	}

	repoName, err := promptRepoName.Run()
	if err != nil {
		return fmt.Errorf("failed to get repository name: %w", err)
	}

	config.CurrentConfig.DotfilesRepo = repoName

	promptLocalPath := promptui.Prompt{
		Label:   "Local Path",
		Default: filepath.Join(os.Getenv("HOME"), "dotfiles"),
	}

	localPath, err := promptLocalPath.Run()
	if err != nil {
		return fmt.Errorf("failed to get local path: %w", err)
	}

	config.CurrentConfig.LocalPath = localPath

	if err := utils.EnsureDirExists(localPath); err != nil {
		return fmt.Errorf("failed to create local directory: %w", err)
	}

	if config.CurrentConfig.MultiOS {
		osFolder := config.GetOSFolder()
		osFolderPath := filepath.Join(localPath, osFolder)
		if err := utils.EnsureDirExists(osFolderPath); err != nil {
			return fmt.Errorf("failed to create OS-specific directory: %w", err)
		}
		utils.Success("Created OS-specific directory: %s", osFolderPath)
	}

	if err := config.SaveConfig(); err != nil {
		return fmt.Errorf("failed to save configuration: %w", err)
	}

	utils.Success("Configuration saved to %s", config.ConfigFile())

	if err := git.SetupDefaultRepo(); err != nil {
		return fmt.Errorf("failed to set up GitHub repository: %w", err)
	}

	utils.Success("dfmgr initialized successfully!")
	utils.Info("Your dotfiles repository: https://github.com/%s/%s", 
		config.CurrentConfig.GithubUsername, 
		config.CurrentConfig.DotfilesRepo)
	utils.Info("Local path: %s", config.CurrentConfig.LocalPath)

	return nil
} 
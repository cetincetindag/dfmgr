package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/cetincetindag/dfmgr/pkg/config"
	"github.com/cetincetindag/dfmgr/pkg/git"
	"github.com/cetincetindag/dfmgr/pkg/utils"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var (
	commitMessage string
)

var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Push changes to the remote repository",
	Long:  `Add, commit, and push all changes in your dotfiles to the remote GitHub repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runPushCommand(); err != nil {
			utils.Error("Failed to push: %s", err)
			os.Exit(1)
		}
	},
}

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch changes from the remote repository",
	Long:  `Fetch and apply the latest changes from the remote dotfiles repository.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runFetchCommand(); err != nil {
			utils.Error("Failed to fetch: %s", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(pushCmd)
	rootCmd.AddCommand(fetchCmd)
	
	pushCmd.Flags().StringVarP(&commitMessage, "message", "m", "", "Commit message")
}

func runPushCommand() error {
	localPath := config.CurrentConfig.LocalPath
	
	if !utils.IsGitRepo(localPath) {
		return fmt.Errorf("no Git repository found at %s", localPath)
	}
	
	if err := git.AddFiles(localPath); err != nil {
		return fmt.Errorf("failed to add files: %w", err)
	}
	
	if commitMessage == "" {
		promptCommitMessage := promptui.Prompt{
			Label:   "Commit Message",
			Default: fmt.Sprintf("Update dotfiles - %s", time.Now().Format("2006-01-02")),
		}
		
		var err error
		commitMessage, err = promptCommitMessage.Run()
		if err != nil {
			return fmt.Errorf("failed to get commit message: %w", err)
		}
	}
	
	if err := git.Commit(localPath, commitMessage); err != nil {
		return fmt.Errorf("failed to commit changes: %w", err)
	}
	
	if err := git.Push(localPath); err != nil {
		return fmt.Errorf("failed to push changes: %w", err)
	}
	
	utils.Success("Successfully pushed changes to remote repository")
	return nil
}

func runFetchCommand() error {
	localPath := config.CurrentConfig.LocalPath
	
	if !utils.IsGitRepo(localPath) {
		return fmt.Errorf("no Git repository found at %s", localPath)
	}
	
	if err := git.Pull(localPath); err != nil {
		return fmt.Errorf("failed to pull changes: %w", err)
	}
	
	utils.Success("Successfully fetched latest changes from remote repository")
	return nil
} 
package cmd

import (
	"fmt"
	"os"

	"github.com/cetincetindag/dfmgr/pkg/stow"
	"github.com/cetincetindag/dfmgr/pkg/utils"
	"github.com/spf13/cobra"
)

var (
	applySelectiveFlag bool
)

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Apply dotfiles to the home directory",
	Long:  `Create symlinks for dotfiles in your repository to your home directory using GNU stow.`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := runApplyCommand(); err != nil {
			utils.Error("Failed to apply dotfiles: %s", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(applyCmd)
	applyCmd.Flags().BoolVarP(&applySelectiveFlag, "selective", "s", false, "Selectively apply dotfiles")
}

func runApplyCommand() error {
	utils.Info("Applying dotfiles to home directory...")
	
	if err := stow.ApplyDotfiles(applySelectiveFlag); err != nil {
		return fmt.Errorf("failed to apply dotfiles: %w", err)
	}
	
	utils.Success("Successfully applied dotfiles")
	return nil
} 
package cmd

import (
	"fmt"
	"os"

	"github.com/cetincetindag/dfmgr/pkg/config"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	cfgFile string
	rootCmd = &cobra.Command{
		Use:   "dfmgr",
		Short: "dfmgr - A dotfiles manager",
		Long: `dfmgr is a powerful dotfiles manager that helps you manage, share, and synchronize your configuration files across multiple machines.
It uses GNU stow for symlinking and integrates with GitHub for easy sharing and collaboration.`,
		Version: "0.2.0",
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dfmgr)")
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Only print logo for main commands, not for help or completion
		if cmd.Name() != "completion" && cmd.Name() != "help" {
			printLogo()
		}
	}
}

func initConfig() {
	config.LoadConfig(cfgFile)
}

func Execute() error {
	return rootCmd.Execute()
}

func printLogo() {
	logo := `
██████╗ ███████╗███╗   ███╗ ██████╗ ██████╗ 
██╔══██╗██╔════╝████╗ ████║██╔════╝ ██╔══██╗
██║  ██║█████╗  ██╔████╔██║██║  ███╗██████╔╝
██║  ██║██╔══╝  ██║╚██╔╝██║██║   ██║██╔══██╗
██████╔╝██║     ██║ ╚═╝ ██║╚██████╔╝██║  ██║
╚═════╝ ╚═╝     ╚═╝     ╚═╝ ╚═════╝ ╚═╝  ╚═╝
`
	fmt.Fprintf(os.Stderr, "%s\n", color.GreenString(logo))
	fmt.Fprintf(os.Stderr, "%s\n", color.CyanString("dotfiles manager - v0.2.0"))
} 
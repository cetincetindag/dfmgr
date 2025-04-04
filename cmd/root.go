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
		Version: "0.1.0",
	}
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.dfmgr)")
}

func initConfig() {
	config.LoadConfig(cfgFile)
}

func Execute() error {
	printLogo()
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
	fmt.Fprintf(os.Stderr, "%s\n", color.CyanString("dotfiles manager - v0.1.0"))
} 
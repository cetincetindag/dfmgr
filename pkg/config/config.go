package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/fatih/color"
)

type Config struct {
	GithubUsername string            `json:"github_username"`
	MultiOS        bool              `json:"multi_os"`
	OSSeparation   map[string]string `json:"os_separation"`
	DotfilesRepo   string            `json:"dotfiles_repo"`
	LocalPath      string            `json:"local_path"`
}

var (
	DefaultConfig = Config{
		GithubUsername: "",
		MultiOS:        false,
		OSSeparation:   make(map[string]string),
		DotfilesRepo:   "dotfiles",
		LocalPath:      filepath.Join(os.Getenv("HOME"), "dotfiles"),
	}

	CurrentConfig = DefaultConfig
)

func init() {
	CurrentConfig.OSSeparation = map[string]string{
		"darwin":  "macos",
		"linux":   "linux",
		"windows": "windows",
	}
}

func ConfigFile() string {
	if os.Getenv("DFMGR_CONFIG") != "" {
		return os.Getenv("DFMGR_CONFIG")
	}
	return filepath.Join(os.Getenv("HOME"), ".dfmgr")
}

func LoadConfig(cfgFile string) {
	if cfgFile == "" {
		cfgFile = ConfigFile()
	}

	_, err := os.Stat(cfgFile)
	if os.IsNotExist(err) {
		return
	}

	data, err := os.ReadFile(cfgFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s Unable to read config file: %s\n", color.RedString("[ERROR]"), err)
		return
	}

	err = json.Unmarshal(data, &CurrentConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s Unable to parse config file: %s\n", color.RedString("[ERROR]"), err)
		return
	}
}

func SaveConfig() error {
	data, err := json.MarshalIndent(CurrentConfig, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	cfgFile := ConfigFile()
	err = os.WriteFile(cfgFile, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
}

func GetCurrentOS() string {
	return runtime.GOOS
}

func GetOSFolder() string {
	if !CurrentConfig.MultiOS {
		return ""
	}
	
	os := GetCurrentOS()
	if folder, ok := CurrentConfig.OSSeparation[os]; ok {
		return folder
	}
	
	return os
} 
package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/cetincetindag/dfmgr/pkg/config"
	"github.com/cetincetindag/dfmgr/pkg/utils"
)

func CreateRepo(name string) error {
	utils.Info("Creating GitHub repository: %s", name)
	
	if !utils.IsCommandAvailable("gh") {
		return fmt.Errorf("GitHub CLI (gh) is not installed")
	}
	
	cmd := exec.Command("gh", "repo", "create", name, "--public", "--confirm")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func CloneRepo(username, repoName, destPath string) error {
	utils.Info("Cloning repository: %s/%s", username, repoName)
	
	repoURL := fmt.Sprintf("git@github.com:%s/%s.git", username, repoName)
	
	cmd := exec.Command("git", "clone", repoURL, destPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func ForkRepo(username, repoName string) error {
	utils.Info("Forking repository: %s/%s", username, repoName)
	
	if !utils.IsCommandAvailable("gh") {
		return fmt.Errorf("GitHub CLI (gh) is not installed")
	}
	
	repoFullName := fmt.Sprintf("%s/%s", username, repoName)
	cmd := exec.Command("gh", "repo", "fork", repoFullName, "--clone")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func InitRepo(path string) error {
	if utils.IsGitRepo(path) {
		return nil
	}

	utils.Info("Initializing Git repository at: %s", path)
	
	cmd := exec.Command("git", "init")
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func AddFiles(repoPath string) error {
	utils.Info("Adding files to Git repository")
	
	cmd := exec.Command("git", "add", ".")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func Commit(repoPath, message string) error {
	if message == "" {
		message = "Update dotfiles"
	}
	
	utils.Info("Committing changes: %s", message)
	
	cmd := exec.Command("git", "commit", "-m", message)
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func Push(repoPath string) error {
	utils.Info("Pushing changes to remote repository")
	
	cmd := exec.Command("git", "push", "origin", "HEAD")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func Pull(repoPath string) error {
	utils.Info("Pulling latest changes from remote repository")
	
	cmd := exec.Command("git", "pull")
	cmd.Dir = repoPath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	
	return cmd.Run()
}

func SetupDefaultRepo() error {
	localPath := config.CurrentConfig.LocalPath
	if err := utils.EnsureDirExists(localPath); err != nil {
		return err
	}
	
	if !utils.IsGitRepo(localPath) {
		if err := InitRepo(localPath); err != nil {
			return err
		}
		
		readmePath := filepath.Join(localPath, "README.md")
		readmeContent := GenerateReadme()
		
		if err := os.WriteFile(readmePath, []byte(readmeContent), 0644); err != nil {
			return err
		}
		
		if err := AddFiles(localPath); err != nil {
			return err
		}
		
		if err := Commit(localPath, "Initial commit"); err != nil {
			return err
		}
		
		username := config.CurrentConfig.GithubUsername
		repoName := config.CurrentConfig.DotfilesRepo
		
		if err := CreateRepo(repoName); err != nil {
			return err
		}
		
		remoteURL := fmt.Sprintf("git@github.com:%s/%s.git", username, repoName)
		cmd := exec.Command("git", "remote", "add", "origin", remoteURL)
		cmd.Dir = localPath
		if err := cmd.Run(); err != nil {
			return err
		}
		
		if err := Push(localPath); err != nil {
			return err
		}
	}
	
	return nil
}

func GenerateReadme() string {
	username := config.CurrentConfig.GithubUsername
	currentOS := config.GetCurrentOS()
	
	content := []string{
		"# Dotfiles",
		"",
		fmt.Sprintf("This repository contains my personal dotfiles managed using [dfmgr](https://github.com/%s/dfmgr).", username),
		"",
		"## Contents",
		"",
	}
	
	if config.CurrentConfig.MultiOS {
		content = append(content, fmt.Sprintf("These dotfiles are organized by operating system. Currently tracking configurations for **%s**.", strings.Title(currentOS)))
	} else {
		content = append(content, "This repository contains configuration files for various applications and tools.")
	}
	
	content = append(content,
		"",
		"## Installation",
		"",
		"To use these dotfiles on your system:",
		"",
		"1. Install dfmgr:",
		"```bash",
		"go install github.com/cetince/dfmgr@latest",
		"```",
		"",
		"2. Clone and apply the dotfiles:",
		"```bash",
		fmt.Sprintf("dfmgr clone %s", username),
		"```",
		"",
		"## License",
		"",
		"These dotfiles are provided as-is under the MIT License.",
	)
	
	return strings.Join(content, "\n")
} 
package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

var (
	InfoColor    = color.New(color.FgCyan).SprintFunc()
	SuccessColor = color.New(color.FgGreen).SprintFunc()
	WarningColor = color.New(color.FgYellow).SprintFunc()
	ErrorColor   = color.New(color.FgRed).SprintFunc()
)

func Info(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s %s\n", InfoColor("[INFO]"), fmt.Sprintf(format, args...))
}

func Success(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s %s\n", SuccessColor("[SUCCESS]"), fmt.Sprintf(format, args...))
}

func Warning(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s %s\n", WarningColor("[WARNING]"), fmt.Sprintf(format, args...))
}

func Error(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "%s %s\n", ErrorColor("[ERROR]"), fmt.Sprintf(format, args...))
}

func ExecuteCommand(name string, args ...string) (string, error) {
	cmd := exec.Command(name, args...)
	output, err := cmd.CombinedOutput()
	return string(output), err
}

func IsCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

func EnsureDirExists(path string) error {
	return os.MkdirAll(path, 0755)
}

func IsGitRepo(path string) bool {
	_, err := os.Stat(filepath.Join(path, ".git"))
	return err == nil
}

func FindConfigFiles(patterns []string) (map[string]string, error) {
	results := make(map[string]string)
	home := os.Getenv("HOME")

	for _, pattern := range patterns {
		if !strings.HasPrefix(pattern, "/") {
			pattern = filepath.Join(home, pattern)
		}
		
		matches, err := filepath.Glob(pattern)
		if err != nil {
			return nil, err
		}
		
		for _, match := range matches {
			relPath, err := filepath.Rel(home, match)
			if err != nil {
				continue
			}
			results[match] = relPath
		}
	}

	return results, nil
}

func BackupFile(filePath string, backupDir string) (string, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", nil
	}

	fileName := filepath.Base(filePath)
	backupPath := filepath.Join(backupDir, fileName)
	
	err := EnsureDirExists(backupDir)
	if err != nil {
		return "", err
	}
	
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	
	err = os.WriteFile(backupPath, data, 0644)
	if err != nil {
		return "", err
	}
	
	return backupPath, nil
}

func IsValidGitHubUsername(username string) bool {
	pattern := `^[a-z\d](?:[a-z\d]|-(?=[a-z\d])){0,38}$`
	matched, _ := regexp.MatchString(pattern, strings.ToLower(username))
	return matched
}

func FileExtension(filename string) string {
	ext := filepath.Ext(filename)
	if ext == "" {
		return ""
	}
	return ext[1:]
} 
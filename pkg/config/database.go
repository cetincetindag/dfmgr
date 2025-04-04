package config

type ConfigFileInfo struct {
	Name        string
	Description string
	Category    string
}

var ConfigFileDatabase = map[string]ConfigFileInfo{
	// Shell
	".zshrc":     {Name: "Zsh Config", Description: "Configuration for the Z shell", Category: "Shell"},
	".bashrc":    {Name: "Bash Config", Description: "Configuration for Bash shell", Category: "Shell"},
	".profile":   {Name: "Shell Profile", Description: "Environment variables and shell startup commands", Category: "Shell"},
	".aliases":   {Name: "Shell Aliases", Description: "Custom command aliases for shell", Category: "Shell"},
	".functions": {Name: "Shell Functions", Description: "Custom shell functions", Category: "Shell"},
	
	// Terminal
	".tmux.conf":      {Name: "Tmux Config", Description: "Configuration for the Tmux terminal multiplexer", Category: "Terminal"},
	".alacritty.yml":  {Name: "Alacritty Config", Description: "Configuration for the Alacritty terminal emulator", Category: "Terminal"},
	".kitty.conf":     {Name: "Kitty Config", Description: "Configuration for the Kitty terminal emulator", Category: "Terminal"},
	".terminfo":       {Name: "Terminal Info", Description: "Terminal capability database", Category: "Terminal"},
	".screenrc":       {Name: "Screen Config", Description: "Configuration for GNU Screen", Category: "Terminal"},
	".hyper.js":       {Name: "Hyper Config", Description: "Configuration for the Hyper terminal", Category: "Terminal"},
	".wezterm.lua":    {Name: "WezTerm Config", Description: "Configuration for the WezTerm terminal", Category: "Terminal"},
	
	// Editors
	".vimrc":         {Name: "Vim Config", Description: "Configuration for Vim editor", Category: "Editor"},
	".vim":           {Name: "Vim Directory", Description: "Directory containing Vim plugins and settings", Category: "Editor"},
	".emacs":         {Name: "Emacs Config", Description: "Configuration for Emacs editor", Category: "Editor"},
	".emacs.d":       {Name: "Emacs Directory", Description: "Directory containing Emacs packages and settings", Category: "Editor"},
	".spacemacs":     {Name: "Spacemacs Config", Description: "Configuration for Spacemacs", Category: "Editor"},
	".ideavimrc":     {Name: "IdeaVim Config", Description: "Vim emulation for IntelliJ IDEA", Category: "Editor"},
	".nanorc":        {Name: "Nano Config", Description: "Configuration for the Nano editor", Category: "Editor"},
	".config/nvim":   {Name: "Neovim Config", Description: "Configuration for Neovim", Category: "Editor"},
	
	// Version Control
	".gitconfig":    {Name: "Git Config", Description: "Global Git configuration", Category: "Version Control"},
	".gitignore":    {Name: "Git Ignore", Description: "Global Git ignore patterns", Category: "Version Control"},
	".gitattributes": {Name: "Git Attributes", Description: "Attributes for Git repositories", Category: "Version Control"},
	".hgrc":         {Name: "Mercurial Config", Description: "Configuration for Mercurial", Category: "Version Control"},
	
	// Window Managers
	".xinitrc":      {Name: "X Init", Description: "X Window System initialization", Category: "Window Manager"},
	".Xresources":   {Name: "X Resources", Description: "X Window System resources", Category: "Window Manager"},
	".i3":           {Name: "i3 Config", Description: "Configuration for the i3 window manager", Category: "Window Manager"},
	".config/i3":    {Name: "i3 Config Directory", Description: "Directory containing i3 window manager configuration", Category: "Window Manager"},
	".xmonad":       {Name: "XMonad Config", Description: "Configuration for the XMonad window manager", Category: "Window Manager"},
	".dwm":          {Name: "DWM Config", Description: "Configuration for the Dynamic Window Manager", Category: "Window Manager"},
	".awesomewm":    {Name: "AwesomeWM Config", Description: "Configuration for the Awesome window manager", Category: "Window Manager"},
	".config/sway":  {Name: "Sway Config", Description: "Configuration for the Sway window manager", Category: "Window Manager"},
	
	// Desktop Environment
	".config/gtk-3.0": {Name: "GTK3 Config", Description: "GTK3 configuration", Category: "Desktop"},
	".gtkrc-2.0":     {Name: "GTK2 Config", Description: "GTK2 configuration", Category: "Desktop"},
	".config/picom":  {Name: "Picom Config", Description: "Configuration for the Picom compositor", Category: "Desktop"},
	".config/compton.conf": {Name: "Compton Config", Description: "Configuration for the Compton compositor", Category: "Desktop"},
	".config/rofi":   {Name: "Rofi Config", Description: "Configuration for the Rofi application launcher", Category: "Desktop"},
	".config/polybar": {Name: "Polybar Config", Description: "Configuration for the Polybar status bar", Category: "Desktop"},
	".config/dunst":  {Name: "Dunst Config", Description: "Configuration for the Dunst notification daemon", Category: "Desktop"},
	
	// Shells and Prompts
	".config/starship.toml": {Name: "Starship Config", Description: "Configuration for the Starship prompt", Category: "Shell"},
	".p10k.zsh":     {Name: "Powerlevel10k Config", Description: "Configuration for the Powerlevel10k prompt", Category: "Shell"},
	".zsh":          {Name: "Zsh Directory", Description: "Directory containing Zsh plugins and settings", Category: "Shell"},
	".oh-my-zsh":    {Name: "Oh-My-Zsh", Description: "Oh-My-Zsh configuration and plugins", Category: "Shell"},
	
	// Package Managers
	".npmrc":        {Name: "NPM Config", Description: "Configuration for NPM", Category: "Development"},
	".yarnrc":       {Name: "Yarn Config", Description: "Configuration for Yarn", Category: "Development"},
	".cargo/config": {Name: "Cargo Config", Description: "Configuration for Rust's Cargo", Category: "Development"},
	".pip/pip.conf": {Name: "Pip Config", Description: "Configuration for Python's Pip", Category: "Development"},
	
	// Programming Languages
	".pylintrc":     {Name: "Pylint Config", Description: "Configuration for Python linting", Category: "Development"},
	".config/pycodestyle": {Name: "PEP8 Config", Description: "Configuration for Python code style", Category: "Development"},
	".eslintrc":     {Name: "ESLint Config", Description: "Configuration for JavaScript/TypeScript linting", Category: "Development"},
	".prettierrc":   {Name: "Prettier Config", Description: "Configuration for code formatting", Category: "Development"},
	".editorconfig": {Name: "Editor Config", Description: "Consistent coding styles across editors", Category: "Development"},
	
	// SSH and Security
	".ssh/config":   {Name: "SSH Config", Description: "SSH client configuration", Category: "Security"},
	".gnupg":        {Name: "GnuPG Directory", Description: "GnuPG keys and configuration", Category: "Security"},
	
	// Email
	".mutt":         {Name: "Mutt Config", Description: "Configuration for the Mutt email client", Category: "Communication"},
	".muttrc":       {Name: "Mutt Config File", Description: "Configuration file for the Mutt email client", Category: "Communication"},
	".mbsyncrc":     {Name: "mbsync Config", Description: "Configuration for the mbsync mail syncing tool", Category: "Communication"},
	
	// Music and Media
	".config/mpd":   {Name: "MPD Config", Description: "Configuration for the Music Player Daemon", Category: "Media"},
	".ncmpcpp":      {Name: "ncmpcpp Config", Description: "Configuration for the ncmpcpp music player", Category: "Media"},
	".config/mpv":   {Name: "MPV Config", Description: "Configuration for the MPV media player", Category: "Media"},
	
	// File Managers
	".config/ranger": {Name: "Ranger Config", Description: "Configuration for the Ranger file manager", Category: "Utilities"},
	".config/lf":    {Name: "LF Config", Description: "Configuration for the LF file manager", Category: "Utilities"},
	
	// Browsers and Web
	".config/qutebrowser": {Name: "Qutebrowser Config", Description: "Configuration for the Qutebrowser", Category: "Web"},
	
	// System
	".config/systemd/user": {Name: "Systemd User Units", Description: "User services for systemd", Category: "System"},
	
	// macOS Specific
	".config/karabiner": {Name: "Karabiner Config", Description: "Key remapping for macOS", Category: "macOS"},
	".skhdrc":        {Name: "skhd Config", Description: "Simple hotkey daemon for macOS", Category: "macOS"},
	".yabairc":       {Name: "Yabai Config", Description: "Window manager for macOS", Category: "macOS"},
	
	// Other
	".config":       {Name: "Config Directory", Description: "XDG config home directory", Category: "System"},
	".local/share":  {Name: "Data Directory", Description: "XDG data home directory", Category: "System"},
	".tmuxp":        {Name: "Tmuxp Sessions", Description: "Tmux session manager configurations", Category: "Terminal"},
	".config/bat":   {Name: "Bat Config", Description: "Configuration for the Bat command", Category: "Utilities"},
	".config/lazygit": {Name: "Lazygit Config", Description: "Configuration for the Lazygit TUI", Category: "Development"},
}

func GetConfigFileInfo(filename string) (ConfigFileInfo, bool) {
	info, found := ConfigFileDatabase[filename]
	return info, found
}

func GetConfigFilesInCategory(category string) []ConfigFileInfo {
	var result []ConfigFileInfo
	
	for _, info := range ConfigFileDatabase {
		if info.Category == category {
			result = append(result, info)
		}
	}
	
	return result
}

func ListCategories() []string {
	categories := make(map[string]bool)
	
	for _, info := range ConfigFileDatabase {
		categories[info.Category] = true
	}
	
	var result []string
	for category := range categories {
		result = append(result, category)
	}
	
	return result
} 
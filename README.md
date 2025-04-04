# dfmgr - Dotfiles Manager

A powerful, easy-to-use dotfiles manager written in Go. Simplify the management, sharing, and synchronization of your configuration files across multiple machines.

## Features

- Colorful and well-styled command line interface
- Interactive setup process
- Multi-OS support to manage dotfiles for different operating systems
- Automatic OS detection
- GitHub integration for easy sharing and collaboration
- GNU stow integration for symlinking
- Backup and restore functionality
- Extensive dotfiles database

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (1.16 or later)
- [Git](https://git-scm.com/downloads)
- [GNU Stow](https://www.gnu.org/software/stow/)
- [GitHub CLI](https://cli.github.com/) (optional, for GitHub integration)

### Installing from Source

```bash
go install github.com/cetince/dfmgr@latest
```

## Getting Started

### Initialize dfmgr

Run the setup process to configure dfmgr:

```bash
dfmgr init
```

This will guide you through:
- Setting up your GitHub username
- Configuring multi-OS support if needed
- Creating a local dotfiles directory
- Creating a GitHub repository for your dotfiles

### Clone Existing Dotfiles

To clone and apply someone else's dotfiles:

```bash
dfmgr clone username
```

Use the `-s` or `--selective` flag to interactively choose which configurations to apply:

```bash
dfmgr clone -s username
```

### Fork Dotfiles

To fork someone else's dotfiles repository and make it your own:

```bash
dfmgr fork username
```

### Manage Your Dotfiles

Push your changes to GitHub:

```bash
dfmgr push -m "Update vim configuration"
```

Fetch the latest changes from GitHub:

```bash
dfmgr fetch
```

## Command Reference

| Command | Description |
|---------|-------------|
| `dfmgr init` | Initialize dfmgr and set up your dotfiles repository |
| `dfmgr clone [username]` | Clone a dotfiles repository and apply configurations |
| `dfmgr clone -s [username]` | Clone a repository and selectively apply configurations |
| `dfmgr fork [username]` | Fork someone else's dotfiles repository |
| `dfmgr push` | Add, commit, and push changes to your dotfiles repository |
| `dfmgr fetch` | Pull the latest changes from your dotfiles repository |
| `dfmgr sync [file_paths...]` | Add configuration files to your dotfiles repository |
| `dfmgr sync -o [file_paths...]` | Add and automatically organize files by category |

## FAQ

### Why use dfmgr instead of other dotfiles managers?

dfmgr combines the power of Git for version control, GNU stow for symlink management, and GitHub for sharing in one easy-to-use tool. It also provides features like OS-specific configurations, interactive setup, and automatic backups.

### Do I need a GitHub account to use dfmgr?

While dfmgr works best with GitHub integration, you can use it without a GitHub account for local dotfiles management. However, you'll miss out on the sharing and synchronization features.

### How does dfmgr handle conflicts with existing dotfiles?

When applying dotfiles that would conflict with existing ones, dfmgr will prompt you to backup the existing files before replacing them. Backups are stored in `~/.dfmgr_backup/`.

### Can I manage dotfiles for multiple operating systems?

Yes! dfmgr allows you to organize your dotfiles in OS-specific directories (e.g., `dotfiles/macos`, `dotfiles/linux`) and will automatically detect your current OS.

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 
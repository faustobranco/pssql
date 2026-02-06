# pssql

[![Go Version](https://img.shields.io/github/go-mod/go-version/youruser/pssql)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**pssql** is a lightweight CLI connection manager for PostgreSQL. It allows you to quickly switch between multiple database servers defined in a JSON configuration, supporting both interactive selection and direct connection flags.

[Image of CLI command line interface with database connections]

## Features

* **Interactive Menu**: Beautiful fuzzy-search selection powered by `promptui`.
* **Direct Connect**: Use `--connect <name>` for instant access.
* **Smart Defaults**: Automatic fallback for ports (5432) and CLI tools (pgcli).
* **Shell Autocomplete**: Full support for Zsh and Bash completions (including server names).
* **Environment Aware**: Prompts for missing credentials (User/Database) if not defined in config.

## Installation

### Using Homebrew (Recommended)

```bash
brew tap youruser/tap
brew install pssql
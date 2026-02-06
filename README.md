# pssql

[](https://github.com/faustobranco/pssql#pssql)

[![Go Version](https://img.shields.io/github/go-mod/go-version/youruser/pssql)](https://golang.org)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
**pssql**  is a lightweight CLI connection manager for PostgreSQL. It allows you to quickly switch between multiple database servers defined in a JSON configuration, supporting both interactive selection and direct connection flags.

[Image of CLI command line interface with database connections]

## Features

[](https://github.com/faustobranco/pssql#features)

-   **Interactive Menu**: Beautiful fuzzy-search selection powered by  `promptui`.
-   **Direct Connect**: Use  `--connect <name>`  for instant access.
-   **Smart Defaults**: Automatic fallback for ports (5432) and CLI tools (pgcli).
-   **Shell Autocomplete**: Full support for Zsh and Bash completions (including server names).
-   **Environment Aware**: Prompts for missing credentials (User/Database) if not defined in config.

## Installation

[](https://github.com/faustobranco/pssql#installation)

### Using Homebrew (Recommended)

[](https://github.com/faustobranco/pssql#using-homebrew-recommended)

```
brew tap youruser/tap
brew install pssql
```

### Manual Installation

1.  Ensure you have [Go](https://golang.org/doc/install) installed.
    
2.  Clone the repository and build:
    

Bash

```
go build -o pssql .
sudo mv pssql /usr/local/bin/
```

## Configuration

The tool expects a JSON file at `~/.pssql/pssql.json`.

JSON

```
{
	"postgresql": [
					{
					"name": "Postgres Prod",
					"host": "myprodpostgresql.eu-central-1.rds.amazonaws.com",
					"port": 5432,
					"database": "postgres",
					"user": "admin",
					"cli": "pgcli"
					},
					{
					"name": "Postgres Dev",
					"host": "mydevpostgresql.eu-central-1.rds.amazonaws.com",
					"port": 5432,
					"database": "postgres",
					"user": "admin",
					"cli": "pgcli"
					}
			]
}
```

## Usage

Command

Description

`pssql`

Open interactive menu

`pssql --connect "Name"`

Connect directly to a specific server

`pssql --list`

List all configured servers

`pssql --config ./alt.json`

Use a different config file

|Command| Description |
|--|--|
|pssql |Open interactive menu|
|pssql --connect "Name"  |Connect directly to a specific server|
|pssql --list |List all configured servers|
|pssql --config ./alt.json |Use a different config file|


## Shell Completion

If installed via Homebrew, completions are handled automatically. For manual setup, refer to the `completion/` directory.

## Requirements

-   [pgcli](https://www.pgcli.com/) (recommended) or `psql`.
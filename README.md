
# pssql

  

[](https://github.com/faustobranco/pssql#pssql)

  

[![Go Report Card](https://goreportcard.com/badge/github.com/faustobranco/pssql)](https://goreportcard.com/report/github.com/faustobranco/pssql)

[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

![GitHub License](https://img.shields.io/github/license/faustobranco/pssql)

  

![Views](https://hits.dwyl.com/faustobranco/pssql.svg)

![GitHub release (latest by date)](https://img.shields.io/github/v/release/faustobranco/pssql)

![Homebrew](https://img.shields.io/badge/homebrew-tap-orange)


### Help:
<img width="740" height="189" alt="Screenshot 2026-02-09 at 15 22 59" src="https://github.com/user-attachments/assets/e8408bb3-481b-4b3e-b527-42e447b5521e" />


### List:
<img width="757" height="85" alt="Screenshot 2026-02-09 at 15 23 08" src="https://github.com/user-attachments/assets/a08d3ce0-18ff-4341-984a-b013a5af3bd5" />


### Direct Connect
<img width="724" height="168" alt="Screenshot 2026-02-09 at 15 24 38" src="https://github.com/user-attachments/assets/a18951bd-6781-479d-8f32-e6d34941fe9f" />


### Select / List Connection 
<img width="1519" height="123" alt="Screenshot 2026-02-09 at 15 25 01" src="https://github.com/user-attachments/assets/bb741c6f-0c05-4377-8bbb-d343e9f6cf37" />


### Connected
<img width="732" height="146" alt="Screenshot 2026-02-09 at 15 25 19" src="https://github.com/user-attachments/assets/381e981a-eff7-4bb2-a76e-f999328ef4cc" />


### AWS/RDS IAM connect
<img width="721" height="117" alt="Screenshot 2026-02-09 at 15 29 41" src="https://github.com/user-attachments/assets/f39ac709-650d-4b2a-8d48-da5347a8853a" />

  

**pssql** is a lightweight CLI connection manager for PostgreSQL. It allows you to quickly switch between multiple database servers defined in a JSON configuration, supporting both interactive selection and direct connection flags.

  



  

## Features

  

[](https://github.com/faustobranco/pssql#features)

  

-  **Interactive Menu**: Beautiful fuzzy-search selection powered by `promptui`.

-  **Direct Connect**: Use `--connect <name>` for instant access.

-  **Smart Defaults**: Automatic fallback for ports (5432) and CLI tools (pgcli).

-  **Shell Autocomplete**: Full support for Zsh and Bash completions (including server names).

-  **Environment Aware**: Prompts for missing credentials (User/Database) if not defined in config.

  

## Installation

  

[](https://github.com/faustobranco/pssql#installation)

  

### Using Homebrew (Recommended)

  

[](https://github.com/faustobranco/pssql#using-homebrew-recommended)

  

```
brew tap youruser/tap
brew install pssql
```

  

### Manual Installation

  

1. Ensure you have [Go](https://golang.org/doc/install) installed.

2. Clone the repository and build:

  

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
		},
		{
			"name": "Postgres Test - AWS",
			"host": "mytestpostgresql.jio43u089fge.eu-central-1.rds.amazonaws.com",
			"port": 5432,
			"database": "postgres",
			"user": "your.user",
			"auth": "aws-iam",
			"aws-iam": {
				"region": "eu-central-1",
				"profile": "test"
			},
			"cli": "pgcli"
		}
	]
}

```

  

## Usage

  
  

|Command| Description |
|--|--|
|pssql |Open interactive menu|
|pssql --connect "Name" |Connect directly to a specific server|
|pssql --list |List all configured servers|
|pssql --config ./alt.json |Use a different config file|

  
  

## Shell Completion

  

If installed via Homebrew, completions are handled automatically. For manual setup, refer to the `completion/` directory.

  

## Requirements

  

- [pgcli](https://www.pgcli.com/) (recommended) or `psql`.

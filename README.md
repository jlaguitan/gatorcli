# Gator

A command-line RSS feed aggregator written in Go.

## Prerequisites

- Go
- PostgreSQL

## Installation

Install gator to your Go bin path:

```
go install github.com/jlaguitan/gator@latest
```

## Configuration

Create a config file in your home directory named `.gatorconfig.json` and enter the following:

```json
{
    "db_url": "postgres://username:password@localhost:5432/gator",
    "current_user_name": "yourusername"
}
```
*Replace db_url value with your PostgreSQL connection string



## Setting up the database schema:
1. Install Goose: 
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
2. Navigate to the `sql/schema` directory.
3. Run.
```bash
goose postgres <connection_string> up
```

## Commands

After building the project, you can run commands in this format:

```bash
gator <command> [args...]
```

Available commands:
- `register <name>` - Create a new user.
- `login <name>` - Log in as an existing user.
- `addfeed <name> <url>` - Add a new RSS feed and follow it.
- `agg <time_between_reqs>` - Continuously fetch and save new posts (eg. `gator agg 1m`)
- `browse <limit>` - View recent posts from feeds you follow.
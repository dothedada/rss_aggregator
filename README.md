# RSS Aggregator CLI

RSS Aggregator is a command-line interface (CLI) for aggregating RSS feeds. It allows you to register as a user, add feeds, follow and unfollow feeds, and browse the latest posts from your followed feeds.

## Prerequisites

To run this program, you need to have the following installed:

* [Go](https://golang.org/)
* [PostgreSQL](https://www.postgresql.org/)

## Installation

To install the `rss_aggregator` CLI, run the following command:

```bash
go install github.com/dothedada/rss_aggregator
```

This will install the `rss_aggregator` binary in your Go bin directory.

## Configuration

Before running the program, you need to set up the configuration file.

1. Create a file named `.rss_aggregator.json` in your home directory.
2. Add the following content to the file:

```json
{
  "db_url": "YOUR_POSTGRES_CONNECTION_STRING"
}
```

Replace `"YOUR_POSTGRES_CONNECTION_STRING"` with your actual PostgreSQL connection string. For example:

```json
{
  "db_url": "postgres://user:password@localhost:5432/rss_aggregator?sslmode=disable"
}
```

## Running the Program

Once you have installed and configured the program, you can run it from the command line.

### Commands

Here are some of the available commands:

* `register <username> <password>`: Register a new user.
* `login <username> <password>`: Log in as a user.
* `addfeed <name> <url>`: Add a new RSS feed.
* `follow <url>`: Follow a feed.
* `following`: List the feeds you are currently following.
* `browse`: Browse the latest posts from your followed feeds.
* `unfollow <url>`: Unfollow a feed.
* `users`: List all registered users.
* `feeds`: List all available feeds.
* `agg <duration>`: Periodically scrapes feeds. e.g., `agg 1m`.

### Example Usage

```bash
# Register a new user
rss_aggregator register myuser mypassword

# Log in
rss_aggregator login myuser mypassword

# Add a new feed
rss_aggregator addfeed "Bla bla bla" https://some-shit.dev/rss.xml

# Follow a feed
rss_aggregator follow https://some-shit.dev/rss.xml

# List followed feeds
rss_aggregator following

# Browse posts
rss_aggregator browse
```

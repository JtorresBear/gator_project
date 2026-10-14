# Gator CLI


## Overview

Gator is an RSS feed aggregator that works with websites offering RSS feeds. It retrieves data through HTTP requests and parses the XML responses, allowing users to view articles from the blogs and websites they follow in one place.

Users can follow multiple feeds and browse their latest posts directly from the terminal.

## Features

- **User Profiles:** Create and switch between different users without passwords or authentication.
- **Feed Management:** Add RSS feeds and follow or unfollow feeds for each user.
- **Feed Aggregation:** Periodically check RSS feeds for new posts and save them to PostgreSQL.
- **Post Browsing:** View the latest posts from feeds the current user follows.
- **Persistent Storage:** Store users, feeds, and posts in a database for future use.

## Prerequisites

Before installing and running Gator, make sure you have:

- [Go](https://go.dev/doc/install) (version 1.27.1 or newer) — Required to compile and install Gator.
- [PostgreSQL](https://www.postgresql.org/download/) — Required to store users, feeds, and posts.
- [Git](https://git-scm.com/install/) — Required to clone the repository for database setup.
- [Goose](https://github.com/pressly/goose) — Required to apply database migrations.

To install Goose using Go:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Make sure your Go binary directory (typically `~/go/bin`) is included in your system's `PATH`.

## Installation

### 1. Install Gator

Make sure Go is installed, then run:

```bash
go install github.com/JtorresBear/gator_project/cmd/gator@latest
```

This downloads, compiles, and installs Gator.

### 2. Set Up PostgreSQL

Make sure PostgreSQL is installed and running.

Create a database named `gator`:

```sql
CREATE DATABASE gator;
```

You'll need a PostgreSQL user with permission to access the database and create tables.

### 3. Configure Gator

Create a file named `.gatorconfig.json` in your home directory (`~` on Linux/macOS).

Add the following:

```json
{
  "db_url": "postgres://postgres:YOUR_PASSWORD@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```

Replace `YOUR_PASSWORD` with your PostgreSQL password. Change the username, host, or port if your setup differs.

Gator uses this file to connect to PostgreSQL and remember the currently selected user.

### 4. Set Up the Database Schema

Clone the repository to access the database migrations:

```bash
git clone https://github.com/JtorresBear/gator_project.git
cd gator_project
```

With Goose installed, apply the migrations:

```bash
goose -dir sql/schema postgres "postgres://postgres:YOUR_PASSWORD@localhost:5432/gator?sslmode=disable" up
```

Use the same database credentials you configured in `.gatorconfig.json`.

### 5. Run Gator

Once installation and database setup are complete, register a user:

```bash
gator register yourusername
```

You can then run commands such as:

```bash
gator feeds
gator browse 5
```

To periodically fetch and save RSS articles, run the aggregator in a separate terminal:

```bash
gator agg 1m
```

This checks RSS feeds every minute while the aggregator is running. Press `Ctrl + C` to stop it.

## Commands

Gator provides several commands for managing users, RSS feeds, and posts.

| Command | Description |
|---|---|
| `gator register <username>` | Register a new user and log in. |
| `gator login <username>` | Switch to an existing user. |
| `gator users` | List all registered users. |
| `gator addfeed <name> <url>` | Add a new RSS feed and follow it. |
| `gator feeds` | Display all available feeds. |
| `gator follow <url>` | Follow an existing feed. |
| `gator following` | List feeds followed by the current user. |
| `gator unfollow <url>` | Unfollow a feed. |
| `gator agg <duration>` | Periodically fetch and save RSS posts. |
| `gator browse [limit]` | Display the latest posts from followed feeds. Defaults to 2 posts. |
| `gator reset` | Delete all registered users and their associated data. **Warning: This is destructive and cannot be undone.** |

### Example Usage

Register a user:

```bash
gator register sonic
```

Add an RSS feed:

```bash
gator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
```

Start collecting posts:

```bash
gator agg 1m
```

In another terminal, browse the latest posts:

```bash
gator browse 5
```

The `agg` command continues running until stopped with `Ctrl + C`.

## Technical Details

**RSS Parsing:** Each RSS feed is retrieved through an HTTP request and parsed into Go structs using the `encoding/xml` package and `xml.Unmarshal()`.

**Database Management:** Gator uses PostgreSQL to store and manage users, feeds, subscriptions, and posts. SQL queries are used to create, retrieve, and manage the stored information.

**SQL and Migrations:** Goose is used to manage database migrations, including creating tables and defining relationships between them. SQL queries are written separately, and sqlc generates Go code that allows the application to execute those queries.

**Feed Aggregation:** The aggregator runs at a user-specified interval, retrieving RSS data from websites and saving new posts to PostgreSQL. Duplicate posts are skipped to avoid storing the same article multiple times.

## Skills Practiced

- **Go Programming:** Building CLI applications, handling errors, and organizing code into packages.
- **PostgreSQL & SQL:** Writing queries, joins, and managing relational data.
- **Database Design:** Creating tables, foreign keys, and relationships.
- **Goose:** Managing database migrations.
- **sqlc:** Generating type-safe Go code from SQL queries.
- **HTTP & XML:** Retrieving RSS feeds and parsing XML into Go structs.
- **Data Persistence:** Storing and retrieving application data across sessions.   
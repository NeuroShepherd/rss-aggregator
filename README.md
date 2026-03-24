# RSS Aggregator ("gator")

This project uses Go and Postgres to create a minimal RSS aggregator tool. You can download the project code with:

```bash
git clone git@github.com:NeuroShepherd/rss-aggregator.git
cd rss-aggregator
```

You will then need to set up a PostgreSQL database (assuming you already have it installed):

```bash
psql postgres
```

Inside of Postgres:

```sql
CREATE DATABASE rss_aggregator;
```

Add a config file at `~/.gatorconfig.json` of the form:

```json
{
  "db_url": "postgres://username:password@localhost:5432/rss_aggregator?sslmode=disable",
  "current_user_name":""
}
```

Run the migrations to build the databases:

```bash
goose postgres postgres://username:password@localhost:5432/rss_aggregator up ./sql/schema
```

And finally, either run the project:

```bash
go run . <args>
```

Or build and then run:

```bash
go build
rss-aggregator <args>
```

## Available Commands

### Authentication
- **`login <username>`** - Log in as an existing user
- **`register <username>`** - Create a new user and log in

### User Management
- **`users`** - List all registered users

### Feed Management
- **`addfeed <name> <url>`** - Add a new RSS feed (requires login)
- **`feeds`** - List all available feeds with their creators
- **`follow <feed_url>`** - Follow an RSS feed (requires login)
- **`following`** - List feeds you are currently following (requires login)
- **`unfollow <feed_url>`** - Unfollow an RSS feed (requires login)

### Feed Aggregation
- **`agg <duration>`** - Start the feed aggregator that scrapes feeds at the specified interval (e.g., `2m`, `30s`, `1h30m`)

### Database
- **`reset`** - Drop and recreate the database, running all migrations


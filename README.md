# Go Blog Aggregator

A command-line RSS/blog feed aggregator built in Go that allows users to follow multiple RSS feeds and browse posts in one place.

## Features

- **User Management**: Register and login system with persistent user sessions
- **Feed Management**: Add, list, follow, and unfollow RSS feeds
- **Post Aggregation**: Automatically fetch and store posts from followed feeds
- **Browse Posts**: View latest posts from your followed feeds
- **PostgreSQL Backend**: Robust database storage using PostgreSQL
- **Type-safe Database Queries**: Generated using sqlc for compile-time safety

## Prerequisites

- Go 1.24.0 or higher
- PostgreSQL database
- [sqlc](https://docs.sqlc.dev/en/latest/overview/install.html) (for code generation)
- [goose](https://github.com/pressly/goose) (for database migrations)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/lukaszzieba/go-blog-agregator.git
cd go-blog-agregator
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your PostgreSQL database and run migrations:
```bash
# Navigate to sql/schema directory
cd sql/schema
goose postgres "your-database-url" up
```

4. Build the application:
```bash
go build -o gator
```

## Configuration

The application uses a configuration file stored at `~/.gatorconfig.json`. This file is automatically created when you first register or login and contains:
- Database connection URL
- Current logged-in user information

## Usage

### User Management

**Register a new user:**
```bash
./gator register <username>
```

**Login as existing user:**
```bash
./gator login <username>
```

**Reset configuration:**
```bash
./gator reset
```

**List all users:**
```bash
./gator users
```

### Feed Management

**Add a new RSS feed:**
```bash
./gator addfeed <feed-name> <feed-url>
```

**List all feeds:**
```bash
./gator feeds
```

**Follow a feed:**
```bash
./gator follow <feed-url>
```

**Unfollow a feed:**
```bash
./gator unfollow <feed-url>
```

**List feeds you're following:**
```bash
./gator following
```

### Content Aggregation

**Aggregate posts from all feeds:**
```bash
./gator agg <time-between-requests>
```
Example: `./gator agg 1s` (aggregates with 1 second between requests)

**Browse latest posts:**
```bash
./gator browse [limit]
```
Example: `./gator browse 10` (shows 10 latest posts)

## Database Schema

The application uses the following main tables:

- **users**: Store user accounts
- **feeds**: RSS feed information and ownership
- **feed_follows**: Many-to-many relationship between users and feeds
- **posts**: Aggregated blog posts from feeds

## Development

### Code Generation

This project uses sqlc for type-safe database interactions. To regenerate database code after modifying SQL queries:

```bash
sqlc generate
```

### Database Migrations

New migrations should be added to `sql/schema/` following the naming convention:
```
XXX_description.sql
```

Run migrations with goose:
```bash
goose postgres "your-database-url" up
```

## Dependencies

- `github.com/lib/pq`: PostgreSQL driver
- `github.com/google/uuid`: UUID generation

## License

[Add your license information here]

## Contributing

[Add contributing guidelines here]
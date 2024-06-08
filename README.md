# FeedFlow - An Automated RSS Feed Aggregator and Reader API

FeedFlow is a backend service that allows users to manage RSS feeds and posts. It provides API endpoints for user creation, feed management, following feeds, and retrieving posts. The service regularly fetches and processes RSS feed data, storing posts in a PostgreSQL database.

## Features

- **User Management**: Create and retrieve users.
- **Feed Management**: Add and manage RSS feeds.
- **Follow Feeds**: Follow and unfollow feeds.
- **Post Management**: Fetch, store, and retrieve posts from RSS feeds.
- **Automated Scraping**: Regularly fetch and process RSS feed data.

## Technologies Used

- Go
- PostgreSQL
- Chi (router)
- GoDotEnv
- Google UUID

## Getting Started

### Prerequisites

- Go (1.16+)
- PostgreSQL
- A `.env` file with the following environment variables:
  - `PORT`: The port number for the HTTP server.
  - `DATABASE_URL`: The PostgreSQL connection string.

### Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/ndangi168/FeedFlow---An-Automated-RSS-Feed-Aggregator-and-Reader-API
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up the PostgreSQL database and create the required tables.

4. Create a `.env` file in the project root with your configuration:
   ```env
   PORT=8080
   DATABASE_URL=postgres://user:password@localhost/dbname?sslmode=disable
   ```

### Running the Application

Start the application:
```bash
go run main.go
```

### API Endpoints

- **User Management**
  - `POST /v1/users`: Create a new user.
  - `GET /v1/users`: Get user information.

- **Feed Management**
  - `POST /v1/feeds`: Create a new feed (requires authentication).
  - `GET /v1/feeds`: Get all feeds.

- **Feed Follow Management**
  - `GET /v1/feed_follows`: Get all followed feeds (requires authentication).
  - `POST /v1/feed_follows`: Follow a feed (requires authentication).
  - `DELETE /v1/feed_follows/{feedFollowID}`: Unfollow a feed (requires authentication).

- **Post Management**
  - `GET /v1/posts`: Get posts from followed feeds (requires authentication).

- **Health Check**
  - `GET /v1/healthz`: Health check endpoint.

### Project Structure

- `main.go`: The entry point of the application.
- `internal/database/`: Database-related code and queries.
- `handlers.go`: Handlers for API endpoints.
- `models.go`: Data models and conversion functions.
- `scraper.go`: Functions for scraping and processing RSS feeds.
- `utils.go`: Utility functions for JSON responses and error handling.

### Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

### License

This project is licensed under the MIT License.

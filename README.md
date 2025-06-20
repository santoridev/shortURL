

# URL Shortener API

A simple URL shortening service built with Go and the Gin framework.
It allows you to shorten long URLs into short keys and redirect users to the original URLs.

---

## Features

* **Shorten URL:** `POST /shorten` — submit a long URL and get a short URL key.
* **Redirect:** `GET /:shorturl` — redirect short URL to the original long URL.

---

## Getting Started

### Prerequisites

* Go 1.18+ installed


### Setup and Run

1. Clone the repository or create your project folder:

```bash
git clone https://github.com/santoridev/shortURL.git
cd shortURL
```

2. Initialize Go modules (if not done):

```bash
go mod init github.com/santoridev/shortURL
```

3. Download dependencies:

```bash
go mod tidy
```

4. Run the application:

```bash
go run main.go
```

The server will run on `http://localhost:8080`.

---

## API Usage

### Shorten a URL

```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com/some/long/path"}'
```

**Response:**

```json
{
  "short_url": "/a1b2C"
}
```

### Redirect to Original URL

Open in browser or send GET request:

```
http://localhost:8080/a1b2C
```

This will redirect to the original long URL.

---

## How it works

* The service generates a random 5-character alphanumeric string as the short key.
* Stores a mapping in memory (`urlMap`).
* Redirects requests from the short key to the original URL.
* Note: Data is **not persisted**; restarting the server clears all mappings.

---

## Notes

* This is a simple in-memory implementation without persistent storage.
* No URL validation or duplication checks.
* Designed as a learning/demo project.



# Snippetbox

A small web application written in Go for storing and sharing short text snippets. This project was created as a learning exercise to explore building web apps with Go's standard library and common project patterns.

## Features
- Create, view and list text snippets
- Simple HTML templates and static assets
- Built with Go modules and standard library primitives
- Embed Filesystem for static assets and templates
- User authentication (signup, login, logout)
- Unit tests for helpers and commonHeaders middleware with mock

## Prerequisites
- Go 1.18+ installed
- (Optional) A database if your fork adds persistence

## Quick start
1. Clone the repository:
    git clone this project
2. From the project root, run:
    go run ./cmd/web
3. Open your browser at:
    http://localhost:4000

(If the project uses a different entry point, run `go run ./cmd/web` or follow the repository-specific instructions.)

## Development
- Build: `go build ./...`
- Test: `go test ./...`
- Lint/format: `gofmt -w .` and any linters you prefer

## Project layout (typical)
- cmd/web — application entry point
- internal — application code (handlers, models, helpers)
- templates — HTML templates
- static — CSS, JS, images
- go.mod, go.sum — module files

## Contributing
Feel free to open issues or pull requests. Keep changes small and focused, and include tests where appropriate.

## License
MIT — see LICENSE file for details.
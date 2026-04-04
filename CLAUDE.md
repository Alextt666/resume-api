# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

A lightweight Go HTTP API that serves resume data from a JSON file (`data/resume.json`). The API uses in-memory caching with a 5-minute TTL to avoid repeated file reads.

## Architecture

- **Single binary deployment**: Uses multi-stage Dockerfile (golang:1.24-alpine → alpine:3.19) to produce a minimal production image
- **No external dependencies**: Standard library only (no web frameworks)
- **Data persistence**: JSON file in `data/` directory, mounted as volume in Docker for easy updates without rebuild

### Key Files

- `main.go` - Entry point, sets up HTTP routes and starts server
- `handlers/resume.go` - Resume data handler with read-write lock protected caching
- `data/resume.json` - Resume data source (bilingual: Chinese/English)

## Common Commands

```bash
# Run locally
go run main.go

# Build binary
go build -o resume-api .

# Run with Docker
docker-compose up -d

# Run tests (if added)
go test ./...

# Format code
go fmt ./...
```

## Development Notes

- Default port is 8080, configurable via `PORT` environment variable
- The `/health` endpoint returns "ok" for health checks
- CORS is enabled for all origins on `/api/resume`
- Cache TTL is 5 minutes; edits to `data/resume.json` require container restart to take effect immediately, or wait for cache expiration

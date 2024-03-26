# Austere

Austere is a dockerized Golang API built to download music from multiple streaming services, mainly to use as Spotify local files, with Prometheus and Grafana logging.

Supported Sites:

- Youtube
- Soundcloud
- Krakenfiles

Austere uses `yt-dlp` to download media from supported sites (and a custom solution for Krakenfiles).

## Prerequisites

## Local Setup

To get started with the setup, clone the repository and navigate into the directory using the following commands:

```bash
git clone https://github.com/gursheyss/austere
cd austere
```

Once you're inside the project directory, install the project dependencies with:

```bash
go mod download
```

To run the application, run:

```bash
go run main.go
```

## Deployment

To prepare the application for deployment, build the application with:

```bash
go build
```

Or run it with docker

```bash
docker compose up --build
```

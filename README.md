# Logistics app

Live preview:

## Architecture

For this case I've chosen a simple architecture for the live environment. The Golang server is going to be responsible for delivering the static files to the browser and also handling the API requests.

In more serious cases the architecture would be more separated.

## Persistance Layer

The app is storing the package configuration in a DynamoDB table for reference.

## How to start the app

### With Docker

You can just type this command:

```bash
make docker-build-run
```

### Without Docker

There is an option to run it with MAKE but not sure it's gonna work:

```bash
make dev
```

But without MAKE you can simply do:

```bash
cd frontend && pnpm run dev
cd .. && go run cmd/main.go
```

## Depoloyment

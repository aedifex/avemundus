# Go Hello World – Harness CI Demo

This repository contains a minimal, production-grade Go “hello world” service, purpose-built to demo Harness CI features. It’s a fast, clean starting point for testing CI flows, running builds, and generating real artifacts.

## Features

- Simple Go app (`main.go`) that prints “hello, world” and serves a `/healthz` endpoint
- Unit tests included (`*_test.go`)
- Makefile for local builds and testing
- Dockerfile for containerization
- Pipeline covers build, test, lint, and container image creation
- Fast to set up and easy to extend

## Quickstart

1. **Clone this repo**
2. Run tests and build locally:
    ```sh
    make test
    make build
    ```
3. Build the Docker image:
    ```sh
    docker build -t go-hello-world .
    ```
4. Run locally and test health endpoint:
    ```sh
    docker run --rm -p 8080:8080 go-hello-world
    curl localhost:8080/healthz
    ```

## Harness CI Pipeline

- Restores and saves Go build/test cache for fast feedback
- Runs lint (`go vet` and `golangci-lint`) and unit tests (with code coverage)
- Builds the binary with version and commit injected for traceability
- Builds and pushes a Docker image (requires registry secrets set up in Harness)
- Generates and uploads build artifacts (binary, test coverage, SBOM)

## Key Files

- `main.go` – Go app entrypoint
- `main_test.go` – Unit tests
- `Makefile` – Local build/test commands
- `Dockerfile` – Container build instructions
- `.harness/ci.yaml` – Harness CI pipeline definition

## Local Development

All standard dev/test tasks are handled via Makefile:

```sh
make test      # Run unit tests
make build     # Build the binary
make docker    # Build Docker image
```

To run the app and verify it works:

```sh
./go-hello-world
curl localhost:8080/healthz
```

Or with Docker:

```sh
docker run --rm -p 8080:8080 go-hello-world
curl localhost:8080/healthz
```

## Customization

Add more endpoints, tests, or pipeline steps as needed—this repo is meant to be forked and extended.  
Open an issue or PR if you want to contribute or report a problem.

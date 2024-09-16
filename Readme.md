# WebHook API

## Overview

This project is a WebHook API built using Go. It provides functionalities to create and manage webhooks, and includes a
hook handler for processing webhook events.

## Features

- Create and manage webhooks
- Process webhook events
- Swagger documentation for API endpoints

## Project Structure

- `cmd/`: Contains the command-line interface for the API and handler.
- `internal/`: Contains the core business logic and domain models.
- `pkg/`: Contains utility packages like logging and middleware extensions.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- MongoDB
- Redis
- Docker and Docker Compose

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/mehtiumit/webhook-api.git
    cd webhook-api
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

### Running the API

1. Start MongoDB, Redis Api and Job services using Docker Compose:
    ```sh
    docker-compose up -d
    ```

2. Run the API:
    ```sh
    go run main.go hook-api
    ```

3. Access the API at `http://localhost:5030/webhook/api`.

### Running the Hook Handler

1. Run the hook handler:
    ```sh
    go run main.go hook-handler
    ```

## Swagger Documentation

The API documentation is generated using Swagger. To generate the documentation:

1. Run the following command:
    ```sh
    swag init -g ./cmd/hookApi.go -o ./docs/hookApi --instanceName hookApi
    ```

2. Start the API.

3. Open your browser and navigate to [Swagger Documentation](http://localhost:5030/webhook/swagger/index.html).

## Creating Hooks

Before creating a hook, you need to create content first. Ensure that the content is created and available in the
database.

## Development Mode

If you are in development mode, please change the `MONGO_URL` environment variable to point to your local MongoDB
instance. On /internal/mongodb repositories

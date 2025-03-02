# Todo Application

## Overview

This is a simple Todo Application built with **Golang**, designed to manage user-specific Todo lists. Users can create, update, delete Todo items and attach files such as images, videos, or documents. The application leverages **Gin** for routing, **MongoDB** as the database, and follows **Hexagonal Architecture** for clean and maintainable code.

## Features

- **User-specific Todo lists**: Each user can manage their own Todo items.
- **Attachment support**: Attach multiple files (photos, videos, audios, etc.) to each Todo.
- **CRUD Operations**: 
  - Create new Todos.
  - Get Todos by user.
  - Update Todo details.
  - Delete Todo items.
  - Add or delete attachments to Todos.
- **REST API**: A fully functional API to interact with Todo data.
- **Logging**: Uses **Zap** for structured logging.
- **Cloud-Deployable**: Supports Docker, Helm, and environment-based configurations.

## Architecture

The application follows **Hexagonal Architecture** (also known as Ports and Adapters). It isolates core business logic from external systems such as databases, web frameworks, and external APIs.

- **Core**: Contains the business logic (Todo service and entities).
- **Adapters**: Contains infrastructure code (database repositories, API handlers, etc.).
- **Ports**: Define the interfaces between the core business logic and the external systems.

## Technology Stack

- **Go 1.20** (Golang)
- **Gin** (Web framework)
- **MongoDB** (Database)
- **Viper** (Configuration management)
- **Zap** (Logging)
- **Docker** (Containerization)
- **Helm** (Deployment orchestration)

## Requirements

- **Go** (Golang) 1.20 or above
- **Docker** (for building and running in containers)
- **MongoDB** (locally or cloud-based instance)
- **Git** (for version control)

## Setup and Installation

### Clone the Repository

```bash
git clone git@github.com:anilvpatel21/todo.git
cd todo
```
## Install dependencies using Go modules:

```bash
go mod tidy
```

## Running Locally

```bash
go run main.go
```

## Docker
```bash
docker build -t todo-app .

docker run -p 8080:8080 todo-app
```

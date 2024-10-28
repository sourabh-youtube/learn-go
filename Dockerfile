# syntax=docker/dockerfile:1
FROM golang:1.23-alpine

# Install dependencies
RUN apk add --no-cache git

# Install Air for hot-reloading, using the correct path
RUN go install github.com/air-verse/air@v1.61.1

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Set the entry point to Air
CMD ["air"]

# Start from golang base image
FROM golang:alpine as builder

# Add maintainer info
LABEL maintainer="Stephen Aribaba <stephen.aribaba@gmail.com>"

# Install git
RUN apk update && apk add --no-cache git

# Disable gcc
ENV CGO_ENABLED=0

# Set the current work directory inside the container
WORKDIR /api

# Copy go mod and sum file into the work directory
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Download CompileDaemon for hot reload
RUN go get github.com/githubnemo/CompileDaemon

# Copy source code from project directory to the current work directory in the container
COPY . .


# Expose port 3000 to the outside world
# EXPOSE 3000

# Entrypoint
ENTRYPOINT CompileDaemon -pattern="(.+\.go|.+\.c|.+\.gohtml)$" -build="go build -a -o main ./src/main.go" -command="./main"


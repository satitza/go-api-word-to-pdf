# Use an official Golang image as the base image
FROM golang:1.22-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Copy the config file into the container
# COPY configuration/config.yaml /app/config.yaml

# Build the Go app
RUN go build -o /app/gin-api

# Use a new stage to install LibreOffice
FROM alpine:latest

# Install LibreOffice
RUN apk add --no-cache libreoffice

WORKDIR /

COPY --from=builder /app/configuration ./configuration

# Copy the Go app from the builder stage
COPY --from=builder /app/gin-api .

# Copy the config file from the builder stage
# COPY --from=builder /app/config.yaml /app/config.yaml

# Expose port 8080 to the outside world
EXPOSE 4444

# Command to run the executable
CMD ["./gin-api"]
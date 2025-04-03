# Use Go 1.21 as the base image for building
FROM golang:1.23-alpine AS builder


# Set the working directory
WORKDIR /app

# Set Go proxy
RUN go env -w GOPROXY=https://goproxy.cn,direct

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application with CGO enabled
RUN go build -o augeu_server backEnd/cmd/main.go

# Use a minimal base image for the final container
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/augeu_server .

# Copy the configuration file to the correct path
COPY backEnd/etc/config.yaml /app/backEnd/etc/config.yaml

# Expose the necessary ports
EXPOSE 8080 59191

# Run the application
CMD ["./augeu_server"]
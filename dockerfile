# Step 1: Build the Go application
FROM golang:1.23-alpine as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Copy the entire project
COPY . .

# Build the Go app
RUN GOOS=linux go build -o todo-app

# Step 2: Create a smaller image to run the app
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/todo-app .

# Copy the config folder into the container
COPY --from=builder /app/config /root/config

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./todo-app"]

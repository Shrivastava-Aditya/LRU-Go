# Use the official Golang image as the base image
FROM golang:1.17-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY main.go cache.go ./

# Build the Go application
RUN go build -o main .

# Use a lightweight base image
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /root

# Copy the built executable from the builder stage
COPY --from=builder /app/main .

# Command to run the executable
CMD ["./main"]

# Dockerfile
# Use the official Go image with the version specified in your go.mod
FROM golang:1.21-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o main .

# Use a lightweight image for the final build
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/main .

# Expose the port your application will run on
EXPOSE 5000

# Run the application
CMD ["./main"]

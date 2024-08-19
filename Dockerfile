# First stage: build the Go application
FROM golang:1.22.5-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main main.go

# Second stage: create a smaller image for deployment
FROM alpine:3.18

# Set the working directory inside the container
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Command to run the application
CMD ["/app/main"]

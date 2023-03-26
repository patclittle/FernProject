# Start from the official Go image
FROM golang:1.17-alpine3.14

# Set the current working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go binary
RUN go build -o main .

# Expose the port that the application listens on
EXPOSE 8080

# Start the application
CMD ["./main"]

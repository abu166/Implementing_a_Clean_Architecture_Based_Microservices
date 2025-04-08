# Use the official Golang image as the base image
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o api-gateway .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application
CMD ["./api-gateway"]
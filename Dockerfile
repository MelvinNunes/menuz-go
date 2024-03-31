FROM golang:1.21.3-alpine

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN go build -o main cmd/api/main.go

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
ENTRYPOINT ["./main"]
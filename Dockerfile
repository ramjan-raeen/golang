# Dockerfile
FROM golang:1.24

# Install Air for live reload
# RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/air-verse/air@latest

# Set working directory
WORKDIR /app

# Copy Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire app source
COPY . .

# Expose the port used in main.go
EXPOSE 8080
# Create a temporary directory for Air
RUN mkdir -p tmp
# Start app with Air

CMD ["air", "-c", ".air.toml"]

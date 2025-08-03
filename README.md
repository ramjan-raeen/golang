docker run --rm -it -v "$PWD":/app -w /app golang:1.24 go mod init go-project
docker run --rm -v "$PWD":/app -w /app golang:1.24 go mod tidy
docker run --rm -v "$PWD":/app -w /app -p 8080:8080 golang:1.24 go run main.go




# go mod init htmx-go-app
go mod tidy

go run cmd/api/main.go

# go build -o bin/api ./cmd/api && \
# go build -o bin/worker ./cmd/worker

# ./bin/api & 
# ./bin/worker

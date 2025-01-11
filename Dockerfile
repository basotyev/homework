FROM golang:latest
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/basic_server/main.go
COPY ./internal/app/db/migrations ./internal/app/db/migrations

CMD ["go", "run", "/app/main"]
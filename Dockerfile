FROM golang:1.23.4-alpine3.19 AS build-stage
WORKDIR /app
COPY go.mod go.sum /
RUN go mod download
COPY . .
RUN go build -o main ./cmd/basic_server/main.go



FROM alpine:3.19 AS runner
WORKDIR app
COPY --from=build-stage /app/internal/app/db/migrations ./internal/app/db/migrations
COPY --from=build-stage /app/main .
EXPOSE 8080
CMD ["/app/main"]
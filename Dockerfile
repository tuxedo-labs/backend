FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./cmd/server.go

FROM debian:bullseye-slim

COPY --from=builder /app/server /server

EXPOSE 8080

CMD ["/server"]


FROM golang:1.22 AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN GOARCH=amd64 go build -o main ./cmd/main.go

RUN ls -la /app
RUN chmod +x /app/main

FROM alpine:latest

WORKDIR /root/
COPY .env /root/.env
COPY --from=builder /app/main /root/

EXPOSE 8080
CMD ["/root/main"]

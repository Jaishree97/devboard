# Stage 1: Build the Go application
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o devboard-api .


# Stage 2: Run the application
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/devboard-api .

EXPOSE 8080

CMD ["./devboard-api"]
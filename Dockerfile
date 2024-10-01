# Build stage
FROM golang:1.21.0 AS builder

WORKDIR /usr/src/app

COPY . .

RUN go mod download && \
    go mod tidy && \
    go build -o ./bin ./main.go
    

ENTRYPOINT ["./bin"]
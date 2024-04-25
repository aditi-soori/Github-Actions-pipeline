FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY main.go .
RUN go build -o main -ldflags=-X=main.version=${VERSION} main.go 

EXPOSE 8080
CMD ["main"]

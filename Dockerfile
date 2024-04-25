FROM golang:1.19.3-alpine AS builder

ARG VERSION=dev

WORKDIR /go/src/app
COPY main.go .
RUN go build -o main -ldflags="-X=main.version=${VERSION}" main.go 

FROM alpine:latest
COPY --from=builder /go/src/app/main /go/bin/main
EXPOSE 8080
CMD ["/go/bin/main"]

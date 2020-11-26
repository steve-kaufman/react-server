FROM golang:alpine as builder

WORKDIR /go/src/server

COPY main.go .

RUN go build -o server


FROM alpine:latest

COPY --from=builder /go/src/server/server /usr/bin

EXPOSE 8080

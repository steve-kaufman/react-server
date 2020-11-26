FROM golang:alpine as builder

WORKDIR /go/src/server

COPY main.go .

RUN go build -o server


FROM alpine:latest

ENV PATH=$PATH:/bin

COPY --from=builder /go/src/server/server /bin

EXPOSE 8080

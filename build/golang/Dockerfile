# Stage 1: Build
FROM golang:1.22.5 AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0

ENV GOOS linux

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /app

ADD go.mod .
ADD go.sum .

RUN go mod download

COPY . .

RUN go build -ld-flags="-s -w" -o wishverse . /cmd/wishverse/main.go

#Stage 2: Prod. Image
FROM alpine:latest

RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /root

COPY --from=builder /app/wishverse .

CMD ["./wishverse"]

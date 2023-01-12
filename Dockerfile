FROM golang:1.19 AS builder
WORKDIR /go/src/github.com/SirusDoma/btc-wallet
COPY . ./
RUN go mod tidy
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o app ./cmd/app/main.go

FROM alpine:latest
RUN apk update && apk add tzdata &&\
    apk --no-cache add ca-certificates

WORKDIR /usr/local/bin
COPY --from=builder go/src/github.com/SirusDoma/btc-wallet/app ./
CMD ["./app"]
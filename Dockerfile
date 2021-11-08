FROM golang:1.15-stretch AS builder
WORKDIR /src
COPY . /src
RUN cd /src && go build -o main cmd/main.go

FROM alpine:latest as alpine
RUN apk add -U --no-cache ca-certificates tzdata


FROM debian:buster-slim
WORKDIR /app
COPY --from=builder /src/main /app/
COPY --from=alpine /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=alpine /usr/share/zoneinfo /usr/share/zoneinfo
CMD ["/app/main"]
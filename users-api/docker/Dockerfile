FROM golang:1.20-alpine as builder
RUN apk update && apk add --no-cache git make ca-certificates tzdata && update-ca-certificates
RUN adduser -D -g '' appuser
WORKDIR /go/src/app
COPY . .
RUN env GO='CGO_ENABLED=0 go' go build -o server cmd/combined/*.go
FROM alpine:3.13 as production
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/app/server /app
USER appuser
ENTRYPOINT ["/app"]

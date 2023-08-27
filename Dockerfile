# Dockerfile metadata
FROM golang:alpine3.18

# go-sqlite3 requires GCC and setting CGO_ENABLED=1
ENV CGO_ENABLED=1
RUN apk add build-base

WORKDIR /opt/ctfconsole/server
COPY server/ .
RUN go build -o /usr/local/bin/ctfconsole cmd/main.go
CMD ["ctfconsole", "-p", "${CTFCONSOLE_UI_PORT}"]
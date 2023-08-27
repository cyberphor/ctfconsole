# Dockerfile metadata
# Docker metadata
FROM golang:alpine3.18
LABEL version="v0.3.0"
LABEL description="ctfconsole is a Capture The Flag (CTF) server"
LABEL maintainer="Victor Fernandez III, @cyberphor"

# download GCC (required by go-sqlite3)
ENV CGO_ENABLED=1
RUN apk add build-base

# download dependencies
WORKDIR /opt/ctfconsole/server
COPY server/go.mod .
COPY server/go.sum .
RUN go mod download

# copy and build source code
COPY server/ .
RUN go build -o /usr/local/bin/ctfconsole cmd/main.go

# run ctfconsole
CMD ["ctfconsole", "-p", "${CTFCONSOLE_UI_PORT}"]
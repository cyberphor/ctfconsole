# Dockerfile metadata
# Docker metadata
FROM golang:alpine3.18
LABEL version="v0.3.0"
LABEL description="ctfconsole is a Capture The Flag (CTF) server"
LABEL maintainer="Victor Fernandez III, @cyberphor"

# download GCC (required by backend package: go-sqlite3)
ENV CGO_ENABLED=1
RUN apk add build-base

# download backend dependencies
WORKDIR /opt/ctfconsole/backend
COPY backend/go.mod .
COPY backend/go.sum .
RUN go mod download

# copy and build source code
COPY backend/ .
RUN go build -o /usr/local/bin/ctfconsole cmd/main.go

# run ctfconsole
CMD ["ctfconsole", "-c", "config.yaml"]
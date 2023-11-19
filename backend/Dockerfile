# Dockerfile metadata
FROM golang:alpine3.18
LABEL version="${VERSION}"
LABEL description="ctfconsole | a Capture the Flag (CTF) server"
LABEL maintainer="Victor Fernandez III, @cyberphor"

# update execution path
WORKDIR /usr/local/bin/ctfconsole/
ENV PATH="/usr/local/bin/ctfconsole/:${PATH}"

# download backend dependencies
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy and build source code
COPY . .
RUN go build -o ctfconsole cmd/main.go

# start ctfconsole api
ENTRYPOINT [ "ctfconsole" ]
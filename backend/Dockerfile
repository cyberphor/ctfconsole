# Dockerfile metadata
FROM golang:alpine3.18
LABEL version="v0.3.0"
LABEL description="ctfconsole | a Capture the Flag (CTF) server"
LABEL maintainer="Victor Fernandez III, @cyberphor"

# download backend dependencies
WORKDIR /opt/ctfconsole/backend
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy and build source code
COPY . .
RUN go build -o /usr/local/bin/ctfconsole cmd/main.go

# start ctfconsole backend
CMD ctfconsole \
  --log-path ${CTFCONSOLE_API_LOG_PATH} \
  --ui-proto ${CTFCONSOLE_UI_PROTOCOL} \
  --ui-ip ${CTFCONSOLE_UI_IP_ADDRESS} \
  --ui-port "${CTFCONSOLE_UI_PORT}" \ 
  --api-proto ${CTFCONSOLE_API_PROTOCOL} \
  --api-ip ${CTFCONSOLE_API_IP_ADDRESS} \
  --api-port "${CTFCONSOLE_API_PORT}" \
  --db-proto ${CTFCONSOLE_DB_PROTOCOL} \
  --db-ip ${CTFCONSOLE_DB_IP_ADDRESS} \
  --db-port "${CTFCONSOLE_DB_PORT}" \
  --db-name ${CTFCONSOLE_DB_NAME} \
  --db-user ${CTFCONSOLE_DB_USER} \
  --db-password ${CTFCONSOLE_DB_PASSWORD}
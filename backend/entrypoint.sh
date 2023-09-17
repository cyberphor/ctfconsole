#!/bin/sh

set -e

exec ctfconsole \
  --log-path "$CTFCONSOLE_API_LOG_PATH" \
  --ui-proto "$CTFCONSOLE_UI_PROTOCOL" \
  --ui-ip "$CTFCONSOLE_UI_IP_ADDRESS" \
  --ui-port "$CTFCONSOLE_UI_PORT" \
  --api-proto "$CTFCONSOLE_API_PROTOCOL" \
  --api-ip "$CTFCONSOLE_API_IP_ADDRESS" \
  --api-port "$CTFCONSOLE_API_PORT" \
  --db-proto "$CTFCONSOLE_DB_PROTOCOL" \
  --db-ip "$CTFCONSOLE_DB_IP_ADDRESS" \
  --db-port "$CTFCONSOLE_DB_PORT" \
  --db-name "$POSTGRES_DB_NAME" \
  --db-user "$POSTGRES_DB_USER" \
  --db-password "$POSTGRES_DB_PASSWORD"
func Set-Environment {
    # frontend settings
    $env:CTFCONSOLE_UI_IP_ADDRESS="localhost"
    $env:CTFCONSOLE_UI_PORT=80
    $env:CTFCONSOLE_UI_PROTOCOL="http"

    # backend settings
    $env:CTFCONSOLE_API_IP_ADDRESS="localhost"
    $env:CTFCONSOLE_API_PORT=8081
    $env:CTFCONSOLE_API_LOG_PATH="/var/log/ctfconsole/ctfconsole_api.log"
    $env:CTFCONSOLE_API_HEALTH_ENDPOINT="http://localhost:${CTFCONSOLE_API_PORT}/api/v1/ruok" 

    # database settings
    $env:CTFCONSOLE_DB_HOST="localhost"
    $env:CTFCONSOLE_DB_PORT="5432"
    $env:CTFCONSOLE_DB_NAME="ctfconsole"
    $env:CTFCONSOLE_DB_USER="postgres"
    $env:CTFCONSOLE_DB_PASSWORD="postgres"
    $env:CTFCONSOLE_DB_LOG_PATH="/var/log/ctfconsole/"
}

func Start-Database {
    docker compose --profile "ctfconsole_database" up --build -d
}

func Test-API {
    curl -s -X POST -H "Content-Type: application/json" -d '{"Name": "foo", "Password": "bar"}' localhost:8081/api/v1/player | jq
}
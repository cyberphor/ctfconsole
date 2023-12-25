func Set-Environment {
    # frontend settings
    $env:CTFCONSOLE_UI_IP_ADDRESS="localhost"
    $env:PORT=80
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
    Invoke-RestMethod -Uri "http://localhost:8081/api/v1/player" -Method POST -Headers @{"Content-Type"="application/json"} -Body (@{Name="foo"; Password="bar"} | ConvertTo-JSON)
    Invoke-RestMethod -Uri "http://localhost:8081/api/v1/player" -Method GET -Headers @{"Content-Type"="application/json"} -Body (@{Name="foo"} | ConvertTo-JSON)
}

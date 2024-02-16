curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"Name": "foo", "Password": "bar"}' \
  http://localhost:8081/api/v1/player -s | jq

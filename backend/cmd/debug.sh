curl -s -X POST -H "Content-Type: application/json" -d '{"Name": "foo", "Password": "bar"}' localhost:8081/api/v1/player | 
jq
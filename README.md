# StreakAI

## Instructions:

Run the following command:
```go run main.go```

### Sample Curl to check the API:

```curl --location 'http://localhost:8080/find-pairs' \
--header 'Content-Type: application/json' \
--data '{
    "array": [1,2,3,4,5],
    "target": 6
}'```

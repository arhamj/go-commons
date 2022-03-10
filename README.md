# Golang Commons Library

## Components

### HTTP Client

Library

```
github.com/go-resty/resty/v2
```

Example usage

```
client := httpClient.NewHttpClient(true, httpClient.ConfigOption(httpClient.Config{
    ClientTimeout: 10 * time.Second, // default: 5s
    RetryCount:    2,                // default: 3
}))

req := client.NewRequest()
resp, err := req.Get("https://www.boredapi.com/api/activity")
if err != nil {
    panic(err)
}

result := make(map[string]interface{})
body := resp.Body()
err = json.Unmarshal(body, &result)
if err != nil {
    panic(err)
}
```

### HTTP Errors

### HTTP Utils

### Interceptors

### Kafka

### Logger

### Mongo

### Postgres

### Checks and Probes

### Redis

### Tracing

### Utils

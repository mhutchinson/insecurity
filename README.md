# insecurity
Miscellaneous security & discovery tools.

## Portscanner

```bash
go run ./cmd/scan -alsologtostderr -v=3 -host="127.0.0.1" -workers=100
```

## Proxy

```bash
# Run the proxy
go run .\cmd\proxy -alsologtostderr -v=1

# Make requests through it (overriding virtual host header if needed)
curl --header "Host: example.com" localhost:8080
 ```
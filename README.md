# SM-Analytics

### Getting Started

- HTTP server runs on port 8000. 
- Clone the repo.
- Istio should be running in your local system.
- Install minikube, and follow this link ```https://istio.io/docs/setup/platform-setup/minikube/```
- - Deploy a Destination Rule for CBR.
```
 curl -XPOST -H "Content-type: application/json" -d '{
	"name": "cbr",
	"namespace": "default",
	"host": "host.svc.cluster.local",
	"tcpMaxConnections": 1,
	"httpMaxConnections": 2,
	"http2MaxConnections": 0,
	"consecutiveError": 1,
	"interval": 1,
	"baseEjectionTime": 180s,
	"maxEjectionPercent":1
	
}' 'localhost:8000/api/v1/cbr'
```
- See CIRCUITBREAK.md for more.


### Fault Domain

- A fault domain consists of the following aspects related to microservices :

```
Circuit Breaking
Fault Injection
Rate Limiting
Request Routing
Request Timeouts
```

### Istio Proxy Log

- In order to enable logging for istio proxy, enable logs in istio configmap named as istio. 
- ```istioctl manifest apply --set values.global.proxy.accessLogFile="/dev/stdout```
- Here is what a GET request log looks like

```
[2019-12-10T13:12:57.245Z] "GET /status/418 HTTP/1.1" 418 - "-" "-" 0 135 1 1 "-" "curl/7.64.0" "538be813-30b7-467f-b0c7-63b3bf451af0" "httpbin:8000" "127.0.0.1:80" inbound|8000|http|httpbin.bookinfo.svc.cluster.local - 100.99.7.120:80 100.99.50.3:46996 - default
```

### Understanding the Log

- Mapping to the log above
- <https://www.envoyproxy.io/docs/envoy/latest/configuration/observability/access_log>

| FORMAT STRING | LOG  |
| ------ | ------ |
| [START_TIME] |[2019-12-10T13:12:57.245Z]  |
| REQ(:METHOD) | GET |
| REQ(X-ENVOY-ORIGINAL-PATH?:PATH | /status/418 |
| PROTOCOL | HTTP/1.1 |
| RESPONSE_CODE | 418 |
| RESPONSE FLAG | - |
| BYTES RECIEVED | - |
| BYTES SENT | 0 |
| DURATION | 135 |
|RESP(X-ENVOY-UPSTREAM-SERVICE-TIME) | 1 |
|REQ(X-FORWARDED-FOR) | httpbin:8000 |
|REQ(USER-AGENT) | 127.0.0.1:80 |
|REQ(X-REQUEST-ID) | 538be813-30b7-467f-b0c7-63b3bf451af0 |
|REQ(:AUTHORITY)| 100.99.7.120:80 |
|UPSTREAM_HOST| default |

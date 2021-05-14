### Circuit Breaking

#### Introduction

- A circuit breaker can have three states:

1. Closed,
2. Open
3. Half Open

- By default exists in a closed state. In the closed state, requests succeed or fail until the number of failures reach a predetermined threshold, with no interference from the breaker.
- When the threshold is reached, the circuit breaker opens. When calling a service in an open state, the circuit breaker trips the requests, which means that it returns an error without attempting to execute the call.
- In this way, by tripping the request downstream at the client, cascading failures can be prevented in a production system. 
- After a configurable timeout, the circuit breaker enters a half open state, in which the failing service is given time to recover from its broken behavior. 
- If requests continue to fail in this state, then the circuit breaker is opened again and keeps tripping requests.
- Otherwise, if the requests succeed in the half open state, then the circuit breaker will close and the service will be allowed to handle requests again.

### Circuit Breaking In Istio

- In models a struct is defined for the circuit breaker object.
- Deploy a Destination Rule for CBR.
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

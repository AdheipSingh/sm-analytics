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
### Mixer Log

- An end to end trace of how a request traverses from ingress gateway to a pod through istio-proxy.

```
connection.mtls               : false
context.protocol              : http
context.reporter.kind         : outbound
context.reporter.uid          : kubernetes://istio-ingressgateway-cbc9c69b6-r7qfd.istio-system
destination.service.host      : helloworld.default.svc.cluster.local
destination.service.name      : helloworld
destination.service.namespace : default
destination.service.uid       : istio://default/services/helloworld
origin.ip                     : [10 0 83 159]
request.headers               : stringmap[accept:*/* user-agent:curl/7.58.0 x-forwarded-for:10.0.83.159 x-request-id:99d7fe58-8d26-491a-a60c-2ca6a4ea1cdc :authority:a883d391e757c11e9a86a06538fac772-1465315514.us-west-2.elb.amazonaws.com x-envoy-internal:true :method:GET :path:/hello x-forwarded-proto:http x-envoy-decorator-operation:helloworld.default.svc.cluster.local:5000/hello]
request.host                  : a883d391e757c11e9a86a06538fac772-1465315514.us-west-2.elb.amazonaws.com
request.method                : GET
request.path                  : /hello
request.scheme                : http
request.time                  : 2019-05-14 18:03:48.431822233 +0000 UTC
request.url_path              : /hello
request.useragent             : curl/7.58.0
source.namespace              : istio-system
source.uid                    : kubernetes://istio-ingressgateway-cbc9c69b6-r7qfd.istio-system
---
destination.container.name    : unknown
destination.ip                : [0 0 0 0 0 0 0 0 0 0 255 255 0 0 0 0]
destination.labels            : stringmap[]
destination.name              : unknown
destination.namespace         : default
destination.owner             : unknown
destination.serviceAccount    : unknown
destination.uid               : unknown
destination.workload.name     : unknown
destination.workload.namespace: unknown
destination.workload.uid      : unknown
source.ip                     : [0 0 0 0 0 0 0 0 0 0 255 255 100 121 215 132]
source.labels                 : stringmap[pod-template-hash:767572562 release:istio app:istio-ingressgateway chart:gateways heritage:Tiller istio:ingressgateway]
source.name                   : istio-ingressgateway-cbc9c69b6-r7qfd
source.owner                  : kubernetes://apis/apps/v1/namespaces/istio-system/deployments/istio-ingressgateway
source.serviceAccount         : istio-ingressgateway-service-account
source.workload.name          : istio-ingressgateway
source.workload.namespace     : istio-system
source.workload.uid           : istio://istio-system/workloads/istio-ingressgateway
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

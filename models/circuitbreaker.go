package models

// CircuitBreaker hold cbr object
type CircuitBreaker struct {
	Name                     string `json:"name"`
	Namespace                string `json:"namespace"`
	Host                     string `json:"host"`
	TCPMaxConnections        int32  `json:"tcpMaxConnections"`
	HTTP1MaxConnections      int32  `json:"httpMaxConnections"`
	HTTP2MaxConnections      int32  `json:"http2MaxConnections"`
	HTTP1MaxPendingRequests  int32  `json:"http1MaxPendingRequests"`
	MaxRequestsPerConnection int32  `json:"maxRequestsPerConnection"`
	ConsecutiveError         int32  `json:"consecutiveError"`
	Interval                 int64  `json:"interval"`
	BaseEjectionTime         int64  `json:"baseEjectionTime"`
	MaxEjectionPercent       int32  `json:"maxEjectionPercent"`
}

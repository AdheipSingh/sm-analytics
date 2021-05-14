package circuit

import (
	"github.com/gogo/protobuf/types"
	networkingv1alpha3 "istio.io/api/networking/v1alpha3"

	"sm-analytics/models"
	"istio.io/client-go/pkg/apis/networking/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewCircuitbreaker returns a circuit breaker
func NewCircuitbreaker(c *models.CircuitBreaker) *v1alpha3.DestinationRule {
	cb := &v1alpha3.DestinationRule{
		ObjectMeta: v1.ObjectMeta{
			Name:      c.Name,
			Namespace: c.Namespace,
		},
		Spec: networkingv1alpha3.DestinationRule{
			Host: c.Host,
			TrafficPolicy: &networkingv1alpha3.TrafficPolicy{
				ConnectionPool: &networkingv1alpha3.ConnectionPoolSettings{
					Tcp: &networkingv1alpha3.ConnectionPoolSettings_TCPSettings{
						MaxConnections: c.TCPMaxConnections,
					},
					Http: &networkingv1alpha3.ConnectionPoolSettings_HTTPSettings{
						Http2MaxRequests:         c.HTTP2MaxConnections,
						Http1MaxPendingRequests:  c.HTTP1MaxPendingRequests,
						MaxRequestsPerConnection: c.MaxRequestsPerConnection,
					},
				},
				OutlierDetection: &networkingv1alpha3.OutlierDetection{
					ConsecutiveErrors: c.ConsecutiveError,
					Interval: &types.Duration{
						Seconds: c.Interval,
					},
					BaseEjectionTime: &types.Duration{
						Seconds: c.BaseEjectionTime,
					},
					MaxEjectionPercent: c.MaxEjectionPercent,
				},
			},
		},
	}
	return cb
}

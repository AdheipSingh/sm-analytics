package circuit

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"sm-analytics/models"
	"sm-analytics/utils"
)

// CreateCbr shall create a circuit breaker file.
func CreateCbr(w http.ResponseWriter, req *http.Request) {
	// Istio clientset
	ic := utils.GetIstioClientset()

	body, err := ioutil.ReadAll(io.LimitReader(req.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := req.Body.Close(); err != nil {
		panic(err)
	}
	var cbr models.CircuitBreaker

	if err := json.Unmarshal(body, &cbr); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	cb := NewCircuitbreaker(&cbr)
	// TODO:Conditional check for existence
	//  Create cb in the given namespace
	cb, err = ic.NetworkingV1alpha3().DestinationRules(cbr.Namespace).Create(cb)
	if err != nil {
		log.Printf("%d", err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cb); err != nil {
		panic(err)
	}
	log.Printf("Destination Rule created for Circuit Breaker  with the Name %s in namespace %s", cb.Name, cb.Namespace)

}

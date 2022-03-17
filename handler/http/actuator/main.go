package actuator

import (
	"encoding/json"
	"net/http"
)

//HealthBody usada para criar um status da aplicação
type HealthBody struct {
	Status string `json:"status"`
}

//Health verifica se a aplicação está de pé
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-type", "application/json")

	var status = HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(status)

}

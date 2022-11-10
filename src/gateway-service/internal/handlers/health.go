package handlers

import "net/http"

func (gs *GatewayService) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

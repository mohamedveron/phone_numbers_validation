package api

import (
	"encoding/json"
	"net/http"
)

func (s *Server) GetPhoneNumbers(w http.ResponseWriter, r *http.Request){
	employees , err := s.svc.GetPhoneNumbers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}

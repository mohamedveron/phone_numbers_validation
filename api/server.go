package api

import "github.com/mohamedveron/phone_numbers_validation/service"

type Server struct {
	svc *service.Service
}

func NewServer(svc *service.Service) *Server {
	return &Server{
		svc: svc,
	}
}

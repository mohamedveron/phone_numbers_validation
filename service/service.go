package service

import "github.com/mohamedveron/phone_numbers_validation/persistence"

type Service struct {
	persistence *persistence.Persistence
}

func NewService(persistence *persistence.Persistence) *Service {

	return &Service{
		persistence: persistence,
	}
}

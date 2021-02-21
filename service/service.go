package service

import "github.com/mohamedveron/phone_numbers_validation/persistence"

type Service struct {
	persistence *persistence.Persistence
	countriesCodeMap map[string]Country
}

func NewService(persistence *persistence.Persistence, countriesCodeMap map[string]Country) *Service {

	return &Service{
		persistence: persistence,
		countriesCodeMap: countriesCodeMap,
	}
}

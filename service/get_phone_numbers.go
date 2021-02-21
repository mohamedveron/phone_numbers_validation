package service

import "github.com/mohamedveron/phone_numbers_validation/domains"

func (s *Service) GetPhoneNumbers() ([]domains.PhoneNumber, error){

	phonesList, _ := s.persistence.GetPhoneNumbers()


	return phonesList, nil
}

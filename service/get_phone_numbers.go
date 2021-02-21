package service

import (
	"fmt"
	"github.com/mohamedveron/phone_numbers_validation/domains"
	"regexp"
)

func (s *Service) GetPhoneNumbers() ([]domains.PhoneNumber, error) {

	phonesList, _ := s.persistence.GetPhoneNumbers()

	for idx := range phonesList {
		for k := range s.countriesCodeMap {
			if match, _ := regexp.MatchString(s.countriesCodeMap[k].Regex, phonesList[idx].Number); match {
				fmt.Print(k)
			}
		}
	}

	return phonesList, nil
}

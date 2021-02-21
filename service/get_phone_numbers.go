package service

import (
	"github.com/mohamedveron/phone_numbers_validation/domains"
	"regexp"
)

func (s *Service) GetPhoneNumbers() (map[string]*domains.ResponseList, error) {

	phonesList, _ := s.persistence.GetPhoneNumbers()
	listMap := make(map[string]*domains.ResponseList)

	for idx := range phonesList {
		for k := range s.countriesCodeMap {

			match, _ := regexp.MatchString(s.countriesCodeMap[k].Regex, phonesList[idx].Number);
			flag := false

			// check if number is valid or not
			if  match {
				flag = true
			}else{

				flag = false
			}

			if _, ok := listMap[k]; !ok {

				obj := domains.ResponseList{
					Code:            s.countriesCodeMap[k].Code,
					Country:         k,
					ValidNumbers:    nil,
					NotValidNumbers: nil,
				}
				listMap[k] = &obj
			}

			// if valid number
			if flag {

				listMap[k].ValidNumbers = append(listMap[k].ValidNumbers, domains.PhoneNumber{
					CustomerName: phonesList[idx].CustomerName,
					Number:       phonesList[idx].Number,
				})

			}else{

				listMap[k].NotValidNumbers = append(listMap[k].NotValidNumbers, domains.PhoneNumber{
					CustomerName: phonesList[idx].CustomerName,
					Number:       phonesList[idx].Number,
				})
			}
		}
	}

	return listMap, nil
}

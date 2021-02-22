package service

import (
	"github.com/mohamedveron/phone_numbers_validation/domains"
	"regexp"
	"strings"
)

func (s *Service) GetPhoneNumbers() (CategoryListResponse, error) {

	phonesList, _ := s.persistence.GetPhoneNumbers()
	listMap := make(map[string]*domains.ResponseList)

	for idx := range phonesList {
		for k := range s.countriesCodeMap {

			// check if current number match this country regex
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

			}else if(!flag && strings.Contains(phonesList[idx].Number, listMap[k].Code[1:])){

				listMap[k].NotValidNumbers = append(listMap[k].NotValidNumbers, domains.PhoneNumber{
					CustomerName: phonesList[idx].CustomerName,
					Number:       phonesList[idx].Number,
				})
			}
		}
	}

	return CategoryListResponse{len(phonesList), listMap}, nil
}

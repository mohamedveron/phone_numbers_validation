package service

import (
	"github.com/mohamedveron/phone_numbers_validation/domains"
	"regexp"
	"strings"
)

func (s *Service) GetPhoneNumbers(country string, state string) (CategoryListResponse, error) {

	phonesList, _ := s.persistence.GetPhoneNumbers()
	listMap := make(map[string]*domains.ResponseList)
	count := 0

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

			// if valid number and filter not equal not valid
			if flag && state != "notvalid"{

				listMap[k].ValidNumbers = append(listMap[k].ValidNumbers, domains.PhoneNumber{
					CustomerName: phonesList[idx].CustomerName,
					Number:       phonesList[idx].Number,
				})

				count++

			}else if(!flag && strings.Contains(phonesList[idx].Number, listMap[k].Code[1:])) && state != "valid" {

				listMap[k].NotValidNumbers = append(listMap[k].NotValidNumbers, domains.PhoneNumber{
					CustomerName: phonesList[idx].CustomerName,
					Number:       phonesList[idx].Number,
				})

				count++
			}
		}
	}

	filteredListMap := make(map[string]*domains.ResponseList)
	// add country filters
	if country != "" {
		for k := range listMap {
			if k == country {
				filteredListMap[k] = listMap[k]
			}
		}
	}else{
		filteredListMap = listMap
	}

	return CategoryListResponse{count, filteredListMap}, nil
}

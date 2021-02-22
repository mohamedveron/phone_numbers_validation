package service

import "github.com/mohamedveron/phone_numbers_validation/domains"

type Country struct {
	Code  string
	Regex string
}

type CategoryListResponse struct {
	Count int      `json:"count"`
	Data  map[string]*domains.ResponseList `json:"data"`
}

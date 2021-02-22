package tests

import "github.com/mohamedveron/phone_numbers_validation/service"

func GetCountriesMap() map[string]service.Country{

	countriesCodeMap := make(map[string]service.Country)

	countriesCodeMap["Cameroon"] = service.Country{"+237", "\\(237\\)\\ ?[2368]\\d{7,8}$"}

	countriesCodeMap["Ethiopia"] = service.Country{
		"+251",
		"\\(251\\)\\ ?[1-59]\\d{8}$",
	}

	countriesCodeMap["Morocco"] = service.Country{
		"+212",
		"\\(212\\)\\ ?[5-9]\\d{8}$",
	}

	countriesCodeMap["Mozambique"] = service.Country{
		"+258",
		"\\(258\\)\\ ?[28]\\d{7,8}$",
	}

	countriesCodeMap["Uganda"] = service.Country{
		"+256",
		"\\(256\\)\\ ?\\d{9}$",
	}

	return countriesCodeMap
}

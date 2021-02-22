package tests

import (
	"github.com/mohamedveron/phone_numbers_validation/persistence"
	"github.com/mohamedveron/phone_numbers_validation/service"
	"testing"
)

func TestGetAllNumbers(t *testing.T){
	persistenceLayer := persistence.NewPersistence("./sample.db")
	serviceLayer := service.NewService(persistenceLayer, GetCountriesMap())
	_, err := serviceLayer.GetPhoneNumbers("", "")

	if err != nil {
		t.Error("error while getting phone numbers list")
	}
}

func TestGetAllValidNumbers(t *testing.T){

	persistenceLayer := persistence.NewPersistence("./sample.db")
	serviceLayer := service.NewService(persistenceLayer, GetCountriesMap())
	numbers, err := serviceLayer.GetPhoneNumbers("", "valid")

	if err != nil || numbers.Data["Cameroon"].NotValidNumbers != nil{
		t.Error("there must be only valid numbers in the response")
	}
}

func TestGetSpecificCountryNumbers(t *testing.T){

	persistenceLayer := persistence.NewPersistence("./sample.db")
	serviceLayer := service.NewService(persistenceLayer, GetCountriesMap())
	numbers, err := serviceLayer.GetPhoneNumbers("Cameroon", "")
	_, exist := numbers.Data["Morocco"]

	if err != nil || exist{
		t.Error("there must be only numbers from just Cameroon country in the response")
	}
}

func TestGetSpecificCountryValidNumbers(t *testing.T){

	persistenceLayer := persistence.NewPersistence("./sample.db")
	serviceLayer := service.NewService(persistenceLayer, GetCountriesMap())
	numbers, err := serviceLayer.GetPhoneNumbers("Cameroon", "valid")
	_, exist := numbers.Data["Morocco"]

	if err != nil || exist || numbers.Data["Cameroon"].NotValidNumbers != nil{
		t.Error("there must be only valid numbers from just Cameroon country in the response")
	}
}

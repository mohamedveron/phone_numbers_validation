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

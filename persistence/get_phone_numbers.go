package persistence

import (
	"github.com/mohamedveron/phone_numbers_validation/domains"
)

func (db *Persistence) GetPhoneNumbers() ([]domains.PhoneNumber, error) {
	phonesList := []domains.PhoneNumber{}

	rows, err := db.database.Query("SELECT name, phone FROM customer")

	if err != nil {
		return phonesList, err
	}

	var name string
	var phone string

	for rows.Next() {
		p := domains.PhoneNumber{}
		rows.Scan(&name, &phone)

		p.CustomerName = name
		p.Number = phone
		phonesList = append(phonesList, p)
	}

	return phonesList, nil
}

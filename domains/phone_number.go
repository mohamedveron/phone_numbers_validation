package domains

// PhoneNumber type
type PhoneNumber struct {
	CustomerName string `json:"customer_name"`
	Number       string `json:"number" `
}

type ResponseList struct {
	Code            string
	Country         string
	ValidNumbers    []PhoneNumber
	NotValidNumbers []PhoneNumber
}

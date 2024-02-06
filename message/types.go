package message

type Invoice struct {
	BillId    string             `json:"billId"`
	Amount    float64            `json:"amount"`
	Currency  string             `json:"currency"`
	Addresses map[string]Address `json:"addresses"`
}

type Address struct {
	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Zip    string `json:"zip"`
}

type Generic map[string]interface{}

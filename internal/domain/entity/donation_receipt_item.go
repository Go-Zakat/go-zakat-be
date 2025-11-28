package entity

type DonationReceiptItem struct {
	ID          string  `json:"id"`
	FundType    string  `json:"fundType"`
	ZakatType   *string `json:"zakatType"`
	PersonCount *string `json:"personCount"`
	Amount      string  `json:"amount"`
	RiceKG      *string `json:"riceKG"`
	Notes       string  `json:"notes"`
}

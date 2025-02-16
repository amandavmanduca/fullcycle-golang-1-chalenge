package structs

type ExchangeRate struct {
	ID  string  `json:"id" gorm:"primaryKey"`
	Bid float64 `json:"bid"`
}

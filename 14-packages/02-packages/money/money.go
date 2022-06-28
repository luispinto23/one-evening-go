package money

type Money struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

func New(amount int, currency string) Money {
	return Money{
		Amount:   amount,
		Currency: currency,
	}
}

package types



//Money - представялет собой денежную сумму в минимальных единицах (центы, копейки, дирамы и т.д)
type Money int64

// Currency - представляет собой код валюты
type Currency string


// Коды валют

const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"

)

// PAN - представялет номер карты
type PAN string


type Card struct {
	ID int
	PAN PAN
	Balance Money
	MinBalance Money
	Currency Currency
	Color string
	Name string
	Active bool
}



type Payment struct{
	ID    int
	Amount Money
}
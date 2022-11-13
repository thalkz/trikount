package models

import (
	"time"

	"github.com/thalkz/trikount/format"
)

type Expense struct {
	Id        int
	Title     string
	Amount    float64
	PaidBy    Member
	SpentBy   []Member
	UpdatedAt time.Time
}

func (e Expense) AmountPerMember() float64 {
	return float64(e.Amount) / float64(len(e.SpentBy))
}

func (e Expense) FormattedAmountPerMember() string {
	return format.ToEuro(e.AmountPerMember())
}

func (e Expense) FormattedAmount() string {
	return format.ToEuro(e.Amount)
}

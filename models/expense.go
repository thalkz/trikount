package models

import (
	"fmt"
	"time"
)

type Expense struct {
	Id        int
	Title     string
	Amount    float32
	PaidBy    Member
	SpentBy   []Member
	UpdatedAt time.Time
}

func (e Expense) AmountPerMember() float64 {
	return float64(e.Amount) / float64(len(e.SpentBy))
}

func (e Expense) FormattedAmountPerMember() string {
	return fmt.Sprintf("%.2f€", e.AmountPerMember())
}

func (e Expense) FormattedAmount() string {
	return fmt.Sprintf("%.2f€", e.Amount)
}

package models

import (
	"time"

	"github.com/thalkz/trikount/format"
	"github.com/xeonx/timeago"
)

type Expense struct {
	Id         int
	Title      string
	Amount     float64
	PaidBy     Member
	SpentBy    []*Member
	CreatedAt  time.Time
	IsTransfer bool
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

func (e Expense) FormattedTimeAgo() string {
	return timeago.French.Format(e.UpdatedAt)
}

func (e Expense) HasSpent(id int) bool {
	for _, member := range e.SpentBy {
		if member.Id == id {
			return true
		}
	}
	return false
}

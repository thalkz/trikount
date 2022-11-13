package models

import (
	"fmt"

	"github.com/thalkz/trikount/format"
)

type Transfer struct {
	Amount float64
	From   Member
	To     Member
}

func (t Transfer) FormattedAmount() string {
	return format.ToEuro(t.Amount)
}

func (t Transfer) String() string {
	return fmt.Sprintf("%v -> %v %v", t.From, t.To, format.ToEuro(t.Amount))
}

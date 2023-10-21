package models

import "github.com/thalkz/trikount/format"

type ExpenseWithPart struct {
	Expense
	Part float64
}

func (e ExpenseWithPart) FormattedPart() string {
	return format.ToEuro(e.Part)
}

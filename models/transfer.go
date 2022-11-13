package models

import "fmt"

type Transfer struct {
	Amount float64
	From   Member
	To     Member
}

func (t Transfer) String() string {
	return fmt.Sprintf("%v -> %v %.2f€", t.From, t.To, t.Amount)
}

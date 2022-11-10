package models

import "time"

type Expense struct {
	Id        int
	Title     string
	Amount    float32
	PaidBy    Member
	SpentBy   []Member
	CreatedAt time.Time
	UpdatedAt time.Time
}

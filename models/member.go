package models

type Member struct {
	Id         int
	Name       string
	TotalPaid  float32 // Computed value
	TotalSpent float32 // Computed value
}

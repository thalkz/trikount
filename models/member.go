package models

type Member struct {
	Id   int
	Name string
}

type MemberBalance struct {
	Member
	Spent float64
	Paid  float64
}

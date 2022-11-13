package models

type MemberBalance struct {
	Member
	Spent float64
	Paid  float64
}

func (t MemberBalance) String() string {
	return t.Name
}

func (b MemberBalance) Balance() float64 {
	return b.Paid - b.Spent
}

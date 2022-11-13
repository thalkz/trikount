package models

type Member struct {
	Id   int
	Name string
}

func (t Member) String() string {
	return t.Name
}

package models

import "time"

type Project struct {
	Id         string
	Name       string
	CreatedAt  time.Time
	TotalSpent float32 // Computed value
}

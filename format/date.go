package format

import (
	"fmt"
	"time"
)

var months = []string{
	"janvier",
	"février",
	"mars",
	"avril",
	"mai",
	"juin",
	"juillet",
	"août",
	"septembre",
	"octobre",
	"novembre",
	"décembre",
}

func FormatDateFrench(t time.Time) string {
	t = t.Add(time.Second) // Add one second to prevent midnight error
	year, month, day := time.Now().Date()
	today := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	if t.After(today) {
		return "aujourd'hui"
	} else if t.After(today.AddDate(0, 0, -1)) {
		return "hier"
	} else if year == t.Year() {
		return fmt.Sprintf("%v %v",
			t.Day(),
			months[t.Month()-1],
		)
	} else {
		return fmt.Sprintf("%v %v %v",
			t.Day(),
			months[t.Month()-1],
			t.Year(),
		)
	}
}

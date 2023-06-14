package format

import "fmt"

func ToEuro(amount float64) string {
	return fmt.Sprintf("%.2f€", amount)
}

func ToSignedEuro(amount float64) string {
	if amount > -0.01 && amount < 0.01 {
		return "+0.00€"
	}
	return fmt.Sprintf("%+.2f€", amount)
}

package format

import "fmt"

func ToEuro(amount float64) string {
	return fmt.Sprintf("%.2f€", amount)
}

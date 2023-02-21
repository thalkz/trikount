package models

import (
	"math"

	"github.com/thalkz/trikount/format"
)

type Balance struct {
	Members []*MemberBalance
}

func (b Balance) TotalSpent() float64 {
	var total float64
	for _, member := range b.Members {
		total += member.Spent
	}
	return total
}

func (b Balance) FormattedTotalSpent() string {
	return format.ToEuro(b.TotalSpent())
}

func (b Balance) GetTransfers() []*Transfer {
	transfers := make([]*Transfer, 0)
	balances := make(map[*Member]float64)
	for _, member := range b.Members {
		balances[&member.Member] = member.Balance()
	}

	var topPayer *Member
	var topPayerAmount float64
	var topSpender *Member
	var topSpenderAmount float64

	for {
		topPayerAmount = 0
		topSpenderAmount = 0
		for member, amount := range balances {
			if amount < topSpenderAmount {
				topSpender = member
				topSpenderAmount = amount
			}
			if amount > topPayerAmount {
				topPayer = member
				topPayerAmount = amount
			}
		}
		transferAmount := math.Min(topPayerAmount, -topSpenderAmount)

		if transferAmount < 0.01 {
			break
		}

		transfers = append(transfers, &Transfer{
			From:   *topSpender,
			To:     *topPayer,
			Amount: transferAmount,
		})
		balances[topPayer] -= transferAmount
		balances[topSpender] += transferAmount
	}
	return transfers
}

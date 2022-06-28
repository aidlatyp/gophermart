package core

import (
	"errors"
	"time"
)

type withdraw struct {
	OrderNumber int
	Amount      int
	time        int64
}

type Account struct {
	Id              string
	User            string //owner
	Points          int
	WithdrawHistory []withdraw
}

func NewAccount(user string) Account {
	return Account{
		User: user,
	}
}

func (a *Account) CurrentPoints() int {
	return a.Points
}

func (a *Account) AddPoints() { /* calculations.. */ }

func (a *Account) WithdrawPoints(order int, amount int) error {
	if amount > a.Points {
		return errors.New("not enough funds")
	}
	w := withdraw{
		OrderNumber: order,
		Amount:      amount,
		time:        time.Now().Unix(),
	}
	a.Points = -amount
	a.WithdrawHistory = append(a.WithdrawHistory, w)
	return nil
}

// ... other funcs

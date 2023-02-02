package accounts

import "errors"

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("you can't withdraw money")

func NewAccount(owner string, balance int) *Account {
	account := Account{
		owner:   owner,
		balance: balance,
	}
	return &account
}

func (a *Account) Deposit(amount int) {
	(*a).balance += amount
}

func (a *Account) GetBalance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(owner string) {
	a.owner = owner
}

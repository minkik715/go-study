package main

import (
	"fmt"
	"learngo/accounts"
	"log"
)

func main() {
	account := accounts.NewAccount("minki", 9000000)
	account.Deposit(300000)
	err := account.Withdraw(1000000000)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.GetBalance())
}

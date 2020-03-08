package main

import (
	"fmt"

	"github.com/changhoi/learngo/account"
)

type person struct {
	name         string
	age          int
	favoriteFood []string
}

func main() {
	account := account.NewAccount("changhoi")
	account.Deposit(5000)
	fmt.Println(account.Balance())
	err := account.Withdraw(6000)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(account.Balance())
	fmt.Println(account)
}

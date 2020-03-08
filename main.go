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
	account.Withdraw(4000)
	fmt.Println(account.Balance())
}

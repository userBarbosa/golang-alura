package main

import (
	"fmt"

	"github.com/userbarbosa/golang-alura/orientacao_a_objetos/project/v2/accounts"
)

func PayBill(account validateAccountType, billAmount float64) {
	account.Withdraw(billAmount)
}

type validateAccountType interface {
	Withdraw(amount float64) string
}

func main() {
	TestAccountOne := accounts.SavingsAccount{}
	TestAccountOne.Deposit(1000)
	PayBill(&TestAccountOne, 100)

	fmt.Println(TestAccountOne.GetBalance())

	TestAccountTwo := accounts.CheckingAccount{}
	TestAccountTwo.Deposit(100)
	PayBill(&TestAccountTwo, 200)

	fmt.Println(TestAccountTwo.GetBalance())
}

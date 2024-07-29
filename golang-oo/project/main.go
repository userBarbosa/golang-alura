package main

import (
	"fmt"

	"github.com/userbarbosa/golang-alura/golang-oo/project/v2/accounts"
)

func PayBill(account validateAccountType, billAmount float64) {
	account.Withdraw(billAmount)
}

type validateAccountType interface {
	Withdraw(amount float64) string
}

func main() {
	fmt.Println("Test case: instantiate a account with $1000 and pay a $100 bill:")
	TestAccountOne := accounts.SavingsAccount{}
	TestAccountOne.Deposit(1000)
	PayBill(&TestAccountOne, 100)
	fmt.Printf("First account after paying bill: %f\n\n", TestAccountOne.GetBalance())

	fmt.Println("Test case: instantiate a second account with $100 and try to pay a $200 bill:")
	TestAccountTwo := accounts.CheckingAccount{}
	TestAccountTwo.Deposit(100)
	PayBill(&TestAccountTwo, 200)
	fmt.Printf("Second account after paying bill: %f\n\n", TestAccountTwo.GetBalance())

	fmt.Println("Test case: instantiate a third account and transfer $50 from the second account:")
	TestAccountThree := accounts.CheckingAccount{}
	TestAccountTwo.Transfer(50, &TestAccountThree)
	fmt.Printf("Second account after transferring values: %f\n", TestAccountTwo.GetBalance())
	fmt.Printf("Third account after receiving values: %f\n", TestAccountThree.GetBalance())
}

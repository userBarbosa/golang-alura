package accounts

import "github.com/userbarbosa/golang-alura/golang-oo/project/v2/clients"

type SavingsAccount struct {
	Holder                                       clients.Holder
	AccountNumber, AgencyNumber, OperationNumber int
	balance                                      float64
}

func (ca *SavingsAccount) Withdraw(amount float64) string {
	validValue := amount > 0 && amount <= ca.balance
	if validValue {
		ca.balance -= amount
		return "Withdrawal successful"
	} else {
		return "Insufficient balance"
	}
}

func (ca *SavingsAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		ca.balance += amount
		return "Deposit successful", ca.balance
	} else {
		return "Invalid deposit amount", ca.balance
	}
}

func (ca *SavingsAccount) GetBalance() float64 {
	return ca.balance
}

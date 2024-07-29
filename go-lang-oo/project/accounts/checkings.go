package accounts

import "github.com/userbarbosa/golang-alura/orientacao_a_objetos/project/v2/clients"

type CheckingAccount struct {
	Holder                      clients.Holder
	AccountNumber, AgencyNumber int
	balance                     float64
}

func (ca *CheckingAccount) Withdraw(amount float64) string {
	validValue := amount > 0 && amount <= ca.balance
	if validValue {
		ca.balance -= amount
		return "Withdrawal successful"
	} else {
		return "Insufficient balance"
	}
}

func (ca *CheckingAccount) Deposit(amount float64) (string, float64) {
	if amount > 0 {
		ca.balance += amount
		return "Deposit successful", ca.balance
	} else {
		return "Invalid deposit amount", ca.balance
	}
}

func (ca *CheckingAccount) Transfer(amount float64, destinationAccount *CheckingAccount) bool {
	if amount > 0 && amount < ca.balance {
		ca.balance -= amount
		destinationAccount.Deposit(amount)
		return true
	} else {
		return false
	}
}

func (ca *CheckingAccount) GetBalance() float64 {
	return ca.balance
}

package main

import "fmt"

type InsufficientBalanceError struct {
	CurrentBalance  float64
	RequestedAmount float64
}

func (e *InsufficientBalanceError) Error() string {
	return fmt.Sprintf("saldo insuficiente: saldo atual é %.2f, mas o valor requerido é %.2f", e.CurrentBalance, e.RequestedAmount)
}

func withdraw(balance, amount float64) error {
	if amount > balance {
		return &InsufficientBalanceError{CurrentBalance: balance, RequestedAmount: amount}
	}
	return nil
}

func main() {
	err := withdraw(100, 150)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Saque permitido")
	}

}

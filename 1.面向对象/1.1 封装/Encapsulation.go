package main

import (
	"errors"
	"fmt"
	"time"
)

// Wallet wallet struct
type Wallet struct {
	id               int
	balance          float64
	lastModifiedTime time.Time
}

// GetWalletBalance get current balance
func (w *Wallet) GetWalletBalance() float64 {
	return w.balance
}

// IncreaseBalance get wallet id
func (w *Wallet) IncreaseBalance(increaseAmount float64) error {
	if increaseAmount < 0 {
		return errors.New("Invalid increase amount ")
	}
	w.lastModifiedTime = time.Now()
	w.balance = w.balance + increaseAmount
	return nil
}

// DecreaseBalance get wallet id
func (w *Wallet) DecreaseBalance(decreaseAmount float64) error {
	if decreaseAmount < 0 {
		return errors.New("Invalid decrease amount ")
	}
	if w.balance < decreaseAmount {
		return errors.New("Insufficient amount ")
	}
	w.lastModifiedTime = time.Now()
	w.balance = w.balance - decreaseAmount
	return nil
}

// Transfer to wallet
func (w *Wallet) Transfer(toWallet *Wallet, transferAmount float64) error {
	if transferAmount < 0 {
		return errors.New("Invalid transfer amount ")
	}
	if w.balance < transferAmount {
		return errors.New("Insufficient amount ")
	}
	transferTime := time.Now()
	w.lastModifiedTime = transferTime
	toWallet.lastModifiedTime = transferTime
	w.balance = w.balance - transferAmount
	toWallet.balance = toWallet.balance + transferAmount
	return nil
}

func main() {
	fmt.Println("123")
}

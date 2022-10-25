package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Money int
}

func (w *Wallet) Pay(amount int) error {
	if w.Money < amount {
		return fmt.Errorf("Недостаточно средств")
	} else {
		w.Money -= amount
	}
	return nil
}

type ApplePay struct {
	Account int
}

func (a *ApplePay) Pay(amount int) error {
	if a.Account < amount {
		return fmt.Errorf("Недостаточно средств")
	} else {
		a.Account -= amount
	}
	return nil

}

type Cash struct {
	Hm int
}

func (c *Cash) Pay(amount int) error {
	if c.Hm < amount {
		return fmt.Errorf("Недостаточно средств")
	} else {
		c.Hm -= amount
	}
	return nil
}
func Buy(p Payer) {
	err := p.Pay(100)
	if err != nil {
		panic(err)
	}
	fmt.Println("Спасибо за покупку!!!")
}

func main() {
	myWallet := Wallet{Money: 1000}
	myApplePay := ApplePay{Account: 10}
	myCash := Cash{Hm: 2000}
	Buy(&myWallet)
	Buy(&myApplePay)
	Buy(&myCash)

}

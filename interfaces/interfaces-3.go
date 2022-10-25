package main

import "fmt"

type Wallet struct {
	Money int
}

func (w *Wallet) Pay(amount int) error {
	if w.Money < amount {
		return fmt.Errorf("Недостаточно средств в кошельке")
	}
	w.Money -= amount
	return nil
}

type Card struct {
	Balance int
	Owner   string
	CVV     string
	Number  string
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return fmt.Errorf("%v,Недостаточно средств в кошельке", c.Owner)
	}
	c.Balance -= amount
	return nil

}

type Payer interface {
	Pay(int) error
}

func Buy(p Payer) {
	err := p.Pay(500)
	if err != nil {
		panic(err)
	}
	fmt.Println("Спасибо за покупку!!!")
}

func main() {
	//myWallet := Wallet{Money: 100}
	myCard := Card{Balance: 5000, Owner: "Kitri"}
	// defer Buy(&myWallet)
	Buy(&myCard)

	var person Payer
	person = &Card{Balance: 300, Owner: "Rigard"}
	Buy(person)
}

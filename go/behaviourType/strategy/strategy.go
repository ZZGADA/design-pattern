package strategy

import "fmt"

type strategy interface {
	Pay(amount int)
}

type CashPayStrategy struct{}

func (cashPayStrategy *CashPayStrategy) Pay(amount int) {
	fmt.Println("using cash to pay , the amount multi 10  is ", amount*10)
}

type CreditPayStrategy struct{}

func (creditPayStrategy *CreditPayStrategy) Pay(amount int) {
	fmt.Println("using credit cart to pay , the amount plus 10 is ", amount+10)
}

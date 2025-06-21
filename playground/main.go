package main

import "fmt"

type CartItem struct {
	Code     string
	Quantity int
}

type Cart struct {
	Items []CartItem
}

type InventoryItem struct {
	Code      string
	UnitPrice float64
}

var inventory = []InventoryItem{
	{
		Code:      "P001",
		UnitPrice: 10.0,
	},
	{
		Code:      "P002",
		UnitPrice: 10.0,
	},
	{
		Code:      "P003",
		UnitPrice: 10.0,
	},
	{
		Code:      "P004",
		UnitPrice: 10.0,
	},
	{
		Code:      "P005",
		UnitPrice: 10.0,
	},
	{
		Code:      "P006",
		UnitPrice: 10.0,
	},
	{
		Code:      "P007",
		UnitPrice: 10.0,
	},
	{
		Code:      "P008",
		UnitPrice: 10.0,
	},
	{
		Code:      "P009",
		UnitPrice: 10.0,
	},
	{
		Code:      "P010",
		UnitPrice: 10.0,
	},
}

type ItemMap map[string]InventoryItem

type PaymentProcessor struct {
	// implementar atributos necesarios aquí
	PaymentMethod PaymentMethod
}

type PaymentMethod interface {
	Pay() error
	GetType() string
}

type CreditCard struct {
	CardNumber  string
	ExpiryYear  string
	ExpiryMonth string
	CVV         string
}

func (c *CreditCard) Pay() error {
	return nil
}

func (c *CreditCard) GetType() string {
	return "Credit Card"
}

type Paypal struct {
	Email string
}

func (p *Paypal) Pay() error {
	return nil
}

func (p *Paypal) GetType() string {
	return "Paypal"
}

type CryptoMoney struct {
	WalletAddress string
}

func (c *CryptoMoney) Pay() error {
	return nil
}

func (c *CryptoMoney) GetType() string {
	return "CryptoMoney"
}

func (p *PaymentProcessor) ProcessPayment(amount float64) {
	// implementar lógica de la función aquí
	if p.PaymentMethod == nil {
		fmt.Println("Empty payment method")
		return
	}
	fmt.Println("Payment type used: ", p.PaymentMethod.GetType())
}

func main() {

	itemM := ItemMap{}
	for _, item := range inventory {
		itemM[item.Code] = item
	}

	cart := Cart{
		Items: []CartItem{
			{Code: "P001", Quantity: 2},
			{Code: "P010", Quantity: 1},
		},
	}

	// PUNTOS A EVALUAR
	// 1. Calcular el total del carrito en una variable llamada totalAmountToPay
	totalAmountToPay := 0.0
	for _, item := range cart.Items {
		i, ok := itemM[item.Code]
		if !ok {
			fmt.Println("Item with code not found", item.Code)
		} else {
			totalAmountToPay += i.UnitPrice * float64(item.Quantity)
		}
	}
	fmt.Println(totalAmountToPay)

	//creditCard := &CreditCard{
	//	CardNumber:  "4242424242424242",
	//	ExpiryYear:  "10",
	//	ExpiryMonth: "12",
	//	CVV:         "123",
	//}
	//paypal := &Paypal{
	//	Email: "email@user.com",
	//}
	crypto := &CryptoMoney{
		WalletAddress: "address",
	}
	paymentProcessor := PaymentProcessor{
		PaymentMethod: crypto,
	}

	// 2. Utilizar ProcessPayment(totalAmountToPay) e imprimir en pantalla el método de pago iutilizado
	paymentProcessor.ProcessPayment(totalAmountToPay)

}

/*
CreditCard
*/

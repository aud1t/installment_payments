package main

import (
	"fmt"
	"installment_payments/pkg/repo"
	"log"
)

func main() {
	apple := repo.Product{
		Category: repo.Smartphone,
		Price:    1000,
	}

	client := repo.Client{
		PhoneNum: "+7-977-961-6505",
	}

	credit := repo.Credit{
		Product: apple,
		Client:  client,
	}

	fullPrice, err := repo.FullPrice(apple, apple.Price, "+7-977-961-65-05", 18)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Printf("Польная стоимость продукта %d\n", fullPrice)

	credit.HowMuchLeftPay = fullPrice
	credit.Pay(600)
}

package repo

import (
	"errors"
	"fmt"
)

type Credit struct {
	Product
	Client
	HowMuchLeftPay Money
}

type Client struct {
	PhoneNum string
}

// Pay метод для оплаты кредита
func (c *Credit) Pay(amount Money) error {

	if amount < 0 {
		return errors.New("сумма оплаты должна быть положительной")
	}

	if c.HowMuchLeftPay < amount {
		return errors.New("сумма оплаты не может быть больше остатка")
	}

	c.HowMuchLeftPay -= amount

	sms := fmt.Sprintf("Вы заплатили %d сомони на продукт категории %s. Осталось заплатить еще %d сомони",
		amount, c.Category, c.HowMuchLeftPay)

	c.Client.SendSMS(sms)

	return nil
}

// SendSMS метод для отправки смс на номер телефона клиента
func (c *Client) SendSMS(message string) {
	fmt.Print(message)
}

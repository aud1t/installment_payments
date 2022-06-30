package repo

import "errors"

type Money int

type Categories string

const (
	Smartphone Categories = "Smartphone"
	Computer   Categories = "Computer"
	TV         Categories = "TV"
)

const (
	SmartphonePercent = 3 // 3% на каждую дополнительную единицу диапазона для смартфона
	ComputerPercent   = 4 // 4% на каждую дополнительную единицу диапазона для компьютера
	TVPercent         = 5 // 5% на каждую дополнительную единицу диапазона для ТВ
)

// Product тип продукта
type Product struct {
	Price    Money
	Category Categories
}

// FullPrice функция вычисляющая польную стоимость продукта
func FullPrice(product Product, price Money, phoneNumber string, diapason int) (Money, error) {
	if price < 500 || price > 100_000 {
		return 0, errors.New("Стоимость продукта должна быть больше 500 и меньше 100000 сомони")
	}
	return countPrice(product.Category, price, diapason)
}

func countPrice(category Categories, price Money, diapason int) (Money, error) {
	switch category {
	case Smartphone:
		return priceForSmartphone(price, SmartphonePercent, diapason)
	case Computer:
		return priceForComputer(price, ComputerPercent, diapason)
	case TV:
		return priceForTV(price, TVPercent, diapason)
	default:
		return 0, errors.New("Не существующая категория")
	}
}

func priceForSmartphone(price Money, percent int, diapason int) (Money, error) {
	switch {
	case diapason >= 3 && diapason <= 9:
		return price, nil
	case diapason > 9 && diapason <= 12:
		return price + Money(Money(percent)*price/100), nil
	case diapason > 12 && diapason <= 18:
		return price + Money(Money(percent)*2*price/100), nil
	case diapason > 18 && diapason <= 24:
		return price + Money(Money(percent)*3*price/100), nil
	default:
		return 0, errors.New("Диапазон рассрочки 3-6-9-12-18-24 месяцев")

	}
}

func priceForComputer(price Money, percent int, diapason int) (Money, error) {
	switch {
	case diapason >= 3 && diapason <= 12:
		return price, nil
	case diapason > 12 && diapason <= 18:
		return price + Money(Money(percent)*price/100), nil
	case diapason > 18 && diapason <= 24:
		return price + Money(Money(percent)*2*price/100), nil
	default:
		return 0, errors.New("Диапазон рассрочки 3-6-9-12-18-24 месяцев")
	}
}

func priceForTV(price Money, percent int, diapason int) (Money, error) {
	switch {
	case diapason >= 3 && diapason <= 18:
		return price, nil
	case diapason > 18 && diapason <= 24:
		return price + Money(Money(percent)*price/100), nil
	default:
		return 0, errors.New("Диапазон рассрочки 3-6-9-12-18-24 месяцев")
	}
}

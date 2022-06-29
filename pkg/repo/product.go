package repo

type Money int

type Categories string

const Smartphone Categories = "Smartphone"
const Computer Categories = "Computer"
const TV Categories = "TV"

type Product struct {
	Price    Money
	Category Categories
}

const SmartphonePercent = 3
const ComputerPercent = 4
const TVPercent = 5

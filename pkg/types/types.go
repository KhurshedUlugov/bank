package types
// Money денежная сумма
type Money int64
// Category категории
type Category string
// Payment информация о платеже
type Payment struct {
	ID int
	Amount Money
	Category Category
}
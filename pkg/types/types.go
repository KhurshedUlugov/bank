package types
// Money денежная сумма
type Money int64
// Category категории
type Category string

type Status string

const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Payment информация о платеже
type Payment struct {
	ID int
	Amount Money
	Category Category
	Status Status
}
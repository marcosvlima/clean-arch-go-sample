package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	FindAll(page, limit int) ([]Order, error)
	GetTotal() (int, error)
}

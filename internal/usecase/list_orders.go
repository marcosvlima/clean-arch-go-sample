package usecase

import (
	"github.com/marcosvlima/clean-arch-go-sample/internal/entity"
	"github.com/marcosvlima/clean-arch-go-sample/pkg/events"
)

type OrderListOutputDTO struct {
	Orders []OrderOutputDTO `json:"orders"`
	Total  int              `json:"total"`
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
	OrderListed     events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
	OrderListed events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *ListOrdersUseCase {

	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
		OrderListed:     OrderListed,
		EventDispatcher: EventDispatcher,
	}
}

func (l *ListOrdersUseCase) Execute(page, limit int) (OrderListOutputDTO, error) {
	orders, err := l.OrderRepository.FindAll(page, limit)
	if err != nil {
		return OrderListOutputDTO{}, err
	}

	total, err := l.OrderRepository.GetTotal()
	if err != nil {
		return OrderListOutputDTO{}, err
	}
	var output []OrderOutputDTO
	for _, order := range orders {
		output = append(output, OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return OrderListOutputDTO{
		Orders: output,
		Total:  total,
	}, nil
}

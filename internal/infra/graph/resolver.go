package graph

import "github.com/marcosvlima/clean-arch-go-sample/internal/usecase"

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

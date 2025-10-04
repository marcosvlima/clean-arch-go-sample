package main

import (
	"database/sql"

	"github.com/google/wire"
	"github.com/marcosvlima/clean-arch-go-sample/internal/entity"
	"github.com/marcosvlima/clean-arch-go-sample/internal/event"
	"github.com/marcosvlima/clean-arch-go-sample/internal/infra/database"
	"github.com/marcosvlima/clean-arch-go-sample/internal/infra/web"
	"github.com/marcosvlima/clean-arch-go-sample/internal/usecase"
	"github.com/marcosvlima/clean-arch-go-sample/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func NewCreateOrderUseCase(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func NewWebOrderHandler(db *sql.DB, eventDispatcher events.EventDispatcherInterface) *web.WebOrderHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.WebOrderHandler{}
}

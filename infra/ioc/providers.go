package ioc

import (
	eventemitter "github.com/melkdesousa/event-emitter/infra/event_emitter"
	"github.com/melkdesousa/event-emitter/usecases"
)

func InitializeUsecases() (usecases.CreateOrder, usecases.PublishOrder, usecases.UpdateStatusOrder) {
	createOrder := usecases.NewCreateOrder(eventemitter.EE)
	publishOrder := usecases.NewPublishOrder(eventemitter.EE)
	updateStatusOrder := usecases.NewUpdateStatusOrder()
	return *createOrder, *publishOrder, *updateStatusOrder
}

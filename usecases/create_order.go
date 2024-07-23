package usecases

import (
	"log"

	eventemitter "github.com/melkdesousa/event-emitter/infra/event_emitter"
)

type CreateOrder struct {
	eventEmitter *eventemitter.EventEmitter
}

func NewCreateOrder(eventEmitter *eventemitter.EventEmitter) *CreateOrder {
	return &CreateOrder{eventEmitter: eventEmitter}
}

func (uc *CreateOrder) Execute() {
	orderId := "123"

	// Business logic
	// ...
	log.Printf("Order %s created\n", orderId)
	uc.eventEmitter.Emit(eventemitter.CreatedOrder, orderId)
}

package usecases

import (
	"log"

	eventemitter "github.com/melkdesousa/event-emitter/infra/event_emitter"
)

type PublishOrder struct {
	eventEmitter *eventemitter.EventEmitter
}

func NewPublishOrder(eventEmitter *eventemitter.EventEmitter) *PublishOrder {
	return &PublishOrder{eventEmitter: eventEmitter}
}

func (uc *PublishOrder) Execute(orderId string) {
	// Business logic
	// ...
	log.Printf("Order %s published\n", orderId)
	uc.eventEmitter.Emit(eventemitter.PublishedOrder, orderId, "published")
}

func (uc *PublishOrder) Wrapper(messages ...interface{}) {
	orderId := messages[0].(string)

	uc.Execute(orderId)
}

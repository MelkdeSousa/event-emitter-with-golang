package main

import (
	eventemitter "github.com/melkdesousa/event-emitter/infra/event_emitter"
	"github.com/melkdesousa/event-emitter/infra/ioc"
)

func main() {
	createOrder, publishOrder, updateStatusOrder := ioc.InitializeUsecases()
	eventemitter.EE.On(eventemitter.CreatedOrder, publishOrder.Wrapper)
	eventemitter.EE.On(eventemitter.PublishedOrder, updateStatusOrder.Wrapper)

	createOrder.Execute()
}

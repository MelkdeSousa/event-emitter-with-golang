package usecases

import "fmt"

type UpdateStatusOrder struct{}

func NewUpdateStatusOrder() *UpdateStatusOrder {
	return &UpdateStatusOrder{}
}

func (uc *UpdateStatusOrder) Execute(orderId string, status string) {
	// Business logic
	// ...
	fmt.Printf("Order %s updated to status %s\n", orderId, status)
}

func (uc *UpdateStatusOrder) Wrapper(messages ...interface{}) {
	orderId := messages[0].(string)
	status := messages[1].(string)

	uc.Execute(orderId, status)
}

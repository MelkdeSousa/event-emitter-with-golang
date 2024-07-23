package eventemitter

type EventToken string

const (
	CreatedOrder   EventToken = "CreatedOrder"
	UpdatedOrder   EventToken = "UpdatedOrder"
	DeletedOrder   EventToken = "DeletedOrder"
	PublishedOrder EventToken = "PublishedOrder"
)

type EventEmitter struct {
	events map[EventToken][]func(messages ...interface{})
}

func newEventEmitter() *EventEmitter {
	return &EventEmitter{events: make(map[EventToken][]func(messages ...interface{}))}
}

// On registers a listener function for the specified event.
// The listener function will be called whenever the event is emitted.
func (emitter *EventEmitter) On(event EventToken, listener func(messages ...interface{})) {
	emitter.events[event] = append(emitter.events[event], listener)
}

// Emit sends the specified event with the given messages to all registered listeners.
// It iterates over the list of listeners for the specified event and invokes each listener
// with the provided messages as arguments.
//
// Parameters:
// - event: The event token representing the event to be emitted.
// - messages: The messages to be passed to the listeners.
//
// Example:
//
//	emitter.Emit("user.created", "John Doe", 30)
//
// This will emit the "user.created" event with the messages "John Doe" and 30 to all
// registered listeners for the "user.created" event.
func (emitter *EventEmitter) Emit(event EventToken, messages ...interface{}) {
	for _, listener := range emitter.events[event] {
		listener(messages...)
	}
}

var EE *EventEmitter

func init() {
	EE = newEventEmitter()
}

package order

// ConfirmOrderRequest represents the body payload for confirming an order.
// This operation transitions an order from pending to a confirmed state.
type ConfirmOrderRequest struct {
	// ID is the unique identifier of the order to confirm.
	ID string `json:"id,omitempty"`
	// Status is typically set to "confirmed" when confirming an order.
	Status string `json:"status,omitempty"`
}

// SimulateOrderEventRequest represents the body payload for simulating an event
// on an order (e.g., for testing webhooks or state transitions).
type SimulateOrderEventRequest struct {
	// EventType is the type of event to simulate (e.g., "payment.updated", "order.completed").
	EventType string `json:"event_type,omitempty"`
	// Data contains additional event-specific data.
	Data map[string]interface{} `json:"data,omitempty"`
}
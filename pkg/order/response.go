package order

import "time"

// ConfirmOrderResponse represents the response returned after confirming an order.
// It contains the updated order details following the confirmation operation.
type ConfirmOrderResponse struct {
	// ID is the unique identifier for the confirmed order.
	ID string `json:"id"`
	// Status is the current status of the order (e.g., "confirmed", "processed").
	Status string `json:"status"`
	// TotalAmount is the total amount of the order.
	TotalAmount string `json:"total_amount"`
	// Currency is the ISO 4217 currency code (e.g., "BRL", "USD").
	Currency string `json:"currency"`
	// ExternalReference is the integrator-provided identifier for reconciliation.
	ExternalReference string `json:"external_reference"`
	// DateCreated is the date when the order was created.
	DateCreated time.Time `json:"date_created"`
	// DateLastUpdated is the date when the order was last modified.
	DateLastUpdated time.Time `json:"date_last_updated"`
}

// SimulateOrderEventResponse represents the response returned after simulating an event.
// It contains metadata about the simulated event and its outcome.
type SimulateOrderEventResponse struct {
	// EventID is the unique identifier for the simulated event.
	EventID string `json:"event_id"`
	// EventType is the type of event that was simulated.
	EventType string `json:"event_type"`
	// Status indicates whether the simulation was successful.
	Status string `json:"status"`
	// Message provides additional information about the simulation result.
	Message string `json:"message"`
	// DateCreated is the date when the event simulation occurred.
	DateCreated time.Time `json:"date_created"`
}

// GetOrderRefundsResponse represents the response containing all refunds for an order.
// It includes a list of individual refund transactions and summary information.
type GetOrderRefundsResponse struct {
	// Refunds is the list of refund transactions associated with the order.
	Refunds []OrderRefund `json:"refunds"`
	// TotalRefunded is the total amount refunded across all transactions.
	TotalRefunded string `json:"total_refunded"`
	// Currency is the ISO 4217 currency code for the refund amounts.
	Currency string `json:"currency"`
}

// OrderRefund represents a single refund transaction within an order.
type OrderRefund struct {
	// ID is the unique identifier for the refund.
	ID string `json:"id"`
	// TransactionID is the identifier of the transaction being refunded.
	TransactionID string `json:"transaction_id"`
	// Amount is the amount refunded in this transaction.
	Amount string `json:"amount"`
	// Status is the current status of the refund (e.g., "approved", "pending", "rejected").
	Status string `json:"status"`
	// Reason is the reason for the refund, if provided.
	Reason string `json:"reason"`
	// DateCreated is the date when the refund was initiated.
	DateCreated time.Time `json:"date_created"`
	// DateApproved is the date when the refund was approved, if applicable.
	DateApproved *time.Time `json:"date_approved,omitempty"`
}
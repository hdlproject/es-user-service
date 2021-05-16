package entity

type (
	TopUpEvent struct {
		TransactionEventID string
		UserID             uint
		Amount             uint64
	}
)

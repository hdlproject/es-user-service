package input_port

type (
	TopUpRequest struct {
		TransactionEventID string
		UserID             uint
		Increment          uint64
	}

	TopUpResponse struct {
		Ok      bool
		Message string
	}
)

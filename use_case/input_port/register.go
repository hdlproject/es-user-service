package input_port

type (
	RegisterRequest struct {
		Username string
		Password string
	}

	RegisterResponse struct {
		Ok      bool
		Message string
		UserID  uint
	}
)

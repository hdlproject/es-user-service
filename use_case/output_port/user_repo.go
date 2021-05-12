package output_port

import "github.com/hdlproject/es-user-service/entity"

type (
	UserRepo interface {
		Register(user entity.User) (uint, error)
		IncreaseBalance(id uint, increment uint64) error
	}
)

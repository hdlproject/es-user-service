package output_port

import "github.com/hdlproject/es-user-service/entity"

type (
	TopUpEventRepo interface {
		Insert(event entity.TopUpEvent) (string, error)
		IsAlreadyApplied(event entity.TopUpEvent) (bool, error)
	}
)

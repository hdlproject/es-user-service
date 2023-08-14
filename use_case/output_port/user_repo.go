package output_port

import (
	"context"

	"github.com/hdlproject/es-user-service/entity"
)

type (
	UserRepo interface {
		Register(user entity.User) (uint, error)
		IncreaseBalance(id uint, increment uint64) error
		Track(ctx context.Context, location entity.UserLocation) error
		LocateNearbyUser(ctx context.Context, location entity.UserLocation) ([]uint, error)
	}
)

package interactor

import (
	"context"

	"golang.org/x/exp/slices"

	"github.com/hdlproject/es-user-service/entity"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/output_port"
)

type (
	TrackRequest struct {
		UserID uint
		Lon    float64
		Lat    float64
	}

	TrackResponse struct {
		Ok      bool
		Message string
	}

	LocateNearbyUserRequest struct {
		UserID uint
	}

	LocateNearbyUserResponse struct {
		Ok      bool
		Message string
		UserIDs []uint
	}

	Track struct {
		userRepo output_port.UserRepo
	}
)

func NewTrackUseCase(userRepo output_port.UserRepo) *Track {
	return &Track{
		userRepo: userRepo,
	}
}

func (instance *Track) Track(ctx context.Context, request TrackRequest) (response TrackResponse, err error) {
	userLocation := entity.UserLocation{
		UserID: request.UserID,
		Lon:    request.Lon,
		Lat:    request.Lat,
	}

	err = instance.userRepo.Track(ctx, userLocation)
	if err != nil {
		return TrackResponse{
			Ok:      false,
			Message: "internal error",
		}, helper.WrapError(err)
	}

	return TrackResponse{
		Ok:      true,
		Message: "success",
	}, nil
}

func (instance *Track) LocateNearbyUser(ctx context.Context, request LocateNearbyUserRequest) (response LocateNearbyUserResponse, err error) {
	userLocation := entity.UserLocation{
		UserID: request.UserID,
	}

	userIDs, err := instance.userRepo.LocateNearbyUser(ctx, userLocation)
	if err != nil {
		return LocateNearbyUserResponse{
			Ok:      false,
			Message: "internal error",
		}, helper.WrapError(err)
	}

	userIDs = slices.DeleteFunc(userIDs, func(i uint) bool { return i == request.UserID })

	return LocateNearbyUserResponse{
		Ok:      true,
		Message: "success",
		UserIDs: userIDs,
	}, nil
}

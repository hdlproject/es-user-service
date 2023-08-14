package interactor

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hdlproject/es-user-service/config"
	"github.com/hdlproject/es-user-service/interface_adapter/database"
)

func TestTrack_LocateNearbyUser(t *testing.T) {
	ctx := context.Background()

	redisClient := database.GetRedisClient(config.Redis{
		Host:     "localhost",
		Port:     "6379",
		Password: "admin",
	})

	userRepo := database.NewUserRepo(nil, redisClient)

	track := NewTrackUseCase(userRepo)

	seederData := []TrackRequest{
		{
			UserID: 1,
			Lon:    -122.27652,
			Lat:    37.805186,
		},
		{
			UserID: 2,
			Lon:    -122.2674626,
			Lat:    37.8062344,
		},
		{
			UserID: 3,
			Lon:    -122.2469854,
			Lat:    37.8104049,
		},
		{
			UserID: 4,
			Lon:    -110,
			Lat:    30,
		},
		{
			UserID: 5,
			Lon:    -122.2612767,
			Lat:    37.7936847,
		},
	}
	for _, item := range seederData {
		_, err := track.Track(ctx, item)
		if err != nil {
			t.Fatal(err)
		}
	}

	res, err := track.LocateNearbyUser(ctx, LocateNearbyUserRequest{
		UserID: 5,
	})
	if err != nil {
		t.Fatal(err)
	}

	expectedRes := []uint{2, 1, 3}
	if diff := cmp.Diff(expectedRes, res.UserIDs); diff != "" {
		t.Fatalf("(-want/+got)\n%s", diff)
	}
}

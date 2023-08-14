package database

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/hdlproject/es-user-service/config"
)

func TestRedisClient_GeoSearchByRadius(t *testing.T) {
	ctx := context.Background()

	redisClient := GetRedisClient(config.Redis{
		Host:     "localhost",
		Port:     "6379",
		Password: "admin",
	})

	key := "location"

	seederData := []struct {
		name string
		lon  float64
		lat  float64
	}{
		{
			name: "loc A",
			lon:  -122.27652,
			lat:  37.805186,
		},
		{
			name: "loc B",
			lon:  -122.2674626,
			lat:  37.8062344,
		},
		{
			name: "loc C",
			lon:  -122.2469854,
			lat:  37.8104049,
		},
		{
			name: "loc D",
			lon:  -110,
			lat:  30,
		},
	}
	for _, item := range seederData {
		err := redisClient.GeoAdd(ctx, key, item.name, item.lon, item.lat)
		if err != nil {
			t.Fatal(err)
		}
	}

	res, err := redisClient.GeoSearchByRadius(ctx, key, -122.2612767, 37.7936847, 5)
	if err != nil {
		t.Fatal(err)
	}

	expectedRes := []string{"loc B", "loc A", "loc C"}
	if diff := cmp.Diff(expectedRes, res); diff != "" {
		t.Fatalf("(-want/+got)\n%s", diff)
	}
}

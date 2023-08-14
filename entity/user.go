package entity

type (
	User struct {
		ID      uint
		Balance uint64
	}

	UserLocation struct {
		UserID uint
		Lon    float64
		Lat    float64
	}
)

package entity

type (
	User struct {
		ID      uint
		Balance uint64
		Auth    UserAuth
	}

	UserAuth struct {
		UserID   uint
		Username string
		Password string
	}

	UserLocation struct {
		UserID uint
		Lon    float64
		Lat    float64
	}
)

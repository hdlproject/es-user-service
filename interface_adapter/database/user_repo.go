package database

import (
	"context"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/hdlproject/es-user-service/entity"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/output_port"
)

type (
	User struct {
		ID          uint      `gorm:"primary_key"`
		Balance     uint64    `gorm:"not null; default:0"`
		DateCreated time.Time `gorm:"autoCreateTime"`
		DateUpdated time.Time `gorm:"autoUpdateTime"`
		DateDeleted gorm.DeletedAt
	}
)

type (
	userRepo struct {
		client      *PostgresClient
		redisClient *RedisClient
	}
)

var userLocationRedisKey = "user:location"

func NewUserRepo(client *PostgresClient, redisClient *RedisClient) output_port.UserRepo {
	return &userRepo{
		client:      client,
		redisClient: redisClient,
	}
}

func (instance *userRepo) Register(userData entity.User) (uint, error) {
	user := User{}.fromEntity(userData)
	err := instance.client.db.Create(&user).Error
	if err != nil {
		return 0, helper.WrapError(err)
	}

	return user.ID, nil
}

func (instance *userRepo) IncreaseBalance(id uint, increment uint64) error {
	user := User{ID: id}
	err := instance.client.db.Transaction(func(tx *gorm.DB) error {
		err := tx.First(&user).Error
		if err != nil {
			return helper.WrapError(err)
		}

		user.Balance = user.Balance + increment
		err = tx.Save(&user).Error
		if err != nil {
			return helper.WrapError(err)
		}

		return nil
	})
	if err != nil {
		return helper.WrapError(err)
	}

	return nil
}

func (instance *userRepo) Track(ctx context.Context, location entity.UserLocation) error {
	err := instance.redisClient.GeoAdd(ctx, userLocationRedisKey, strconv.FormatUint(uint64(location.UserID), 10), location.Lon, location.Lat)
	if err != nil {
		return helper.WrapError(err)
	}

	return nil
}

func (instance *userRepo) LocateNearbyUser(ctx context.Context, location entity.UserLocation) ([]uint, error) {
	userIDs, err := instance.redisClient.GeoSearchByRadius(ctx, userLocationRedisKey, strconv.FormatUint(uint64(location.UserID), 10), 5)
	if err != nil {
		return nil, helper.WrapError(err)
	}

	var userIDsInt []uint
	for _, userID := range userIDs {
		userIDInt, err := strconv.ParseUint(userID, 10, 64)
		if err != nil {
			return nil, helper.WrapError(err)
		}
		userIDsInt = append(userIDsInt, uint(userIDInt))
	}

	return userIDsInt, nil
}

func (User) fromEntity(user entity.User) User {
	return User{
		ID:      user.ID,
		Balance: user.Balance,
	}
}

func (instance User) getEntity() entity.User {
	return entity.User{
		ID:      instance.ID,
		Balance: instance.Balance,
	}
}

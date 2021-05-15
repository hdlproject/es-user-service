package database

import (
	"fmt"
	"github.com/hdlproject/es-user-service/entity"
	"github.com/hdlproject/es-user-service/helper"
	"github.com/hdlproject/es-user-service/use_case/output_port"
	"gorm.io/gorm"
	"time"
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
		client *PostgresClient
	}
)

func NewUserRepo(client *PostgresClient) output_port.UserRepo {
	return &userRepo{
		client: client,
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

		fmt.Println(user.Balance, increment)

		return nil
	})
	if err != nil {
		return helper.WrapError(err)
	}

	return nil
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

package database

import (
	"fmt"
	"github.com/hdlproject/es-user-service/config"
	"github.com/hdlproject/es-user-service/helper"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type (
	PostgresClient struct {
		db *gorm.DB
	}
)

var (
	postgresClient *PostgresClient
)

const (
	dsnTemplate = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta"
)

func GetPostgresClient(dbConfig config.Database) (*PostgresClient, error) {
	if postgresClient == nil {
		client, err := newPostgresClient(dbConfig)
		if err != nil {
			return nil, helper.WrapError(err)
		}

		postgresClient = client
	}

	return postgresClient, nil
}

func newPostgresClient(dbConfig config.Database) (*PostgresClient, error) {
	dsn := fmt.Sprintf(dsnTemplate,
		dbConfig.Host,
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Name,
		dbConfig.Port,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, helper.WrapError(err)
	}

	return &PostgresClient{
		db: db,
	}, nil
}

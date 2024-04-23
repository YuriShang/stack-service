package postgres

import (
	"context"
	"fmt"

	"stackService/dto"
	"stackService/model"
	"stackService/pkg/logging"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Client struct {
	logger logging.Logger
	db     *gorm.DB
}

func NewClient(ctx context.Context, host, port, username, password, database string, logger logging.Logger) (*Client, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	db.AutoMigrate(&model.StackData{})

	logger.Info("postgresql db initiated")
	return &Client{
		logger: logger,
		db:     db}, nil
}

func (c *Client) Create(ctx context.Context, data dto.StackData) (dto.StackData, error) {
	err := c.db.Create(&data).Error
	return data, err
}

func (c *Client) GetAndDelete(ctx context.Context) (dto.StackData, error) {
	var data dto.StackData
	err := c.db.Last(&data).Error
	if err != nil {
		return data, err
	}
	err = c.db.Delete(data).Error
	return data, err
}

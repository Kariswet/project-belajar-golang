package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnection struct {
	Host     string
	Port     string
	Name     string
	Username string
	Password string
}

func NewPostgresConnection(cfg *PostgresConnection) *PostgresConnection {
	return &PostgresConnection{
		Host:     cfg.Host,
		Port:     cfg.Port,
		Name:     cfg.Name,
		Username: cfg.Username,
		Password: cfg.Password,
	}
}

func (cfg *PostgresConnection) InitConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable timezone=Asia/Jakarta",
		cfg.Host, cfg.Port, cfg.Username, cfg.Name, cfg.Password,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return db, nil
}

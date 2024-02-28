package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

func NewDatabase(config DatabaseConfig) (*gorm.DB, error) {
	fmt.Println(fmt.Sprintf("[DB] Connecting to %s", config.Host))
	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", config.Host, config.Username, config.Password, config.Name, config.Port)

	return gorm.Open(postgres.Open(dsn))
}

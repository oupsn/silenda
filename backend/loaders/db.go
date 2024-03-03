package loaders

import (
	"fmt"
	"github.com/wuhoops/silenda/backend/internal/domains"
	"github.com/wuhoops/silenda/backend/internal/utils/db"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabases() {
	CheckAndConnectDatabase()
	CheckAutoMigrate()
}

func CheckAndConnectDatabase() {
	database, err := db.NewDatabase(db.DatabaseConfig{
		Host:     Env.DBHost,
		Port:     Env.DBPort,
		Username: Env.DBUsername,
		Password: Env.DBPassword,
		Name:     Env.DBName,
	})

	DB = database

	if err != nil {
		panic(fmt.Errorf("fatal connecting to database: %w", err))
	}
}

func CheckAutoMigrate() {
	if Env.DBAutoMigrate {
		fmt.Println(fmt.Sprintf("[DB] Automigrate enabled"))
		err := DB.AutoMigrate(&domains.Secret{}, &domains.Workspace{})
		if err != nil {
			panic(err)
		}
		fmt.Println(fmt.Sprintf("[DB] Automigrated."))
	}
}

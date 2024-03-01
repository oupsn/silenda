package domains

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/wuhoops/silenda/backend/internal/utils/v"
	"gorm.io/gorm"
)

type EnvMode string

const (
	Development EnvMode = "dev"
	Staging     EnvMode = "stg"
	Production  EnvMode = "prod"
)

type Secret struct {
	ID          *uuid.UUID `gorm:"primaryKey"`
	WorkspaceID *uuid.UUID
	EnvMode     *EnvMode
	Key         *string `gorm:"unique"`
	Value       *string
	gorm.Model
}

func (secret *Secret) BeforeCreate(tx *gorm.DB) (err error) {
	secret.ID = v.Ptr(uuid.New())
	return
}

func (secret *Secret) BeforeSave(tx *gorm.DB) error {
	validModes := []EnvMode{Development, Staging, Production}
	for _, validMode := range validModes {
		if *secret.EnvMode == validMode {
			return nil
		}
	}
	return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(" %s is not a valid env mode (try 'dev', 'stg' or 'prod')", *secret.EnvMode))
}

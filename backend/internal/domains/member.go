package domains

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Role string

const (
	Owner Role = "owner"
	Admin Role = "admin"
	Dev   Role = "dev"
)

type Member struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Role        Role
	UserID      uuid.UUID
	WorkspaceID uuid.UUID
	User        User
	Workspace   Workspace
	gorm.Model
}

func (member *Member) BeforeCreate(tx *gorm.DB) (err error) {
	member.ID = uuid.New()
	validRoles := []Role{Owner, Admin, Dev}
	for _, validRole := range validRoles {
		if member.Role == validRole {
			return nil
		}
	}
	return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf(" %s is not a valid role (try 'admin' or 'dev')", member.Role))
}

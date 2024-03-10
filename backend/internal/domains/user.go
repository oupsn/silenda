package domains

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	DisplayName string
	Username    string
	password    string
	Workspace   []Workspace
	gorm.Model
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.ID = uuid.New()
	return nil
}

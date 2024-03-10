package domains

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Workspace struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	Name   string
	UserID uuid.UUID
	User   User
	Secret []Secret
	Member []Member
	gorm.Model
}

func (workspace *Workspace) BeforeCreate(tx *gorm.DB) (err error) {
	workspace.ID = uuid.New()
	return nil
}

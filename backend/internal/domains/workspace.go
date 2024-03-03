package domains

import "github.com/google/uuid"

type Workspace struct {
	ID    uuid.UUID `json:"id" gorm:"primaryKey"`
	Name  string    `json:"name" gorm:"primaryKey"`
	Owner uuid.UUID `json:"owner" gorm:"primaryKey"`
}

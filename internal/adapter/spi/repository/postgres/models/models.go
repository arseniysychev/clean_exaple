package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type ModelUUID struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;"`
}

type ModelBase struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

func (b *ModelBase) BeforeCreate(tx *gorm.DB) (err error) {
	b.ID = uuid.New()
	return nil
}

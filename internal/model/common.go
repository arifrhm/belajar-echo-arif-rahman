package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Common struct {
	ID uuid.UUID `json:"id", gorm:"type:uuid;default:uuid_generate_v4()"`

	CreatedAt time.Time  `json:"created_at"`
	CreatedBy string     `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}

func (c *Common) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreatedAt = time.Now()
	c.CreatedBy = "system"
	return
}

func (c *Common) BeforeUpdate(tx *gorm.DB) (err error) {
	t := time.Now()
	c.UpdatedAt = &t
	var s = "system"
	c.UpdatedBy = &s
	return
}

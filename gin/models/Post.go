package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	ID    uuid.UUID `gorm:"index;primary_key;type:uuid;default:uuid_generate_v4()"`
	Title string    `json:"title" binding:"required"`
	Body  string    `json:"body" binding:"required"`
}

func (p *Post) BeforeCreate(tx *gorm.DB) (err error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	return nil
}
func (p *Post) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return nil
}

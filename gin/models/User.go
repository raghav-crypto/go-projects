package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	Name      string         `gorm:"size:20;not null;unique" json:"name" binding:"required"`
	Email     string         `gorm:"size:255;not null;unique" json:"email" binding:"required"`
	Password  string         `gorm:"size:255;not null" json:"password" binding:"required"`
	CreatedAt time.Time      `gorm:"not null" json:"createdAt"`
	UpdatedAt time.Time      `gorm:"not null" json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

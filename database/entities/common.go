package entities

import (
    "time"
    "gorm.io/gorm"
)

type Common struct {
    ID        string         `gorm:"primaryKey;type:varchar(36)" json:"id"` 
    CreatedAt time.Time      `json:"created_at"`
    UpdatedAt time.Time      `json:"updated_at"`
    DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}
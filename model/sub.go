package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/datatypes"
)

type Sub struct {
	Model
	UserUUID uuid.UUID      `json:"user_uuid"`
	SubWords datatypes.JSON `json:"sub_words"`
}

type Model struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

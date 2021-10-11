package model

import (
	"gorm.io/datatypes"
	"time"
)

type Sub struct {
	Model
	UserId   int            `json:"user_id"`
	Phone    string         `json:"phone"`
	SubWords datatypes.JSON `json:"sub_words"`
}

type AllSubWords struct {
	Model
	SubWord string `json:"sub_word" gorm:"unique"`
}

type Model struct {
	ID        uint      `json:"id" gorm:"primary_key"`
	CreatedAt time.Time `json:"created_at" gorm:"created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"updated_at"`
}

func (s Sub) TableName() string {
	return "sub"
}

func (a AllSubWords) TableName() string {
	return "all_sub_words"
}

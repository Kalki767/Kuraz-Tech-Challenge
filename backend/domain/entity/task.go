package entity

import "time"

type Task struct {
	ID          uint   `json:"id" gorm:"primarykey"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

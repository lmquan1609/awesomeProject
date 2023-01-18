package common

import "time"

type SQLModel struct {
	Id        int        `json:"id" gorm:"column:id;"`
	Status    int        `json:"id" gorm:"status"`
	CreatedAt *time.Time `json:"created_at" gorm:"create_at"`
	UpdatedAt *time.Time `json:"updated_at" gorm:"updated_at"`
}

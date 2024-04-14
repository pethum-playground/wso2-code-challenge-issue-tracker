package entity

import "time"

type Issue struct {
	Id          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title,omitempty" gorm:"type:varchar(50);not null"`
	Description string    `json:"description,omitempty" gorm:"type:longtext;nullable"`
	Status      string    `json:"status,omitempty" gorm:"type:varchar(50);not null;"`
	CreatedAt   time.Time `json:"createdAt,omitempty" gorm:"type:datetime;not null;"`
}

func (Issue) TableName() string {
	return "Issue"
}

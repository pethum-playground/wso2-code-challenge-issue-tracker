package entity

import (
	"gorm.io/gorm"
	"time"
)

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

func CountIssues(db *gorm.DB) (int64, int64, int64, int64) {
	var totalCount int64
	var doneCount int64
	var openCount int64
	var inProgressCount int64

	startOfDay := time.Now().Local().Truncate(24 * time.Hour)
	endOfDay := startOfDay.Add(24 * time.Hour)

	db.Model(&Issue{}).Where("createdAt >= ? AND createdAt < ?", startOfDay, endOfDay).Count(&totalCount)
	db.Model(&Issue{}).Where("status = ? AND createdAt >= ? AND createdAt < ?", "DONE", startOfDay, endOfDay).Count(&doneCount)
	db.Model(&Issue{}).Where("status = ? AND createdAt >= ? AND createdAt < ?", "OPEN", startOfDay, endOfDay).Count(&openCount)
	db.Model(&Issue{}).Where("status = ? AND createdAt >= ? AND createdAt < ?", "IN_PROGRESS", startOfDay, endOfDay).Count(&inProgressCount)

	return totalCount, doneCount, openCount, inProgressCount
}

package models

import "time"

type Post struct {
	Id        int64  `gorm:"primaryKey" json:"id"`
	Title     string `gorm:"type:varchar(255)" json:"title"`
	Content   string `gorm:"type:text"`
	Image     string `gorm:"type:varchar(255)" json:"image"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

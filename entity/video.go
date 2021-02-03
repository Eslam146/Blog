package entity

import "time"

type Video struct{
	ID uint64   `json:"id" gorm:"primary_key;auto_increment"`
	Title string  `json:"title" binding:"min=2,max=10" gorm:"type:varchar(100)"`
	Description string `json:"description" binding:"max=10" gorm:"type:varchar(200)"`
	URL string `json:"url" binding:"required,url" gorm:"type:varchar(256);UNIQUE"`
	Author string `json:"author" binding:"required" gorm:"type:varchar(100)"`
	CreatedAt time.Time `json:"_"  gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `json:"_"  gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
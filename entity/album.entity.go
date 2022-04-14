package entity

import "time"

type Album struct {
	ID     		uint64 		`gorm:"primary_key:auto_increment"`
	Title     	string 		`gorm:"type:varchar(255)"`
	Price     	uint64 		`gorm:"type:bigint"`
	Year		int 		`gorm:"type:int"`
	ArtistID    int64 		`gorm:"not null"`
	Artist 		Artist 		`gorm:"foreignkey:ArtistID;constraint:onUpdate:CASCADE,onDelete:CASCADE"`
	CreatedAt	time.Time 	`gorm:"type:timestamp"`
	UpdatedAt 	time.Time 	`gorm:"type:timestamp"`
}
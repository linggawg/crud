package entity

import "time"

type Artist struct {
	ID  		int64    	`gorm:"primary_key:auto_increment"`
	Name      	string   	`gorm:"type:varchar(255)"`
	Email     	string   	`gorm:"type:varchar(255);unique;"`
	Password  	string    	`gorm:"type:varchar(255)"`
	CreatedAt 	time.Time 	`gorm:"type:timestamp"`
	UpdatedAt 	time.Time 	`gorm:"type:timestamp"`
}

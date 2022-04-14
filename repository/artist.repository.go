package repository

import (
	"crud-gin-gorm/entity"
	"crud-gin-gorm/helper"

	"gorm.io/gorm"
)

type ArtistRepository interface {
	Create(artist entity.Artist) (entity.Artist, error)
	Update(user entity.Artist) (entity.Artist, error)
	FindByEmail(email string) (entity.Artist, error)
	FindById(id int) (entity.Artist, error)
}

type artistRepository struct {
	db *gorm.DB
}

func NewArtistRepository(db *gorm.DB) *artistRepository {
	return &artistRepository{db}
}

func (r *artistRepository) Create(artist entity.Artist) (entity.Artist, error) {
	artist.Password = helper.HashAndSalt([]byte(artist.Password))
	err := r.db.Create(&artist).Error
	return artist, err
}

func (r *artistRepository) Update(artist entity.Artist) (entity.Artist, error) {
	if artist.Password != "" {
		artist.Password = helper.HashAndSalt([]byte(artist.Password))
	} else {
		var tempArtist entity.Artist
		r.db.Find(&tempArtist, artist.ID)
		artist.Password = tempArtist.Password
	}

	err := r.db.Save(&artist).Error
	return artist, err
}

func (r *artistRepository) FindByEmail(email string) (entity.Artist, error) {
	var artist entity.Artist
	err := r.db.Where("email = ?", email).Take(&artist).Error
	return artist, err
}

func (r *artistRepository) FindById(id int) (entity.Artist, error) {
	var artist entity.Artist
	err := r.db.Take(&artist, id).Error
	return artist, err
}

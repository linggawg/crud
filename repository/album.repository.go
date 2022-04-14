package repository

import (
	"crud-gin-gorm/entity"

	"gorm.io/gorm"
)

type AlbumRepository interface {
	FindAll() ([]entity.Album, error)
	FindById(id int) (entity.Album, error)
	Create(album entity.Album) (entity.Album, error)
	Update(album entity.Album) (entity.Album, error)
	Delete(album entity.Album) (entity.Album, error)
}

type albumRepository struct{
	db *gorm.DB
}

func NewAlbumRepository (db *gorm.DB) *albumRepository{
	return &albumRepository{db}
} 

func (r *albumRepository) FindAll() ([]entity.Album, error){
	var albums []entity.Album

	err := r.db.Find(&albums).Error
	r.db.Preload("Artist").Find(&albums)
	return albums, err
}

func (r *albumRepository) FindById(id int) (entity.Album, error){
	var album entity.Album

	err := r.db.Take(&album, id).Error
	r.db.Preload("Artist").Find(&album)
	return album, err
}

func (r *albumRepository) Create(album entity.Album) (entity.Album, error){
	err := r.db.Create(&album).Error
	return album, err
}

func (r *albumRepository) Update(album entity.Album) (entity.Album, error){
	err := r.db.Save(&album).Error
	return album, err
}

func (r *albumRepository) Delete(album entity.Album) (entity.Album, error){
	err := r.db.Delete(&album).Error
	return album, err
}
package service

import (
	"crud-gin-gorm/dto"
	"crud-gin-gorm/entity"
	"crud-gin-gorm/repository"

	"github.com/mashingan/smapping"
)

type AlbumService interface {
	FindAll() ([]entity.Album, error)
	FindById(id int) (entity.Album, error)
	Create(request dto.AlbumRequest) (entity.Album, error)
	Update(id int, request dto.AlbumRequest) (entity.Album, error)
	Delete(id int) (entity.Album, error)
}

type albumService struct {
	repository repository.AlbumRepository
}

func NewAlbumService(repository repository.AlbumRepository) *albumService {
	return &albumService{repository}
}

func (s *albumService) FindAll() ([]entity.Album, error) {
	albums, err := s.repository.FindAll()
	return albums, err
}

func (s *albumService) FindById(id int) (entity.Album, error) {
	album, err := s.repository.FindById(id)
	return album, err
}

func (s *albumService) Create(request dto.AlbumRequest) (entity.Album, error) {
	album := entity.Album{}
	err := smapping.FillStruct(&album, smapping.MapFields(&request))
	if err != nil {
		return album, err
	}

	newAlbum, err := s.repository.Create(album)
	return newAlbum, err
}

func (s *albumService) Update(id int, request dto.AlbumRequest) (entity.Album, error) {
	album, err := s.repository.FindById(id)
	if err != nil {
		return album, err
	}

	err = smapping.FillStruct(&album, smapping.MapFields(&request))
	if err != nil {
		return album, err
	}
	
	updatedAlbum, err := s.repository.Update(album)
	return updatedAlbum, err
}

func (s *albumService) Delete(id int) (entity.Album, error) {
	album, err := s.repository.FindById(id)
	if err != nil {
		return album, err
	}
	deletedAlbum, err := s.repository.Delete(album)
	return deletedAlbum, err
}
package service

import (
	"crud-gin-gorm/dto"
	"crud-gin-gorm/entity"
	"crud-gin-gorm/repository"
)

type ArtistService interface {
	Create(request dto.ArtistRequest) (entity.Artist, error)
	Update(id int, request dto.ArtistRequest) (entity.Artist, error)
	FindByEmail(email string) (entity.Artist, error)
	FindById(id int) (entity.Artist, error)
}

type artistService struct {
	repository repository.ArtistRepository
}

func NewArtistService(repository repository.ArtistRepository) *artistService {
	return &artistService{repository}
}

func (s *artistService) FindById(id int) (entity.Artist, error) {
	artist, err := s.repository.FindById(id)
	return artist, err
}

func (s *artistService) FindByEmail(email string) (entity.Artist, error) {
	artist, err := s.repository.FindByEmail(email)
	return artist, err
}

func (s *artistService) Create(request dto.ArtistRequest) (entity.Artist, error) {
	artist := entity.Artist{
		Name: request.Name,
		Email: request.Email,
		Password: request.Password,
	}

	newArtist, err := s.repository.Create(artist)
	return newArtist, err
}

func (s *artistService) Update(id int, request dto.ArtistRequest) (entity.Artist, error) {
	artist, err := s.repository.FindById(id)
	if err != nil {
		return artist, err
	}
	artist.Name = request.Name
	artist.Email = request.Email
	artist.Password = request.Password

	updatedArtist, err := s.repository.Update(artist)
	return updatedArtist, err
}

package dto

import "crud-gin-gorm/entity"

type ArtistResponse struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func ConvertToArtistResponse(artist entity.Artist) ArtistResponse {
	return ArtistResponse{
		ID:    artist.ID,
		Name:  artist.Name,
		Email: artist.Email,
	}
}
package dto

import "crud-gin-gorm/entity"

type AlbumResponse struct {
	ID     uint64         `json:"id"`
	Title  string         `json:"title"`
	Artist ArtistResponse `json:"artist"`
	Price  uint64         `json:"price"`
	Year   int            `json:"year"`
}

func ConvertToAlbumResponse(album entity.Album) AlbumResponse {
	return AlbumResponse{
		ID:     album.ID,
		Title:  album.Title,
		Year:   album.Year,
		Price:  album.Price,
		Artist: ConvertToArtistResponse(album.Artist),
	}
}

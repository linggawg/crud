package dto

type AlbumRequest struct {
	Title  string        `json:"title"`
	Artist ArtistRequest `json:"artist"`
	Price  uint64        `json:"price"`
	Year   int           `json:"year"`
}
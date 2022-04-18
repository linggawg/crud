package dto

type AlbumRequest struct {
	Title string `json:"title"`
	Price uint64 `json:"price"`
	Year  int    `json:"year"`
}
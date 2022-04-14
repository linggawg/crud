package main

import (
	"crud-gin-gorm/config"
	"crud-gin-gorm/handler"
	"crud-gin-gorm/repository"
	"crud-gin-gorm/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db *gorm.DB										= config.SetupDatabaseConnection()
	albumRepository repository.AlbumRepository 		= repository.NewAlbumRepository(db)
	albumService service.AlbumService 				= service.NewAlbumService(albumRepository)
	albumHandler handler.AlbumHandler 				= handler.NewAlbumHandler(albumService)
	artistRepository repository.ArtistRepository 	= repository.NewArtistRepository(db)
	artistService service.ArtistService 			= service.NewArtistService(artistRepository)
	artistHandler handler.ArtistHandler 			= handler.NewArtistHandler(artistService)

)

func main() {
	godotenv.Load()
	defer config.CloseDatabaseConnection(db)
	server := gin.Default()

	albumRoutes := server.Group("api/albums")
	{
		albumRoutes.GET("/", albumHandler.FindAll)
    	albumRoutes.GET("/:id", albumHandler.FindById)
    	albumRoutes.POST("/", albumHandler.Create)
		albumRoutes.PUT("/:id", albumHandler.Update)
		albumRoutes.DELETE("/:id", albumHandler.Delete)
	}
	artistRoutes := server.Group("api/artist")
	{
    	artistRoutes.GET("/:id", artistHandler.FindById)
    	artistRoutes.POST("/", artistHandler.Create)
		artistRoutes.PUT("/:id", artistHandler.Update)
	}

	server.Run(os.Getenv("SERVER_HOST"))
}
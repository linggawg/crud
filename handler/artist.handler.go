package handler

import (
	"crud-gin-gorm/dto"
	"crud-gin-gorm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArtistHandler interface {
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
}

type artistHandler struct{
	service service.ArtistService
}

func NewArtistHandler(service service.ArtistService) *artistHandler{
	return &artistHandler{service}
}

func (h *artistHandler) FindById(ctx *gin.Context) {
    idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	artist, err := h.service.FindById(id)
	if err != nil{
		response := dto.BuildResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest,response)
		panic(err)
	} 
	
	response := dto.BuildResponse(true, "Success", dto.ConvertToArtistResponse(artist))
	ctx.JSON(http.StatusOK,response)
}

func (h *artistHandler) Create(ctx *gin.Context) {
    var request dto.ArtistRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil{
		panic(err)
	}

	artist, err := h.service.Create(request)
	if err != nil{
		response := dto.BuildResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest,response)
		panic(err)
	} 
	
	response := dto.BuildResponse(true, "Success", dto.ConvertToArtistResponse(artist))
	ctx.JSON(http.StatusOK,response)
}

func (h *artistHandler) Update(ctx *gin.Context) {
    var request dto.ArtistRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil{
		panic(err)
	}

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	artist, err := h.service.Update(id, request)
	if err != nil{
		response := dto.BuildResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest,response)
		panic(err)
	} 
	
	response := dto.BuildResponse(true, "Success", dto.ConvertToArtistResponse(artist))
	ctx.JSON(http.StatusOK,response)
}

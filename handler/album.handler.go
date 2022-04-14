package handler

import (
	"crud-gin-gorm/dto"
	"crud-gin-gorm/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AlbumHandler interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type albumHandler struct{
	service service.AlbumService
}

func NewAlbumHandler(service service.AlbumService) *albumHandler{
	return &albumHandler{service}
}

func (h *albumHandler) FindAll(ctx *gin.Context) {
	albums, err := h.service.FindAll()

	var responses []dto.AlbumResponse
	for _, b := range albums{
		responses = append(responses, dto.ConvertToAlbumResponse(b))
	}

	if err != nil{
		response := dto.BuildResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest,response)
		panic(err)
	} 
	
	response := dto.BuildResponse(true, "Success", responses)
	ctx.JSON(http.StatusOK,response)
}

func (h *albumHandler) FindById(ctx *gin.Context) {
    idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	album, err := h.service.FindById(id)
	if err != nil{
		response := dto.BuildResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest,response)
		panic(err)
	} 
	
	response := dto.BuildResponse(true, "Success", dto.ConvertToAlbumResponse(album))
	ctx.JSON(http.StatusOK,response)
}

func (h *albumHandler) Create(ctx *gin.Context) {
    var request dto.AlbumRequest

	err := ctx.ShouldBindJSON(&request)
	if err != nil{
		panic("failed input data")
	}

	album, err := h.service.Create(request)
	if err != nil{
		response := dto.BuildResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest,response)
		panic(err)
	} 	
	
	response := dto.BuildResponse(true, "Success", dto.ConvertToAlbumResponse(album))
	ctx.JSON(http.StatusOK,response)
}

func (h *albumHandler) Update(ctx *gin.Context) {
    var request dto.AlbumRequest
	err := ctx.ShouldBindJSON(&request)
	if err != nil{
		response := dto.BuildResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest,response)
		panic(err)
	}

	idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)
	album, err := h.service.Update(id, request)
	if err != nil{
		response := dto.BuildResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest,response)
		panic(err)
	} 
	
	response := dto.BuildResponse(true, "Success", dto.ConvertToAlbumResponse(album))
	ctx.JSON(http.StatusOK,response)
}

func (h *albumHandler) Delete(ctx *gin.Context) {
    idParam := ctx.Param("id")
	id, _ := strconv.Atoi(idParam)

	album, err := h.service.Delete(id)
	if err != nil{
		response := dto.BuildResponse(false, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest,response)
		panic(err)
	} 
	
	response := dto.BuildResponse(true, "Success", dto.ConvertToAlbumResponse(album))
	ctx.JSON(http.StatusOK,response)
}

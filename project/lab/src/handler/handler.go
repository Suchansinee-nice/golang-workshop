package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-class/lab/model"
	"github.com/golang-class/lab/service"
)

type Handler struct {
	MovieService    service.MovieService
	FavoriteService service.FavoriteService
}

// ListMovie handles the search movie endpoint
func (h *Handler) ListMovie(c *gin.Context) {
	movie, err := h.MovieService.ListMovie(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
}

// GetMovieDetail handles the get movie detail endpoint
func (h *Handler) GetMovieDetail(c *gin.Context) {
	id := c.Param("id")
	detail, err := h.MovieService.GetMovieDetail(c, id)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
		if err.Error() == "movie not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, detail)
}

// GetFavoriteList handles the get favorite movies endpoint
func (h *Handler) GetFavoriteList(c *gin.Context) {
	favorite, err := h.FavoriteService.GetFavorite(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, favorite)
}

func NewHandler(movieService service.MovieService, favoriteService service.FavoriteService) *Handler {
	return &Handler{
		MovieService:    movieService,
		FavoriteService: favoriteService,
	}
}


//2.add handler addfavorites
func (h *Handler) Addfavorite(c *gin.Context) {
	var favorite model.FavoriteMovieAddRequest //declare variable earn from request
	err := c.ShouldBindBodyWithJSON(&favorite)

	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = h.FavoriteService.AddFavorite(c, favorite.MovieID)

	if err != nil {
		if err.Error() == "movie not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err.Error() == "movie error not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "UNKNOWN ERROR"})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"message": "success"})

}

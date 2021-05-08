package handler

import (
	"log"
	"net/http"
	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
	"vote-app-api/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

type masterMovieHandler struct {
	con controllers.MasterMovieController
}

type MasterMovieHandler interface {
	GetMovie(c *gin.Context)
	GetMovies(c *gin.Context)
	CreateMovie(c *gin.Context)
	UpdateMovie(c *gin.Context)
	DeleteMovie(c *gin.Context)
}

func NewMasterMovieHandler(sc controllers.MasterMovieController) MasterMovieHandler {
	return &masterMovieHandler{sc}
}

func (hn *masterMovieHandler) GetMovie(c *gin.Context) {
	req := new(request.GetMasterMovie)
	if err := c.ShouldBindUri(req); err != nil {
		log.Printf("create master movie request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	movie, err := hn.con.GetMovie(req)
	if err != nil {
		log.Printf("master movie controller err is %+v\n", model.ResponseError{Message: err.Error()})
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
	return
}

func (hn *masterMovieHandler) GetMovies(c *gin.Context) {
	movie, err := hn.con.GetMovies()
	if err != nil {
		log.Printf("master movie controller err is %+v\n", model.ResponseError{Message: err.Error()})
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, movie)
	return
}

func (hn *masterMovieHandler) CreateMovie(c *gin.Context) {
	req := new(request.CreateMasterMovie)
	if err := c.ShouldBind(req); err != nil {
		log.Printf("create master movie request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	err := hn.con.CreateMovie(req)
	if err != nil {
		log.Printf("master movie controller err is %+v\n", model.ResponseError{Message: err.Error()})
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "created")
	return
}

func (hn *masterMovieHandler) UpdateMovie(c *gin.Context) {
	req := new(request.UpdateMasterMovie)

	if err := c.ShouldBindUri(req); err != nil {
		log.Printf("update master movie request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	if err := c.ShouldBind(req); err != nil {
		log.Printf("update master movie request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	err := hn.con.UpdateMovie(req)
	if err != nil {
		log.Printf("master movie controller err is %v+\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "updated")
}

func (hn *masterMovieHandler) DeleteMovie(c *gin.Context) {
	req := new(request.DeleteMasterMovie)

	if err := c.ShouldBindUri(req); err != nil {
		log.Printf("delete master movie request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	err := hn.con.DeleteMovie(req)
	if err != nil {
		log.Printf("master movie controller err is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "deleted")
}

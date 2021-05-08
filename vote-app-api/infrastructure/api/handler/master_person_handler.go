package handler

import (
	"log"
	"net/http"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
	"vote-app-api/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

type masterPersonHandler struct {
	con controllers.MasterPersonController
}

type MasterPersonHandler interface {
	GetPerson(c *gin.Context)
	GetPersons(c *gin.Context)
	CreatePerson(c *gin.Context)
	UpdatePerson(c *gin.Context)
	DeletePerson(c *gin.Context)
}

func NewMasterPersonHandler(sc controllers.MasterPersonController) MasterPersonHandler {
	return &masterPersonHandler{sc}
}

func (hn *masterPersonHandler) GetPerson(c *gin.Context) {
	req := new(request.GetMasterPerson)
	if err := c.ShouldBindUri(req); err != nil {
		log.Printf("get master person request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	person, err := hn.con.GetPerson(req)
	if err != nil {
		log.Printf("master person controller err is %+v\n", model.ResponseError{Message: err.Error()})
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
	return
}

func (hn *masterPersonHandler) GetPersons(c *gin.Context) {
	person, err := hn.con.GetPersons()
	if err != nil {
		log.Printf("master person controller err is %+v\n", model.ResponseError{Message: err.Error()})
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, person)
	return
}

func (hn *masterPersonHandler) CreatePerson(c *gin.Context) {
	req := new(request.CreateMasterPerson)
	if err := c.ShouldBind(req); err != nil {
		log.Printf("create master person request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	err := hn.con.CreateMasterPerson(req)
	if err != nil {
		log.Printf("master person controller err is %+v\n", model.ResponseError{Message: err.Error()})
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, "created")
	return
}

func (hn *masterPersonHandler) UpdatePerson(c *gin.Context) {
	req := new(request.UpdateMasterPerson)

	if err := c.ShouldBindUri(req); err != nil {
		log.Printf("update master person request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	if err := c.ShouldBind(req); err != nil {
		log.Printf("update master person request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	err := hn.con.UpdateMasterPerson(req)
	if err != nil {
		log.Printf("master person controller err is %v+\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "updated")
}

func (hn *masterPersonHandler) DeletePerson(c *gin.Context) {
	req := new(request.DeleteMasterPerson)

	if err := c.ShouldBindUri(req); err != nil {
		log.Printf("delete master person request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	err := hn.con.DeleteMasterPerson(req)
	if err != nil {
		log.Printf("master person controller err is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, "deleted")
}

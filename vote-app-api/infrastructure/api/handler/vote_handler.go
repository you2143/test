package handler

import (
	"log"
	"net/http"

	"vote-app-api/domain/model"
	"vote-app-api/domain/request"
	"vote-app-api/interfaces/controllers"

	"github.com/gin-gonic/gin"
)

type voteHandler struct {
	con controllers.VoteController
}

type VoteHandler interface {
	CreateVote(c *gin.Context)
	GetVoteResult(c *gin.Context)
	GetVoteIsEnd(c *gin.Context)
}

func NewVoteHandler(sc controllers.VoteController) VoteHandler {
	return &voteHandler{sc}
}

func (hn *voteHandler) CreateVote(c *gin.Context) {
	req := new(request.CreateVote)
	if err := c.ShouldBind(req); err != nil {
		log.Printf("create vote request error is %+v\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	err := hn.con.CreateVote(req)
	if err != nil {
		log.Printf("vote controller err is %+v\n", model.ResponseError{Message: err.Error()})
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, "created")
	return
}

func (hn *voteHandler) GetVoteResult(c *gin.Context) {
	vote, err := hn.con.GetALLVote()
	if err != nil {
		log.Printf("vote controller err is %v+\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, vote)
}

func (hn *voteHandler) GetVoteIsEnd(c *gin.Context) {
	isEnd, err := hn.con.GetVoteIsEnd()
	if err != nil {
		log.Printf("vote controller err is %v+\n", err)
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, isEnd)
}

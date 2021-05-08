package router

import (
	"vote-app-api/docs"
	"vote-app-api/registry"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func NewApiV1Router(g *gin.Engine, conn *gorm.DB) {
	r := registry.NewInteractor(conn)
	v1 := g.Group("api/v1")
	docs.SwaggerInfo.BasePath = "/"

	// For Men
	masterMovieHandler := r.NewMasterMovieHandler()
	masterPersonHandler := r.NewMasterPersonHandler()
	voteHandler := r.NewVoteHandler()

	{
		masterMovies := v1.Group("/masterMovie")
		{
			masterMovies.POST("", masterMovieHandler.CreateMovie)
			masterMovies.GET("", masterMovieHandler.GetMovies)
			masterMovies.GET("/:id", masterMovieHandler.GetMovie)
			masterMovies.PUT("/:id", masterMovieHandler.UpdateMovie)
			masterMovies.DELETE("/:id", masterMovieHandler.DeleteMovie)
		}

		masterPerson := v1.Group("/masterPerson")
		{
			masterPerson.POST("", masterPersonHandler.CreatePerson)
			masterPerson.GET("", masterPersonHandler.GetPersons)
			masterPerson.GET("/:id", masterPersonHandler.GetPerson)
			masterPerson.PUT("/:id", masterPersonHandler.UpdatePerson)
			masterPerson.DELETE("/:id", masterPersonHandler.DeletePerson)
		}

		vote := v1.Group("/vote")
		{
			vote.POST("", voteHandler.CreateVote)
			vote.GET("", voteHandler.GetVoteResult)
		}

		voteIsEnd := v1.Group("/voteIsEnd")
		{
			voteIsEnd.GET("", voteHandler.GetVoteIsEnd)
		}

	}

	return
}

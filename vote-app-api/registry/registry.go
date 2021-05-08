package registry

import (
	"vote-app-api/infrastructure/api/handler"
	"vote-app-api/infrastructure/datastore"
	"vote-app-api/interfaces/controllers"
	"vote-app-api/interfaces/presenters"
	"vote-app-api/usecase/presenter"
	"vote-app-api/usecase/repository"
	"vote-app-api/usecase/service"

	"github.com/jinzhu/gorm"
)

type interactor struct {
	conn *gorm.DB
}

type Interactor interface {
	NewMasterMovieHandler() handler.MasterMovieHandler
	NewMasterPersonHandler() handler.MasterPersonHandler
	NewVoteHandler() handler.VoteHandler
}

func NewInteractor(conn *gorm.DB) Interactor {
	return &interactor{
		conn: conn,
	}
}

/* MasterMovie */
func (i *interactor) NewMasterMovieHandler() handler.MasterMovieHandler {
	return handler.NewMasterMovieHandler(i.NewMasterMovieController())
}

func (i *interactor) NewMasterMovieController() controllers.MasterMovieController {
	return controllers.NewMasterMovieController(i.NewMasterMovieService())
}

func (i *interactor) NewMasterMovieService() service.MasterMovieService {
	return service.NewMasterMovieService(i.NewMasterMovieRepository(), i.NewMasterMoviePresenter())
}

func (i *interactor) NewMasterMovieRepository() repository.MasterMovieRepository {
	return datastore.NewMasterMovieRepository(i.conn)
}

func (i *interactor) NewMasterMoviePresenter() presenter.MasterMoviePresenter {
	return presenters.NewMasterMoviePresenter()
}

/* End MasterMovie */

/* MasterPerson */
func (i *interactor) NewMasterPersonHandler() handler.MasterPersonHandler {
	return handler.NewMasterPersonHandler(i.NewMasterPersonController())
}

func (i *interactor) NewMasterPersonController() controllers.MasterPersonController {
	return controllers.NewMasterPersonController(i.NewMasterPersonService())
}

func (i *interactor) NewMasterPersonService() service.MasterPersonService {
	return service.NewMasterPersonService(i.NewMasterPersonRepository(), i.NewMasterPersonPresenter())
}

func (i *interactor) NewMasterPersonRepository() repository.MasterPersonRepository {
	return datastore.NewMasterPersonRepository(i.conn)
}

func (i *interactor) NewMasterPersonPresenter() presenter.MasterPersonPresenter {
	return presenters.NewMasterPersonPresenter()
}

/* End Vote */

/* Vote */
func (i *interactor) NewVoteHandler() handler.VoteHandler {
	return handler.NewVoteHandler(i.NewVoteController())
}

func (i *interactor) NewVoteController() controllers.VoteController {
	return controllers.NewVoteController(i.NewVoteService())
}

func (i *interactor) NewVoteService() service.VoteService {
	return service.NewVoteService(i.NewVoteRepository(), i.NewVotePresenter())
}

func (i *interactor) NewVoteRepository() repository.VoteRepository {
	return datastore.NewVoteRepository(i.conn)
}

func (i *interactor) NewVotePresenter() presenter.VotePresenter {
	return presenters.NewVotePresenter()
}

/* End Vote */

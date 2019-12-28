package registry

import (
	"github.com/easedot/goarc/adapters/controller"
	"github.com/easedot/goarc/adapters/repository"
	"github.com/easedot/goarc/adapters/presenter"
	"github.com/easedot/goarc/usecases/interactor"
)

func(r *registry) NewArticleController() controller.ArticleController {
	ar:= repository.NewArticleRepository(r.db)
	ap:= presenter.NewArticlePresenter()
	ai:= interactor.NewArticleInteractor(ar,ap)
	return controller.NewArticleController(ai)
}
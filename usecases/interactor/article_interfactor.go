package interactor

import (
	"github.com/easedot/goarc/entities"
	"github.com/easedot/goarc/usecases/presenter"
	"github.com/easedot/goarc/usecases/repository"
)

//for use case business logic
type ArticleInteractor interface {
	Query(u []*entities.Article)([]*entities.Article,error)
}

//inject from outer object to interface
func NewArticleInteractor(r repository.ArticleRepository,p presenter.ArticlePresenter) ArticleInteractor{
	return &articleInteractor{p,r}
}

type articleInteractor struct {
	ArticlePresenter presenter.ArticlePresenter
	ArticleRepository repository.ArticleRepository
}

func (as *articleInteractor) Query(u []*entities.Article)([]*entities.Article,error)  {
	a,err:=as.ArticleRepository.Query(u)
	if err!=nil{
		return nil,err
	}
	//user case modify response,use inject adapter response object
	a=as.ArticlePresenter.ResponseArticles(a)
	return a, nil
}
package interactor

import (
	"github.com/easedot/goarc/domain"
	"github.com/easedot/goarc/usecases/presenter"
	"github.com/easedot/goarc/usecases/repository"
)

//for use case business logic
type ArticleInteractor interface {
	Query(u []*domain.Article)([]*domain.Article,error)
	Find(u *domain.Article)(*domain.Article,error)
	Update(u *domain.Article) error
}

//inject from outer object to interface
func NewArticleInteractor(r repository.ArticleRepository,p presenter.ArticlePresenter) ArticleInteractor{
	return &articleInteractor{p,r}
}

type articleInteractor struct {
	ArticlePresenter presenter.ArticlePresenter
	ArticleRepository repository.ArticleRepository
}

func (ai *articleInteractor) Query(u []*domain.Article)([]*domain.Article,error)  {
	a,err:= ai.ArticleRepository.Query(u)
	if err!=nil{
		return nil,err
	}
	//user case modify response,use inject adapter response object
	a= ai.ArticlePresenter.ResponseArticles(a)
	return a, nil
}

func (ai *articleInteractor) Find(u *domain.Article)(*domain.Article,error)  {
	a,err:= ai.ArticleRepository.Find(u)
	if err!=nil{
		return nil,err
	}
	//user case modify response,use inject adapter response object
	a= ai.ArticlePresenter.ResponseArticle(a)
	return a, nil
}
func (ai *articleInteractor) Update(u *domain.Article)(error)  {
	err:= ai.ArticleRepository.Update(u)
	if err!=nil{
		return err
	}
	//user case modify response,use inject adapter response object
	u= ai.ArticlePresenter.ResponseArticle(u)
	return nil
}
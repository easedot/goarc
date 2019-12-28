package controller

import (
	"net/http"

	"github.com/easedot/goarc/entities"
	"github.com/easedot/goarc/usecases/interactor"
)

type ArticleController interface {
	GetArticles(c Context) error
}

type articleController struct {
	//call usecase to implete logic
	articleInteractor interactor.ArticleInteractor
}

func NewArticleController(ai interactor.ArticleInteractor) ArticleController{
	return &articleController{ai}
}

func (ai *articleController)GetArticles(c Context) error  {
	//controller use usecase to query data
	var ars []*entities.Article
	ar:=&entities.Article{ID:1}
	ars=append(ars,ar)
	a,err:=ai.articleInteractor.Query(ars)
	if err!=nil{
		return err
	}
	return c.JSON(http.StatusOK,a)
}


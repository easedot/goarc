package controller

import (
	"net/http"

	"github.com/easedot/goarc/domain"
	"github.com/easedot/goarc/usecases/interactor"
)

type AuthorController interface {
	GetAutuors(c Context) error
}

type authorController struct {
	//call usecase to implete logic
	articleInteractor interactor.ArticleInteractor
}

func NewAuthorController(ai interactor.ArticleInteractor) AuthorController{
	return &authorController{ai}
}

func (ai *authorController)GetAutuors(c Context) error  {
	//controller use usecase to query data
	var ars []*domain.Article
	ar:=&domain.Article{ID: 1}
	ars=append(ars,ar)
	a,err:=ai.articleInteractor.Query(ars)
	if err!=nil{
		return err
	}
	return c.JSON(http.StatusOK,a)
}


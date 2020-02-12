package controller

import (
	"net/http"
	"strconv"

	"github.com/easedot/goarc/domain"
	"github.com/easedot/goarc/usecases/interactor"
)

type ArticleController interface {
	GetArticles(c Context) error
	GetArticle(c Context) error
	UpdateArticle(c Context) error
}

type articleController struct {
	//call usecase to implete logic
	articleInteractor interactor.ArticleInteractor
}

func NewArticleController(ai interactor.ArticleInteractor) ArticleController{
	return &articleController{ai}
}

// GetArticles godoc
// @Summary get article list
// @Description
// @Tags Article
// @Accept json
// @Produce json
// @Router /articles [get]
// @Success 200 {array} domain.Article
func (ac *articleController)GetArticles(c Context) error  {
	//controller use usecase to query data
	var ars []*domain.Article
	ar:=&domain.Article{ID: 1}
	ars=append(ars,ar)
	a,err:= ac.articleInteractor.Query(ars)
	if err!=nil{
		return err
	}
	return c.JSON(http.StatusOK,a)
}

// GetArticles godoc
// @Summary get article list
// @Description
// @Tags Article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Router /article/{id} [get]
// @Success 200 {object} domain.Article
func (ac *articleController)GetArticle(c Context) error  {
	id,_:=strconv.Atoi(c.Param("id"))
	ar:=&domain.Article{ID: int64(id)}
	a,err:= ac.articleInteractor.Find(ar)
	if err!=nil{
		return err
	}
	return c.JSON(http.StatusOK,a)
}

// Update Article godoc
// @Summary update article
// @Description
// @Tags Article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Param  article body domain.Article true "Update article"
// @Router /article/{id} [put]
// @Success 200 {object} domain.Article
func (ac *articleController)UpdateArticle(c Context) error  {
	id,_:=strconv.Atoi(c.Param("id"))
	ar:=&domain.Article{}
	c.Bind(ar)
	ar.ID = int64(id)
	err:= ac.articleInteractor.Update(ar)
	if err!=nil{
		return err
	}
	return c.JSON(http.StatusOK,ar)
}

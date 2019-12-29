package controller

import (
	"net/http"
	"strconv"

	"github.com/easedot/goarc/entities"
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
// @Success 200 {array} entities.Article
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

// GetArticles godoc
// @Summary get article list
// @Description
// @Tags Article
// @Accept json
// @Produce json
// @Param id path int true "Article ID"
// @Router /article/{id} [get]
// @Success 200 {object} entities.Article
func (ai *articleController)GetArticle(c Context) error  {
	id,_:=strconv.Atoi(c.Param("id"))
	ar:=&entities.Article{ID:int64(id)}
	a,err:=ai.articleInteractor.Find(ar)
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
// @Param  article body entities.Article true "Update article"
// @Router /article/{id} [put]
// @Success 200 {object} entities.Article
func (ai *articleController)UpdateArticle(c Context) error  {
	id,_:=strconv.Atoi(c.Param("id"))
	ar:=&entities.Article{}
	c.Bind(ar)
	ar.ID = int64(id)
	err:=ai.articleInteractor.Update(ar)
	if err!=nil{
		return err
	}
	return c.JSON(http.StatusOK,ar)
}

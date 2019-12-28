package presenter

import (
	"github.com/easedot/goarc/entities"
	"github.com/easedot/goarc/usecases/presenter"
)


func NewArticlePresenter() presenter.ArticlePresenter{
	return &articlePresenter{}
}

type articlePresenter struct {
}

func (ap *articlePresenter) ResponseArticles(as [] *entities.Article) []*entities.Article  {
	for _, a := range as{
		a.Title="EASE"+a.Title
	}
	return as
}

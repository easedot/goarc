package presenter

import (
	"github.com/easedot/goarc/domain"
	"github.com/easedot/goarc/usecases/presenter"
)


func NewArticlePresenter() presenter.ArticlePresenter{
	return &articlePresenter{}
}

type articlePresenter struct {
}

func (ap *articlePresenter) ResponseArticles(as [] *domain.Article) []*domain.Article  {
	for _, a := range as{
		a.Title="EASE_"+a.Title
	}
	return as
}
func (ap *articlePresenter) ResponseArticle(as *domain.Article) *domain.Article  {
	as.Title="EASE_"+as.Title
	return as
}

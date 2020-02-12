package presenter

import (
	"github.com/easedot/goarc/domain"
)

type ArticlePresenter interface {
	ResponseArticles(as [] *domain.Article) []*domain.Article
	ResponseArticle(as *domain.Article) *domain.Article
}
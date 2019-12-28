package presenter

import (
	"github.com/easedot/goarc/entities"
)

type ArticlePresenter interface {
	ResponseArticles(as [] *entities.Article) []*entities.Article
}
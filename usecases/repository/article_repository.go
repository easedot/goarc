package repository

import (
	"github.com/easedot/goarc/domain"
)

type ArticleRepository interface {
	Query(u [] *domain.Article ) ([]*domain.Article,error)
	Find(u *domain.Article)(*domain.Article,error)
	Update(u *domain.Article) error
}
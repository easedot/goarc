package repository

import (
	"github.com/easedot/goarc/entities"
)

type ArticleRepository interface {
	Query(u [] *entities.Article ) ([]*entities.Article,error)
}
package repository

import (
	"github.com/easedot/godbs"

	"github.com/easedot/goarc/entities"
)

type articleRepository struct {
	db *godbs.DbHelper
}

type ArticleRepository interface {
	Query(u [] *entities.Article ) ([]*entities.Article,error)
}

func NewArticleRepository(db *godbs.DbHelper) ArticleRepository{
	return &articleRepository{db:db}
}

func (ar * articleRepository) Query(u [] *entities.Article ) ( []*entities.Article,error)  {
	var rst []*entities.Article
	if err:=ar.db.Query(u[0],&rst);err!=nil{
		return nil,err
	}
	return rst,nil
}
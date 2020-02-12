package repository

import (
	"github.com/easedot/godbs"

	"github.com/easedot/goarc/domain"
	"github.com/easedot/goarc/usecases/repository"
)

type articleRepository struct {
	db *godbs.DbHelper
}

func NewArticleRepository(db *godbs.DbHelper) repository.ArticleRepository{
	return &articleRepository{db:db}
}

func (ar * articleRepository) Query(u [] *domain.Article ) ( []*domain.Article,error)  {
	var rst []*domain.Article
	if err:=ar.db.Query(u[0],&rst);err!=nil{
		return nil,err
	}
	return rst,nil
}
func (ar * articleRepository) Find(u *domain.Article ) ( *domain.Article,error)  {
	if err:=ar.db.Find(u);err!=nil{
		return nil,err
	}
	return u,nil
}
func (ar * articleRepository) Update(u *domain.Article ) error {
	if err:=ar.db.Update(u);err!=nil{
		return err
	}
	return nil
}
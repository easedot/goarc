package repository

import (
	"github.com/easedot/godbs"

	"github.com/easedot/goarc/entities"
	"github.com/easedot/goarc/usecases/repository"
)

type articleRepository struct {
	db *godbs.DbHelper
}

func NewArticleRepository(db *godbs.DbHelper) repository.ArticleRepository{
	return &articleRepository{db:db}
}

func (ar * articleRepository) Query(u [] *entities.Article ) ( []*entities.Article,error)  {
	var rst []*entities.Article
	if err:=ar.db.Query(u[0],&rst);err!=nil{
		return nil,err
	}
	return rst,nil
}
func (ar * articleRepository) Find(u *entities.Article ) ( *entities.Article,error)  {
	if err:=ar.db.Find(u);err!=nil{
		return nil,err
	}
	return u,nil
}
func (ar * articleRepository) Update(u *entities.Article ) error {
	if err:=ar.db.Update(u);err!=nil{
		return err
	}
	return nil
}
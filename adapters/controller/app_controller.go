package controller

type AppController interface {
	ArticleController
	AuthorController
}
type Controller struct {
	ArticleController
	AuthorController
}

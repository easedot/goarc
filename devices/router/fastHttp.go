package router

import (
	"encoding/json"

	"github.com/easedot/goarc/adapters/controller"
	"github.com/fasthttp/websocket"
	"github.com/gomodule/redigo/redis"
	"github.com/savsgio/atreugo/v10"
)

type fastWrap struct {
	c *atreugo.RequestCtx
}

func (f *fastWrap) Param(name string) string {
	return string(f.c.QueryArgs().Peek(name))
}
func (f *fastWrap) Bind(i interface{}) error {
	return json.Unmarshal(f.c.PostBody(), i)
}
func (f *fastWrap) JSON(code int, i interface{}) error {
	return f.c.JSONResponse(i, code)
}

func NewFRouter(e *atreugo.Atreugo, c controller.AppController, rds *redis.Pool) *atreugo.Atreugo {
	e.Static("/assets/", "public")
	v1 := e.NewGroupPath("/api/v1")
	v1.GET("/articles", func(context *atreugo.RequestCtx) error {
		return c.GetArticles(&fastWrap{c: context})
	})
	v1.GET("/article/:id", func(context *atreugo.RequestCtx) error { return c.GetArticle(&fastWrap{c: context}) })
	v1.PUT("/article/:id", func(context *atreugo.RequestCtx) error { return c.UpdateArticle(&fastWrap{c: context}) })
	v1.GET("/authors", func(context *atreugo.RequestCtx) error { return c.GetAutuors(&fastWrap{c: context}) })

	chat := NewChatServer(rds)
	e.GET("/ws", func(c *atreugo.RequestCtx) error {
		var upgrader = websocket.FastHTTPUpgrader{}
		//var upgrader = websocket.Upgrader{}
		// Upgrade initial GET request to a websocket
		//ws, err := upgrader.Upgrade( c.Response(), c.Request(), nil)
		upgrader.Upgrade(c.RequestCtx, func(ws *websocket.Conn) {
			chat.ChatHandle(ws)
		})
		return nil
	})

	return e
}

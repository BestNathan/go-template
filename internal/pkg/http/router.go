package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IRouter interface {
	Route(r gin.IRouter)
}

type Router struct {
	r *gin.Engine
}

func NewRouter() *Router {
	r := gin.New()
	r.Use(gin.Recovery()).Use(gin.Logger())

	return &Router{r}
}

func NewWithGinEngine(r *gin.Engine) *Router {
	return &Router{r}
}

func (r *Router) Route(rts ...IRouter) {
	for _, rt := range rts {
		rt.Route(r.r)
	}
}

func (r *Router) Handler() http.Handler {
	return r.r
}

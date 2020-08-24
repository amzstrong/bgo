package app

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Bgo struct {
	//routerMap map[string]HandlerFunc
	Router *Router
}

func NewBgo() *Bgo {
	return &Bgo{Router: newRouter()}
}

func (bgo *Bgo) GET(pattern string, handler HandlerFunc) {
	bgo.Router.AddRoute("GET", pattern, handler)
}

func (bgo *Bgo) POST(pattern string, handler HandlerFunc) {
	bgo.Router.AddRoute("POST", pattern, handler)
}

func (bgo *Bgo) Run(addr string) error {
	err := http.ListenAndServe(addr, bgo)
	return err
}

func (bgo *Bgo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c:=newContext(w,r)
	bgo.Router.Handle(c)
}

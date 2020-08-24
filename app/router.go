package app

import (
	"fmt"
	"log"
)


type Router struct {
	RouteMap map[string]HandlerFunc
}

func newRouter() *Router {
	return &Router{RouteMap: make(map[string]HandlerFunc)}
}

func (r *Router) AddRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %2s - %s", method, pattern)
	key := method + "-" + pattern
	r.RouteMap[key] = handler
}

func (r *Router) Handle(c *Context) {
	var key string
	key = c.Method + "-" + c.Path
	fmt.Println(key)
	if handle, ok := r.RouteMap[key]; ok {
		handle(c)
	} else {
		fmt.Fprintf(c.Writer, "404 NOT FOUND  url:%s", c.Req.URL)
	}
}

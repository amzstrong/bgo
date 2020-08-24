package app

import "net/http"

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Router struct {
	routerMap map[string]HandlerFunc
}

func New() *Router {
	return &Router{routerMap: make(map[string]HandlerFunc)}
}

//添加路由  GET-/index  =>   handleIndex()
func (router *Router) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	router.routerMap[key] = handler
}

func (router *Router) GET(pattern string, handler HandlerFunc) {
	router.addRouter("GET", pattern, handler)
}

func (router *Router) POST(pattern string, handler HandlerFunc) {
	router.addRouter("POST", pattern, handler)
}



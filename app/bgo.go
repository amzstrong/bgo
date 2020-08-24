package app

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Bgo struct {
	routerMap map[string]HandlerFunc
}

func NewBgo() *Bgo {
	return &Bgo{routerMap: make(map[string]HandlerFunc)}
}

//添加路由  GET-/index  =>   handleIndex()
func (bgo *Bgo) addRouter(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	bgo.routerMap[key] = handler
}

func (bgo *Bgo) GET(pattern string, handler HandlerFunc) {
	bgo.addRouter("GET", pattern, handler)
}

func (bgo *Bgo) POST(pattern string, handler HandlerFunc) {
	bgo.addRouter("POST", pattern, handler)
}

func (bgo *Bgo) Run(addr string) error {
	err := http.ListenAndServe(addr, bgo)
	return err
}

func (bgo *Bgo) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	var key string

	key = r.Method+"-"+r.URL.Path

	fmt.Println(key)
	if handle,ok :=bgo.routerMap[key];ok{
		handle(w,r)
	}else{
		fmt.Fprintf(w,"404 NOT FOUND  url:%s",r.URL)
	}
}

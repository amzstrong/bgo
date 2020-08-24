package main

import (
	"bgo/app"
	"net/http"
)

//
//func main() {
//	http.HandleFunc("/", indexHandler)
//	http.HandleFunc("/hello", helloHandler)
//	log.Fatal(http.ListenAndServe("localhost:8000", nil))
//}
//
//// handler echoes r.URL.Path
//func indexHandler(w http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//}
//
//// handler echoes r.URL.Header
//func helloHandler(w http.ResponseWriter, req *http.Request) {
//	for k, v := range req.Header {
//		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
//	}
//	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//	fmt.Fprintf(w, "URL.Path = %q\n", req.Method)
//
//}

func main() {
	bgo := app.NewBgo()
	bgo.GET("/", func(c *app.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})


	bgo.GET("/hello", func(c *app.Context) {
		c.JSON(http.StatusOK, app.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	bgo.Run(":9999")


}


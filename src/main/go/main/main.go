package main

import (
	"fmt"
	logger "main/go/common/utils/log"
	"main/go/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	logger.Info(http.ListenAndServe(":8080", router))
	logger.Flush()
}
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

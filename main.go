package main

import (
	"fmt"
	"github.com/Mainzxq/go_article/utils"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"net/http"
)


func AuthMiddleWareHandler(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// do some middleware things before.
	fmt.Println("good, middleware is up")
	fmt.Println("yes")
	next(w, r)
}


func RegisterHandlers() *mux.Router {
	basicRouter := mux.NewRouter().StrictSlash(true)
	basicRouter.HandleFunc("/", DbConnectTest).Methods("GET")
	basicRouter.HandleFunc("/user/create", CreateUser).Methods("POST")



	return basicRouter
}

func main() {
	r := RegisterHandlers()
	n := negroni.New()
	n.Use(negroni.HandlerFunc(AuthMiddleWareHandler))
	n.UseHandler(r)
	fmt.Println(utils.MakeUUIDS())
	n.Run(":8000")
}

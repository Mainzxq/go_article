package main

import (
	"fmt"
	"github.com/Mainzxq/go_article/utils"
	"github.com/gorilla/mux"
	"net/http"
)

type middleWrerHandler struct {
	r *mux.Router
}

func RegisterHandlers() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", DbConnectTest).Methods("GET")
	r.HandleFunc("/user/{user_name}", CreateUser).Methods("POST")


	//router := httprouter.New()
	//router.POST("/user/:user_name/:pwd", CreateUser)
	//router.GET("/", DbConnectTest)
	//router.POST("/user/:user_name", Login)
	return r
}

func main() {
	r := RegisterHandlers()
	fmt.Println(utils.MakeUUIDS())
	http.ListenAndServe(":8000", r)
}

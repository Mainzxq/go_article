package main

import (
	"fmt"
	"github.com/Mainzxq/go_article/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)



func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.GET("/", DbConnectTest)

	router.POST("/user/:user_name", Login)
	return router
}

func main() {
	r := RegisterHandlers()
	fmt.Println(utils.MakeUUIDS())
	http.ListenAndServe(":8000", r)
}

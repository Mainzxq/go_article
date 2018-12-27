package main

import (
	"fmt"
	"github.com/Mainzxq/go_article/dbops"
	"github.com/Mainzxq/go_article/defs"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	un := p.ByName("user_name")
	pwd := p.ByName("pwd")
	var newUser = defs.UserCredential{
		Username:un,
		Pwd: pwd,
	}
	res, err := dbops.CreateUser(newUser)
	if err != nil {
		fmt.Println("this is a handler error ")
		log.Fatal(err)
	}
	var str = "create user success:"+strconv.FormatBool(res)
	io.WriteString(w, str)
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uname := p.ByName("user_name")
	io.WriteString(w, "user was Created!: "+uname)
}

func DbConnectTest(w http.ResponseWriter, r *http.Request) {
	feedback_words := dbops.TestDbConnet()
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, feedback_words)
}
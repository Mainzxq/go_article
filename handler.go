package main

import (
	"github.com/Mainzxq/go_article/dbops"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	io.WriteString(w, "create user handler")
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params)  {
	uname := p.ByName("user_name")
	io.WriteString(w, "user was Created!: "+uname)
}

func DbConnectTest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	feedback_words := dbops.TestDbConnet()
	io.WriteString(w, feedback_words)
}
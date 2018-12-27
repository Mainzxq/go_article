package main

import (
	"encoding/json"
	"fmt"
	"github.com/Mainzxq/go_article/dbops"
	"github.com/Mainzxq/go_article/defs"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func AuthForMe(w http.ResponseWriter, r *http.Request) {

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	vals := mux.Vars(r)
	un, _ := vals["user_name"]
	var req map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("body have some problems")
		log.Fatal(err)
		return
	}

	var newUser = defs.UserCredential{
		Username:un,
		Pwd: req["pwd"].(string),
		Email: req["email"].(string),
		Phone: req["phone"].(string),
		Token: req["token"].(string),
	}

	res, err := dbops.CreateUser(newUser)
	if err != nil {
		fmt.Println("this is a handler error ", res)
		log.Fatal(err)
	}

	var result = map[string]string{"result":"success", "name":un,"type":"CreateUser"}
	response, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(response)
	if err != nil {
		log.Fatal(err)
		return
	}
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
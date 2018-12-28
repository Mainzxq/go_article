package main

import (
	"encoding/json"
	"fmt"
	"github.com/Mainzxq/go_article/dbops"
	"github.com/Mainzxq/go_article/defs"
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"

	"io"
	"io/ioutil"
	"log"
	"net/http"
)


var Secret = []byte("lmshiwomendelianmeng")

func GetOneToken() (string, error) {
	claim := defs.SignClaim {
		"mainzxq",
		"192.168.43.171",
		1,
		2,
		2,
		jwt.StandardClaims{
			//ExpiresAt: 15000,
			Issuer: "mainzxq",
	},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenStr, err := token.SignedString(Secret)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	fmt.Println(tokenStr)
	return tokenStr, nil
}

func ParseOneToken(token string) (string, error) {
	token1, err := jwt.Parse(token, func(token1 *jwt.Token)(interface{}, error){
		return []byte(Secret),nil
	})
	if err != nil {
		log.Fatal(err)
		return "", nil
	}
	if token1.Valid {
		fmt.Println(token1)
		buf, err := json.Marshal(token1)
		if err != nil {
			log.Fatal(err)
			return "", nil
		}
		fmt.Println("up here is a Claim structure")
		return string(buf), nil
	}else {
		fmt.Println("token not valid!")
		return "", nil
	}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var req map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &req)
	if err != nil {
		fmt.Println("body have some problems")
		log.Fatal(err)
		return
	}

	var newUser = defs.UserCredential{
		Username: req["user_name"].(string),
		Pwd: req["pwd"].(string),
		Email: req["email"].(string),
		Phone: req["phone"].(string),
		Token: req["token"].(string),
	}

	res, err := dbops.CreateUser(newUser)
	if err != nil {
		log.Println("this is a handler error ", res)
		log.Fatal(err)
	}

	var result = map[string]string{"result":"success", "name":req["user_name"].(string),"type":"CreateUser"}
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
	fmt.Println(r.RemoteAddr)
	feedbackWords := dbops.TestDbConnet()
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, feedbackWords)
}
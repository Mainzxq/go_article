package defs

import "github.com/dgrijalva/jwt-go"

// requests
type UserCredential struct {
	// 反点代表tag，解析时会自动转换json
	UID string `json:"uid"`
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Token string `json:"token"`
}

// data model

type ArticleInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
	Contents string
	Title string
}

type Comment struct {
	id string
	ArticleId string
	AuthorId int
	content string
	DisplayCtime string
}

type SimpleSession struct {
	UserName string
	TTL int64
}


type SignClaim struct {
	User string `json:"user"`
	Addr string `json:"addr"`
	AuthLevel int `json:"auth_lv"`
	SystemLevel int `json:"system_lv"`
	SecurityLevel int `json:"security_lv"`
	jwt.StandardClaims
}
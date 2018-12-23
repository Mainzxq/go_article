package defs

// requests
type UserCredential struct {
	// 反点代表tag，解析时会自动转换json
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
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
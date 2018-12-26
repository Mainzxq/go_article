package dbops

import (
	"context"
	"fmt"
	"github.com/Mainzxq/go_article/defs"
	"log"
)

//var dbClient mongo.Client
//  连接数据库


func TestDbConnet() string {
	err = dbClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
		return "Not good, Ops!"
	}
	return "DB connected!"
}


func CreateUser(newUser defs.UserCredential) (bool, error) {
	collection := dbClient.Database("mainzxq").Collection("user_info")
	res, err := collection.InsertOne(context.TODO(), newUser)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	fmt.Println("Insert one doc by ID:", res.InsertedID)
	return true, nil
}
//// 创建用户
//func AddUserCredential(loginName string, pwd string) error {
//	stmtIns,err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES(?, ?)")
//	if err != nil{
//		return err
//	}
//	// 处理执行错误的报错
//	_, err = stmtIns.Exec(loginName, pwd)
//	if err != nil {
//		return err
//	}
//	// 无论何时都会关闭
//	defer stmtIns.Close()
//	return nil
//}
//
//
//// 获取用户名信息
//func GetUserCredential(loginName string) (string, error) {
//	stmtOut, err := dbConn.Prepare("SELECT  pwd FROM users WHERE login_name = ?")
//	if err != nil {
//		log.Printf("%s", err)
//		return "", err
//	}
//	var pwd string
//
//	err = stmtOut.QueryRow(loginName).Scan(&pwd)
//	// 处理执行错误报错
//	// 注意sql.ErrNoRows代表查询结果为空，并非报错类型
//	if err != nil && err != sql.ErrNoRows {
//		return "", err
//	}
//
//	defer stmtOut.Close()
//
//	return pwd, nil
//}
//
//// 删除用户
//func DeleteUser(loginName string, pwd string) error {
//	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
//	if err != nil {
//		log.Printf("DeleteUser error: %s", err)
//		return  err
//	}
//	// 同上
//	_, err = stmtDel.Exec(loginName, pwd)
//	if err != nil {
//		return err
//	}
//	defer stmtDel.Close()
//
//	return nil
//}
//
//
//func AddNewArticle(aid int, name string) (*defs.ArticleInfo, error) {
//	// create uuid
//
//	atid, err := utils.MakeUUIDS()
//	if err != nil {
//		return nil, err
//	}
//
//	t := time.Now()
//	ctime := t.Format("Jan 02 2006, 15:04:05")
//	stmtIns, err := dbConn.Prepare(`INSERT INTO article_info(id, author_id, name, display_time) VALUES(?, ?, ?, ?)`)
//	if err != nil {
//		return nil, err
//	}
//
//	_, err = stmtIns.Exec(atid, aid, name, ctime)
//	if err != nil {
//		return nil, err
//	}
//
//	res := &defs.ArticleInfo{
//		Id:atid,
//		AuthorId: aid,
//		Name: name,
//		DisplayCtime: ctime,
//	}
//
//	defer dbConn.Close()
//
//	return res, nil
//}
//
//func DelArticle(atid string, aid int) error {
//	stmtDell, err := dbConn.Prepare("DELETE FROM article_info WHERE article_id=? AND aid=?" )
//	if err != nil {
//		log.Printf("DeleteUser error: %s", err)
//		return err
//	}
//
//	_, err = stmtDell.Exec(atid, aid)
//	if err != nil {
//		log.Printf("got an excutive error: %s", err)
//		return err
//	}
//	defer stmtDell.Close()
//	return nil
//}
//
//func GetAticlesByAid(aid int) ([]*defs.ArticleInfo, error) {
//	stmtGet, err := dbConn.Prepare("SELECT id, name, display_ctime, contents, title FROM article_info WHERE author_id=?")
//	if err != nil {
//		log.Printf("failed to connect to db: %s", err)
//		return _, err
//	}
//
//	var res []*defs.ArticleInfo
//	rows, err := stmtGet.Query(aid)
//	if err != nil {
//		return res, nil
//	}
//
//	for rows.Next() {
//		var id, name, contents, title, display_ctime string
//		if err := rows.Scan(&id, &name, &display_ctime, &contents, &title); err != nil {
//			return res, err
//		}
//
//		c := &defs.ArticleInfo{
//			Id: id,
//			Name: name,
//			DisplayCtime: display_ctime,
//			Contents: contents,
//			Title: title,
//		}
//
//		res = append(res, c)
//	}
//
//	defer stmtGet.Close()
//	return res, nil
//}
//
//func GetArticles(from int, to int,) ([]*defs.ArticleInfo, error) {
//	stmtOut, err := dbConn.Prepare("SELECT id, name, display_ctime, contents, title FROM article_info WHERE ctime < ? AND ctime >= ?")
//	if err != nil {
//		return _, err
//	}
//
//	var res []*defs.ArticleInfo
//	rows, err := stmtOut.Query(to, from)
//	if err != nil {
//		return res, err
//	}
//
//	for rows.Next() {
//		var id, name, display_ctime, contents, title string
//		if err := rows.Scan(&id, &name, &display_ctime, &contents, &title); err!=nil {
//			return  _, err
//		}
//
//		c := &defs.ArticleInfo{
//			Id: id,
//			Name: name,
//			DisplayCtime: display_ctime,
//			Contents: contents,
//			Title: title,
//		}
//		res = append(res, c)
//	}
//
//	defer stmtOut.Close()
//	return res, nil
//}
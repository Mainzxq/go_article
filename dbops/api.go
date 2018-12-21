package dbops

import (
	"database/sql"
	"github.com/Mainzxq/go_article/defs"
	"github.com/Mainzxq/go_article/utils"
	"log"
	"time"
)

//  连接数据库


// 创建用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns,err := dbConn.Prepare("INSERT INTO users (login_name, pwd) VALUES(?, ?)")
	if err != nil{
		return err
	}
	// 处理执行错误的报错
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	// 无论何时都会关闭
	defer stmtIns.Close()
	return nil
}


// 获取用户名信息
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("SELECT  pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string

	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	// 处理执行错误报错
	// 注意sql.ErrNoRows代表查询结果为空，并非报错类型
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}

	defer stmtOut.Close()

	return pwd, nil
}

// 删除用户
func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name=? AND pwd=?")
	if err != nil {
		log.Printf("DeleteUser error: %s", err)
		return  err
	}
	// 同上
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()

	return nil
}


func AddArticle(aid int, name string) (*defs.ArticleInfo, error) {
	// create uuid

	atid, err := utils.MakeUUIDS()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")
	stmtIns, err := dbConn.Prepare(`INSERT INTO article_info(id, author_id, name, display_time) VALUES(?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}

	_, err = stmtIns.Exec(atid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.ArticleInfo{
		Id:atid,
		AuthorId: aid,
		Name: name,
		DisplayCtime: ctime,
	}

	defer dbConn.Close()

	return res, nil
}
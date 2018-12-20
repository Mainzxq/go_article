package dbops

import (
	"api/defs"
	"database/sql"
	"log"
)

//  连接数据库


// 创建用户
func AdduserCredential(loginName string, pwd string) error {
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

}
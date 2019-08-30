package main

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// openDB 打开数据库
func openDB() *sql.DB {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(db, err)
	return db
}

// initDB 初始化数据库
func initDB() {
	//创建表
	sqlTable := `
			CREATE TABLE IF NOT EXISTS userinfo(
					uid INTEGER PRIMARY KEY AUTOINCREMENT,
					username VARCHAR(64) NULL,
					city VARCHAR(64) NULL,
					skills VARCHAR(128) NULL,
					created DATE NULL
			);
			`
	db.Exec(sqlTable)
}

func insert(u UserModel) (sql.Result, error) {
	stmt, err := db.Prepare("INSERT INTO userinfo(username, city, skills, created) values(?,?,?,?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(u.username, u.city, u.skills, time.Now().Format("2006-01-02 15:04:05"))
	if err != nil {
		return nil, err
	}
	return res, nil
}
func delete(id int64) int64 {
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(stmt, err)

	res, err := stmt.Exec(id)
	checkErr(res, err)

	affect, err := res.RowsAffected()
	checkErr(affect, err)

	return affect
}

// update	更新用户技能
func update(id int, u UserModel) int64 {
	stmt, err := db.Prepare("update userinfo set skills=? where uid=?")
	checkErr(stmt, err)
	res, err := stmt.Exec(u.skills, id)
	checkErr(res, err)

	affect, err := res.RowsAffected()
	checkErr(affect, err)
	return affect
}
func query() ([]UserModel, error) {
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(rows, err)
	var userList = []UserModel{}
	for rows.Next() {
		var user = UserModel{}
		err = rows.Scan(&user.uid, &user.username, &user.city, &user.skills, &user.created)
		checkErr(nil, err)
		userList = append(userList, user)
	}
	rows.Close()
	return userList, nil
}

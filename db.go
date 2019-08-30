package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// openDB 打开数据库
func openDB() *sql.DB {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)
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

func insert() (sql.Result, error) {
	stmt, err := db.Prepare("INSERT INTO userinfo(username, city, skills, created) values(?,?,?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec("谷中仁", "西安", "golang", "2017-04-21")
	checkErr(err)
	return res, nil
}
func delete() {

}
func update() {

}
func query() {

}

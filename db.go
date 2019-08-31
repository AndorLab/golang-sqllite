package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// openDB 打开数据库
func openDB() *sql.DB {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./fcc.db")
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
					created BIGINT NULL
			);
			`
	db.Exec(sqlTable)
}

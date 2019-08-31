package main

import (
	"database/sql"
	"fmt"
)

var db = new(sql.DB)

func main() {
	startApp()
}

func startApp() {
	db = openDB()
	defer db.Close()
	initDB()

	user := UserModel{
		username: "谷中仁",
		city:     `西安`,
		skills:   `js`,
	}
	// insert
	result, err := user.insert()
	id, err := result.LastInsertId()
	checkErr(id, err)
	fmt.Println("操作数据的id:", id)
	// // update
	updateUser := UserModel{
		username: "test",
		city:     `xian`,
		skills:   `golang`,
	}
	affectedRow := updateUser.update(1)
	fmt.Println("影响的行数：", affectedRow)
	// // query
	queryUser := UserModel{}
	list, _ := queryUser.query()
	fmt.Printf("%v", list)
	// delete
	affect := queryUser.delete(4)
	fmt.Println(affect)
}

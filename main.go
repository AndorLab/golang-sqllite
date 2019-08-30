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

	// insert
	result, err := insert()
	id, err := result.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	// // update
	// stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec("wangshubo_new", id)
	// checkErr(err)

	// affect, err := res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)

	// // query
	// rows, err := db.Query("SELECT * FROM userinfo")
	// checkErr(err)
	// var uid int
	// var username string
	// var department string
	// var created time.Time

	// for rows.Next() {
	// 	err = rows.Scan(&uid, &username, &department, &created)
	// 	checkErr(err)
	// 	fmt.Println(uid)
	// 	fmt.Println(username)
	// 	fmt.Println(department)
	// 	fmt.Println(created)
	// }

	// rows.Close()

	// // delete
	// stmt, err = db.Prepare("delete from userinfo where uid=?")
	// checkErr(err)

	// res, err = stmt.Exec(id)
	// checkErr(err)

	// affect, err = res.RowsAffected()
	// checkErr(err)

	// fmt.Println(affect)
}

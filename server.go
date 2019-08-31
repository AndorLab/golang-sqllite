package main

import (
	"github.com/labstack/gommon/log"
)

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
	log.Info("增：操作数据的id:", id)
	// update
	updateUser := UserModel{
		username: "test",
		city:     `xian`,
		skills:   `golang`,
	}
	affectedRow := updateUser.update(1)
	log.Info("改：影响的行数：", affectedRow)
	// query
	queryUser := UserModel{}
	list, _ := queryUser.query()
	log.Info("查：", list)
	// delete
	affect := queryUser.delete(4)
	log.Info("删：", affect)
}

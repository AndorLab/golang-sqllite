package main

import (
	"database/sql"
	"time"
)

// UserModel 用户模型
type UserModel struct {
	uid      int64
	username string
	city     string
	skills   string
	created  int64
}

func (u UserModel) insert() (sql.Result, error) {
	stmt, err := db.Prepare("INSERT INTO userinfo(username, city, skills, created) values(?,?,?,?)")
	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(u.username, u.city, u.skills, time.Now().Unix())
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (u UserModel) delete(id int64) int64 {
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(stmt, err)

	res, err := stmt.Exec(id)
	checkErr(res, err)

	affect, err := res.RowsAffected()
	checkErr(affect, err)

	return affect
}

// update	更新用户技能
func (u UserModel) update(id int) int64 {
	stmt, err := db.Prepare("update userinfo set skills=? where uid=?")
	checkErr(stmt, err)
	res, err := stmt.Exec(u.skills, id)
	checkErr(res, err)

	affect, err := res.RowsAffected()
	checkErr(affect, err)
	return affect
}
func (u UserModel) query() ([]UserModel, error) {
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

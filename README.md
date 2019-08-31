# 1. Golang with SQLLite Practice
<!-- TOC -->

- [Golang with SQLLite Practice](#golang-with-sqllite-practice)
  - [简介](#简介)
  - [目标](#目标)
  - [目的](#目的)
  - [Coding](#coding)
    - [目录结构](#目录结构)
    - [封装 error 函数](#封装-error-函数)
    - [安装 SQLLite 库及其他库](#安装-sqllite-库及其他库)
    - [申明 DB 全局变量](#申明-db-全局变量)
    - [初始化数据库](#初始化数据库)
    - [用户模型构建及原子操作](#用户模型构建及原子操作)
      - [用户模型](#用户模型)
      - [新增](#新增)
      - [删除](#删除)
      - [修改](#修改)
      - [查询](#查询)
    - [在应用中启动并调用用户模型的方法](#在应用中启动并调用用户模型的方法)
    - [运行结果展示](#运行结果展示)
  - [总结](#总结)

<!-- /TOC -->

## 1.1. 简介

SQLite 是一个进程内的库，实现了自给自足的、无服务器的、零配置的、事务性的 SQL 数据库引擎。它是一个零配置的数据库，这意味着与其他数据库一样，你不需要在系统中配置。在 Golang 中使用SQLLite 也相当简单，只需要安装 SQLLite 的Golang  包即可使用；
Golang 就不多介绍了，能看到这个肯定对 Golang 有一定的了解。

仓库地址：<https://github.com/AndorLab/golang-sqllite>

## 1.2. 目标

使用 SQLLite 通过构建一个社区用户表，包含如下字段; 通过 SQLLite 的 API 实现对社区用户表进行增删改查。

| 序号 | 字段     | 类型   | 说明     |
| ---- | -------- | ------ | -------- |
| 1    | uid      | int64  | id       |
| 2    | username | string | 用户名   |
| 3    | city     | string | 城市     |
| 4    | skills   | string | 技能     |
| 5    | created  | int64  | 创建时间 |

## 1.3. 目的

了解 SQLLite ，学习 Golang 操作 SQLLite, 巩固 Golang 基础知识。

## 1.4. Coding

### 1.4.1. 目录结构

项目采用 Golang 传统的平铺式目录

```shell
.
├── LICENSE
├── Makefile      # 构建工具
├── README.md     # README
├── db.go         # 数据库操作
├── error.go      # 错误处理工具方法
├── fcc.db        # sqllite 数据库
├── go.mod        # go modules
├── go.sum        # go modules
├── main.go       # 项目入口
├── server.go     # 应用程序入口
└── userModel.go  # 用户模型

```

### 1.4.2. 封装 error 函数

因为在 go 中会有很多的 error 的判断，为了代码精简，我们特封装一下 error; 下面的 *interface{}* 代表任何类型，类似 TypeScript 中的 *any*。

```golang
# error.go
func checkErr(data interface{}, err error) (interface{}, error) {
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return data, err
}
```

### 1.4.3. 安装 SQLLite 库及其他库

使用 go modules 之后，将所需的包放在 import 中，使用 *go mod tidy* 命令后，go 会自动安装程序使用到的包。

 日志相关的库，主要用于在控制台打印结果
```golang
# server.go
import (
	"github.com/labstack/gommon/log"
)

```
SQLLite 包
```golang
# db.go
_ "github.com/mattn/go-sqlite3"
```

### 1.4.4. 申明 DB 全局变量

因为在程序中，我们要通过数据库来获取数据，那么存在一个全局的数据库指针是很有必要的。

```golang
# main.go
var db = new(sql.DB)
```

### 1.4.5. 初始化数据库

SQLLite 初始化数据库非常简单，只要指定数据库驱动和数据库文件就可以。为了在程序的整个生命周期中操作数据库，我们将 db 返回。

```golang
// openDB 打开数据库
func openDB() *sql.DB {
	//打开数据库，如果不存在，则创建
	db, err := sql.Open("sqlite3", "./fcc.db")
	checkErr(db, err)
	return db
}
```
创建好 db 后，需要创建表结构，执行如下数据库操作命令即可完成用户表的创建。
```golang
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
```

### 1.4.6. 用户模型构建及原子操作

构建现代程序，强调程序的健壮性，封装就是比较重要的；用 MVC、 MVVM 的观点，我们需要有一个 Model 来提供对象的原子操作。在这，我们将用户抽象为UserModel，对用户的增删改查封装到 *insert*、*dleete*、*update* 和 *query*。

#### 1.4.6.1. 用户模型

```golang
// UserModel 用户模型
type UserModel struct {
	uid      int64
	username string
	city     string
	skills   string
	created  int64
}

```
对用户的原子操作

#### 1.4.6.2. 新增

```golang
// insert 新增
func (u UserModel) insert() (sql.Result, error) {
	stmt, err := db.Prepare("insert into userinfo(username, city, skills, created) values(?,?,?,?)")
	checkErr(stmt, err)
	res, err := stmt.Exec(u.username, u.city, u.skills, time.Now().Unix())
	checkErr(res, err)
	return res, nil
}
```

#### 1.4.6.3. 删除

```golang
// delete 删除
func (u UserModel) delete(id int64) int64 {
	stmt, err := db.Prepare("delete from userinfo where uid=?")
	checkErr(stmt, err)
	res, err := stmt.Exec(id)
	checkErr(res, err)
	affect, err := res.RowsAffected()
	checkErr(affect, err)
	return affect
}
```

#### 1.4.6.4. 修改

```golang
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
```

#### 1.4.6.5. 查询

```golang
// query 查询
func (u UserModel) query() ([]UserModel, error) {
	rows, err := db.Query("select * from userinfo")
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
```

### 1.4.7. 在应用中启动并调用用户模型的方法

在上面我们完成了对用户模型及原子操作的封装，那么接下来就是通过应用程序将分装的内容调用，传入正确的参数进行调用。
我们在此封装一个 *startAPP* 方法，在这个里面我们调用封装好的用户操作的接口，实现功能。

因为数据库要在整个生命周期存在，当程序结束的时候，我们应该将数据库链接释放，所以我们用到了 go 的 *defer* 关键字
```golang
# server.go
  db = openDB()
  defer db.Close()
  initDB()
```

调用用户操作的增删改查并打印结果, 对于不同的操作，我们应该有不同的数据，所以在程序中会有 *user*、和 *updateUser* 两个对象

```golang
# server.go
  user := UserModel{
  	username: "谷中仁",
  	city:     `西安`,
  	skills:   `TypeScript`,
  }
  // insert
  result, err := user.insert()
  id, err := result.LastInsertId()
  checkErr(id, err)
  log.Info("增：操作数据的id:", id)
  // update
  updateUser := UserModel{
  	skills: `golang`,
  }
  affectedRow := updateUser.updateSkills(1)
  log.Info("改：影响的行数：", affectedRow)
  // query
  queryUser := UserModel{}
  list, _ := queryUser.query()
  log.Info("查：", list)
  // delete
  affect := queryUser.delete(1)
  log.Info("删：", affect)  
  // query
  list, _ = queryUser.query()
  log.Info("查：", list)
```

### 1.4.8. 运行结果展示

```shell
$ make run
go run *.go
{"time":"2019-08-31T14:21:48.941164+08:00","level":"INFO","prefix":"-","file":"server.go","line":"21","message":"增：操作数据的id:1"}
{"time":"2019-08-31T14:21:48.941842+08:00","level":"INFO","prefix":"-","file":"server.go","line":"27","message":"改：影响的行数：1"}
{"time":"2019-08-31T14:21:48.942034+08:00","level":"INFO","prefix":"-","file":"server.go","line":"31","message":"查：[{1 谷中仁 西安 golang 1567232508}]"}
{"time":"2019-08-31T14:21:48.942599+08:00","level":"INFO","prefix":"-","file":"server.go","line":"34","message":"删：1"}
{"time":"2019-08-31T14:21:48.942696+08:00","level":"INFO","prefix":"-","file":"server.go","line":"38","message":"查：[]"}
```

## 1.5. 总结

SQLLite 对开发者非常友好，不用安装在机器上，只要指定SQLLite的驱动和数据库存储文件即可对 SQLLite 数据库进行操作；Golang 作为比较流行的语言，对数据库也非常友好，提供了基本的数据库接口，
至于用户需要什么样的数据库，自己开发对应的数据库驱动即可。当然在 GitHub 已经有很多开源爱好者开发了比较流行的数据库的驱动可以直接拿来用。

SQLLite 使用的也是标准的 SQL 语法，可以让不同的开发者快速入手。

为什么没有用到 Golang 的 Web 框架？

因为我们的侧重点在 Golang 与 SQLLite，不在 API 实现上，最小化的实现目标，才是我们学习知识最快速的途径。

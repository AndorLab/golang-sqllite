# 1. Golang with SQLLite Practice

<!-- TOC -->

- [1. Golang with SQLLite Practice](#1-golang-with-sqllite-practice)
  - [1.1. 简介](#11-%e7%ae%80%e4%bb%8b)
  - [1.2. 目标](#12-%e7%9b%ae%e6%a0%87)
  - [1.3. 目的](#13-%e7%9b%ae%e7%9a%84)
  - [1.4. Coding](#14-coding)
    - [1.4.1. 目录结构](#141-%e7%9b%ae%e5%bd%95%e7%bb%93%e6%9e%84)
    - [1.4.2. 安装 SQLLite 库及其他库](#142-%e5%ae%89%e8%a3%85-sqllite-%e5%ba%93%e5%8f%8a%e5%85%b6%e4%bb%96%e5%ba%93)
    - [1.4.3. 申明 DB 全局变量](#143-%e7%94%b3%e6%98%8e-db-%e5%85%a8%e5%b1%80%e5%8f%98%e9%87%8f)
    - [1.4.4. 用户模型构建及原子操作](#144-%e7%94%a8%e6%88%b7%e6%a8%a1%e5%9e%8b%e6%9e%84%e5%bb%ba%e5%8f%8a%e5%8e%9f%e5%ad%90%e6%93%8d%e4%bd%9c)
    - [1.4.5. 在应用中启动并调用用户模型的方法](#145-%e5%9c%a8%e5%ba%94%e7%94%a8%e4%b8%ad%e5%90%af%e5%8a%a8%e5%b9%b6%e8%b0%83%e7%94%a8%e7%94%a8%e6%88%b7%e6%a8%a1%e5%9e%8b%e7%9a%84%e6%96%b9%e6%b3%95)
    - [1.4.6. 运行结果展示](#146-%e8%bf%90%e8%a1%8c%e7%bb%93%e6%9e%9c%e5%b1%95%e7%a4%ba)
  - [1.5. 总结](#15-%e6%80%bb%e7%bb%93)

<!-- /TOC -->

## 1.1. 简介

SQLite 是一个进程内的库，实现了自给自足的、无服务器的、零配置的、事务性的 SQL 数据库引擎。它是一个零配置的数据库，这意味着与其他数据库一样，你不需要在系统中配置。在 Golang 中使用SQLLite 也相当简单，只需要安装 SQLLite 的Golang  包即可使用；
Golang 就不多介绍了，能看到这个肯定对 Golang 有一定的了解。

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
├── foo.db        # sqllite 数据库
├── go.mod        # go modules
├── go.sum        # go modules
├── main.go       # 项目入口
├── server.go     # 应用程序入口
└── userModel.go  # 用户模型

```

### 1.4.2. 安装 SQLLite 库及其他库

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

### 1.4.3. 申明 DB 全局变量

因为在程序中，我们要通过数据库来获取数据，那么存在一个全局的数据库指针是很有必要的。

```golang
# main.go
var db = new(sql.DB)
```

### 1.4.4. 用户模型构建及原子操作

### 1.4.5. 在应用中启动并调用用户模型的方法

### 1.4.6. 运行结果展示

## 1.5. 总结

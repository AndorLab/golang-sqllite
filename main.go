package main

import (
	"database/sql"
)

var db = new(sql.DB)

func main() {
	startApp()
}

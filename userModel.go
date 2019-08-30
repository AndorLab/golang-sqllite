package main

import (
	"time"
)

// UserModel 用户模型
type UserModel struct {
	uid      int64
	username string
	city     string
	skills   string
	created  time.Time
}

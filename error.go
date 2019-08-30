package main

import (
	log "github.com/sirupsen/logrus"
)

func checkErr(err error) {
	if err != nil {
		log.Error(err)
		panic("操作失败")
	}
}

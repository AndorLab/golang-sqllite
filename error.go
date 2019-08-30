package main

import (
	log "github.com/sirupsen/logrus"
)

func checkErr(data interface{}, err error) (interface{}, error) {
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return data, err
}

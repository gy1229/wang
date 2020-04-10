package test

import (
	"github.com/gy1229/wang/database"
	"github.com/gy1229/wang/util"
)

func InitTestConfig1() {
	util.InitViper1()
	util.InitID()
	database.InitDB()
	//gredis.Setup()
}

func InitTestConfig2() {
	util.InitViper2()
	util.InitID()
	database.InitDB()
	//gredis.Setup()
}

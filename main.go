package main

import (
	"go-react/mysql"
	"go-react/router"
)

func main() {
	r := router.Init()
	mysql.InitMysql()
	r.Run(":8080")
}

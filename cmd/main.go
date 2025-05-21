package main

import (
	"net/http"
	"serverDemo/database"
	"serverDemo/router"
)

func main() {
	//初始化数据库
	err := database.InitMysql()
	if err != nil {
		panic(err)
	}

	router := router.SetupRoutes()
	http.ListenAndServe(":8080", router)
}

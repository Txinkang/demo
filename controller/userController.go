package controller

import (
	"fmt"
	"net/http"
	"serverDemo/database"
)

func Login(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		http.Error(response, "方法不允许", http.StatusMethodNotAllowed)
		return
	}

	user_account := request.URL.Query().Get("user_account")
	if user_account != "" {
		rows, err := database.DB.Query("select user_account from user where user_account = ?", user_account)
		if err != nil {
			http.Error(response, "数据库查询有误", http.StatusInternalServerError)
		}
		defer rows.Close()

		for rows.Next() {
			var username string
			rows.Scan(&username)
			fmt.Printf("用户：%s\n", username)
		}
	} else {
		http.Error(response, "账号密码不正确", http.StatusUnauthorized)
	}
}

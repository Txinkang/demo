package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitMysql() error {
	dsn := "root:root@tcp(127.0.0.1:3306)/ch?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	DB = db
	if err := DB.Ping(); err != nil {
		return err
	}
	fmt.Println("MySQL连接成功")
	return nil
}

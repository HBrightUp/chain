package main

import (
	"database/sql"
	"fmt"
	"strings"
)

const (
	userName = "root"
	password = "123456"
	ip       = "localhost"
	port     = "3306"
	dbName   = "go_test"
)

var db *sql.DB

func initDB() error {
	//构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
	path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")

	//打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
	db, _ = sql.Open("mysql", path)
	//设置数据库最大连接数
	db.SetConnMaxLifetime(100)
	//设置上数据库最大闲置连接数
	db.SetMaxIdleConns(10)
	//验证连接
	if err := db.Ping(); err != nil {
		fmt.Println("open database fail")
		return err
	}
	fmt.Println("connect success")
	return nil
}

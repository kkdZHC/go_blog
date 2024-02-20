package dao

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:028135Cc@tcp(127.0.0.1:3306)/goblog?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Println("连接数据库异常")
		panic(err)
	}
	//最大空闲连接数，默认不配置，是2个最大空闲连接
	db.SetMaxIdleConns(5)
	// 连接最大存活时间
	db.SetConnMaxLifetime(time.Minute * 5)
	//空闲连接最大存活时间
	db.SetConnMaxIdleTime(time.Minute * 2)
	err = db.Ping()
	if err != nil {
		log.Println("数据库连接失效")
		db.Close()
		panic(err)
	}
	DB = db
}

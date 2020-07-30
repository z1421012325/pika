package db

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm"
	_ "gorm/dialects/mysql"
	"os"
	"time"
)

var SDB *gorm.DB

func NewMySqlConn() {

	name := os.Getenv("MYSQL_NAME")
	pswd := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	dbname := os.Getenv("MYSQL_DB")

	db, err := gorm.Open("mysql",
		fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", name, pswd, host, port, dbname))
	if err != nil {
		panic(err.Error())
	}

	// 开发模式和生产模式日志开关,环境
	debug := os.Getenv("GIN_MODE")
	if gin.DebugMode != debug {
		db.LogMode(false)
	} else {
		db.LogMode(true)
	}

	// 设置可重用连接的最大时间量 如果d<=0，则永远重用连接
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//设置到数据库的最大打开连接数 如果n<=0，则不限制打开的连接数 默认值为0
	db.DB().SetMaxOpenConns(0)
	// 设置空闲中的最大连接数 默认最大空闲连接数当前为2 如果n<=0，则不保留空闲连接
	db.DB().SetMaxIdleConns(10)

	SDB = db
}

/*
	开启 mysql 事务操作
	支持一次传递多个 *gorm.DB 执行语句(exce)
*/
func Transaction(dbs ...*gorm.DB) bool {

	tx := SDB.Begin()

	for _, db := range dbs {
		tx = db
		if tx.Error != nil {
			tx.Rollback()
			return false
		}
	}

	tx.Commit()
	return true
}

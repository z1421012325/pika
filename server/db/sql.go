package db

import (
	"gorm"
	_ "gorm/dialects/mysql"
)

const (
	defaultHost = ""		// 默认ip
	defaultPort = ""		// 默认端口
	defaultMaxConn = 20		// 默认连接池最大数
)

var SDB *gorm.DB


func NewMySqlConn(host,port string,maxConn int){

	if host == "" {
		host = defaultHost
	}

	if port == "" {
		port = defaultPort
	}

	if maxConn == 0 {
		maxConn = defaultMaxConn
	}

	// logic

	AddressStr := host + port 	// todo
	//var err error
	conn,err := gorm.Open(AddressStr)
	if err != nil {
		panic(err)
	}

	// set setting outtime,maxconn,maxopen
	//conn.DB().SetConnMaxLifetime()
	//conn.DB().SetMaxIdleConns()
	//conn.DB().SetMaxOpenConns()

	SDB = conn

}



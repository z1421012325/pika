package db

import "redigo/redis"

const (
	DefaultRedisAddres = ""		// redis 默认地址
	DefaultRedisMaxConn = 50  	// redis 连接池最大数
)


var RDB *redis.Pool



// 被db/init.go引用
func NewRedisPool(address string,maxConn int){

	if address == ""{
		address = DefaultRedisAddres
	}

	if maxConn == 0 {
		maxConn = DefaultRedisMaxConn
	}


	// logic
}
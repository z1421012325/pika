package db

// db init
func init() {
	NewRedisPool()
	NewMySqlConn()
}

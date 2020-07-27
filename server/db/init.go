package db



// db init
func init() {
	NewRedisPool("",0)
	NewMySqlConn("","",0)
}
package main

import (
	"encoding/json"
	"fmt"
	"time"

	"gorm"
	_ "gorm/dialects/mysql"
)

var (
	conn *gorm.DB
)

func init() {
	db,err := gorm.Open("mysql",
		"root:zyms90bdcs@tcp(localhost:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}

	// 开发模式和生产模式日志开关,环境
	//debug := os.Getenv("GIN_MODE")
	//if gin.DebugMode != debug {
	//	db.LogMode(false)
	//}else {
	//	db.LogMode(true)
	//}

	// 设置可重用连接的最大时间量 如果d<=0，则永远重用连接
	db.DB().SetConnMaxLifetime(time.Second * 30)
	//设置到数据库的最大打开连接数 如果n<=0，则不限制打开的连接数 默认值为0
	db.DB().SetMaxOpenConns(0)
	// 设置空闲中的最大连接数 默认最大空闲连接数当前为2 如果n<=0，则不保留空闲连接
	db.DB().SetMaxIdleConns(10)

	db.LogMode(true)

	conn = db
}

type a struct {
	Aint string `json:"astring"`
	b		`json:"b_struct"`
	c		`json:"c_struct"`
	DD []d `json:"dd_slice"`
}
type b struct {
	Bint string `json:"b_string"`
}
type c struct {
	Cint string `json:"c_string"`
}
type d struct {
	Dint string `json:"d_string"`
}





// 测试读取配置文件
func main() {

	//s1 := "insert into curriculums (c_name,u_id,price) values(?,?,?)"
	//s2 := "set @id = LAST_INSERT_ID()"
	//s3 := "insert into catalog (c_id,name,url) values (@id,?,?)"
	//
	//var c Curriculum
	//
	//link := conn.Exec(s1,"golang语言" , 1 , 123.99).Exec(s2).Exec(s3,"golang语言test","https://www.xxx.com")
	//link.Raw("select @id").Scan(&c)
	//
	//
	//fmt.Println("@id : ",c)

	//aa := a{
	//	Aint: "a",
	//	b:    b{Bint:"b"},
	//	c:    c{Cint:"c"},
	//	DD:   []d{
	//		{Dint:"d1"},
	//		{Dint:"d2"},
	//		{Dint:"d3"},
	//	},
	//}
	//
	//b,_ := json.Marshal(aa)
	//fmt.Println(string(b))

	var res Curriculums

	sql := "insert into curriculums (c_name,u_id,price) values(?,?,?)"
	sql2 := "set @id = LAST_INSERT_ID()"
	sql3 := "select * from curriculums where c_id = @id"
	conn.Exec(sql,"java语言" , 1 , 3.99).Exec(sql2).Raw(sql3).Scan(&res)

	b , _ := json.Marshal(res)
	fmt.Println(string(b))
}


type Curriculums struct {
	CID 		int			`gorm:"column:c_id" json:"cid"`
	UID 		int			`gorm:"column:u_id" json:"uid"`
	//TID 		int			`gorm:"column:t_id" json:"tid"`
	Name 		string		`gorm:"column:c_name" json:"name"`
	Price 		float64		`gorm:"column:price" json:"price"`
	Info		string		`gorm:"column:info" json:"info"`
	Image		string		`gorm:"column:c_image" json:"img"`
	CreateTime 	*time.Time	`gorm:"column:create_at" json:"at"`
	DeleteTime 	*time.Time	`gorm:"column:delete_at" json:"et"`

	AdminDelTime  	*time.Time		`gorm:"column:admin_del" json:"a_del"`		// 后台人员删除时间
	Aid       		int				`gorm:"column:a_id" json:"aid"`				// 后台执行人信息
}


func (cr *Curriculums)TableName()string{
	return "curriculums"
}

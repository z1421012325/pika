package models

import (
	"time"

	"pika/server/db"
	"pika/tools"
)

const (
	// 站主,管理员,上传用户,普通用户
	Level1 = iota + 1
	Level2
	Level3
	Level4

	Registry_User_Pass = -1		// 普通用户不需要审核直接pass,未通过审核数据库中字段默认显示为null
)

// 用户表

/*
create table `users` (
`u_id` INT AUTO_INCREMENT PRIMARY KEY,
`nickname` VARCHAR(20) NOT NULL,
`username` VARCHAR(20) NOT NULL,
`pswd` VARCHAR(255) NOT NULL,
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '用户创建时间',
`grade_level` INT NOT NULL COMMENT '用户等级 grade表外键',
`verify_at` DATETIME COMMENT '上传人账号通过注册时间',
`exec_pass` INT COMMENT '验证账号通过执行人id  ',
`delete_id` INT COMMENT '删除账号执行人id',
`delete_at` DATETIME,
UNIQUE KEY `_username` (`username`),
CONSTRAINT `_grade` FOREIGN KEY (`grade_level`) REFERENCES `grade` (`g_id`)
)ENGINE=InnoDB AUTO_INCREMENT=100000 DEFAULT CHARSET=utf8;
*/

type User struct {
	Uid        int       `gorm:"column:u_id" json:"u_id,omitempty"`
	NickName   string    `gorm:"column:nickname" json:"nickname,omitempty"`
	Username   string    `gorm:"column:username" json:"username,omitempty" form:"username" binding:"required,min=2,max=20"`
	PassWord   string    `gorm:"column:pswd" json:"password,omitempty" form:"password" binding:"required,min=5,max=30"`
	CreateAt   time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt   time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`

	GradeLevel int       `gorm:"column:grade_level" json:"grade_level,omitempty" form:"grade" binding:"required,oneof=1 2 3 4"`

	VerifyAt   time.Time `gorm:"column:verify_at" json:"verify_at,omitempty"`
	ExecPaas   int       `gorm:"column:exec_pass" json:"exec_pass,omitempty"`
	DeleteId   int       `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (u *User) TableName() string {
	return "users"
}






// 查询普通用户
func (u *User)QueryUser(){
	db.SDB.Where("username = ? AND exec_pass = ? ",u.Username,Registry_User_Pass).First(u)
}

// 验证密码
func (u *User)CheckUserPswd(verifyStr string) bool {
	return tools.CheckPassword(u.PassWord,verifyStr)
}

// 查询并验证密码
func (u *User)VerifyUser(verifyStr string) bool {
	u.QueryUser()
	return u.CheckUserPswd(verifyStr)
}

// 密码加密
func (u *User)Encryption(){
	u.PassWord = tools.EnCryptionPassword(u.PassWord)
}


// 检测用户等级并添加数据        pass
func (u *User)CheckUserGrade(){
	// 看看字段grade 的tag oneof = 1 2 3 4 能不嫩过滤不属于执行数据
	//var checkList = []int{Level1,Level2,Level3,Level4}
	//for _,v := range checkList{
	//	if u.GradeLevel == v {
	//		return true
	//	}
	//}
	if u.GradeLevel == Level4 {
		u.ExecPaas = Registry_User_Pass
	}
}

// 注册
func (u *User)RegistryUser(nickname string) error {
	sql := "INSERT INTO users (nickname,username,pswd,grade_level,exec_pass) VALUES (?,?,?,?,?)"
	return db.SDB.Raw(sql,nickname,u.Username,u.PassWord,u.GradeLevel,u.ExecPaas).Error
}


// 查询用户信息
func QueryUserInfo(uid string,in interface{}) error{
	return db.SDB.Where("uid = ? AND delete_at IS NULL ",uid).First(&in).Error
}

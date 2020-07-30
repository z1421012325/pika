package models

import "time"

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
	Username   string    `gorm:"column:username" json:"username,omitempty"`
	PassWord   string    `gorm:"column:pswd" json:"-"`
	CreateAt   time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt   time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
	GradeLevel int       `gorm:"column:grade_level" json:"grade_level,omitempty"`
	VerifyAt   time.Time `gorm:"column:verify_at" json:"verify_at,omitempty"`
	ExecPaas   int       `gorm:"column:exec_pass" json:"exec_pass,omitempty"`
	DeleteId   int       `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (u User) TableName() string {
	return "users"
}

package models

import "time"

// 本子表

/*
CREATE TABLE `benzi` (
`b_id` INT AUTO_INCREMENT PRIMARY KEY,
`title` VARCHAR(100) NOT NULL,
`b_cover` VARCHAR(255) NOT NULL,
`author` VARCHAR(20) DEFAULT NULL,
`upload_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间  ',
`updata_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更行时间  ',
`upload_id` INT NOT NULL COMMENT '上传人id  ',
`delete_id` INT NOT NULL COMMENT '删除id  ',
UNIQUE KEY `_title` (`title`),
UNIQUE KEY `_author` (`author`),
CONSTRAINT `_upload_id` FOREIGN KEY (`upload_id`) REFERENCES `users` (`u_id`),
CONSTRAINT `_delete_id` FOREIGN KEY (`delete_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000  DEFAULT CHARSET=utf8;
*/

type BenZi struct {
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

func (b BenZi) TableName() string {
	return "benzi"
}

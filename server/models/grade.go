package models

// 用户等级表

/*
create table `grade` (
`g_id` INT AUTO_INCREMENT PRIMARY KEY,
`_explain` VARCHAR(20) COMMENT '1,站主 2,管理员 3,上传人员 4,普通用户'
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO grade (_explain) VALUES ("站主");
INSERT INTO grade (_explain) VALUES ("管理员");
INSERT INTO grade (_explain) VALUES ("上传人员");
INSERT INTO grade (_explain) VALUES ("普通用户");
*/

type Grade struct {
	Gid     int    `gorm:"column:g_id" json:"g_id,omitempty"`
	Explain string `gorm:"column:_explain" json:"explain,omitempty"`
}

func (b Grade) TableName() string {
	return "grade"
}

package models

import "time"

// tag

/*
CREATE TABLE `classify_tag` (
`t_id`	INT AUTO_INCREMENT PRIMARY KEY,
`b_id`  INT NUT NULL COMMENT "本子id",
`t_level` INT COMMENT "TAG等级,1代表分类tag,2表示细分tag",
`tag_name` VARCHAR(10) COMMENT "TAG名字",
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME ,
`delete_id` INT COMMENT "这是删除人员id,如果是创建人员id则在一对多表中user_tag ",
CONSTRAINT `_b_id_9` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`),
CONSTRAINT `_u_id_9` FOREIGN KEY (`delete_id`) REFERENCES `users` (`u_id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

const (
	TYPE_INT = 1		// 分类
	TAG_INT  = 2		// 标签
)

type Tags struct {
	Tid      int       `gorm:"column:t_id" json:"t_id,omitempty"`
	Bid   	 int 		`gorm:"column:b_id" json:"b_id,omitempty"`
	TLevel   int       `gorm:"column:t_level" json:"t_level,omitempty"`
	TagName  string    `gorm:"column:tag_name" json:"tag_name,omitempty"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
	DeleteId int       `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b Tags) TableName() string {
	return "classify_tag"
}

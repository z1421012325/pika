package models

import "time"

// 收藏本子
/*
CREATE TABLE `collections` (
`collect_id` INT AUTO_INCREMENT PRIMARY KEY,
`u_id` INT NOT NULL,
`b_id` INT NOT NULL,
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME,
UNIQUE INDEX `_u_id_6` (`u_id`),
CONSTRAINT `_b_id_6` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`),
CONSTRAINT `_u_id_6` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;
*/
type Collection struct {
	CollectId int       `gorm:"column:collect_id" json:"collect_id,omitempty"`
	Bid       int       `gorm:"column:b_id" json:"b_id,omitempty"`
	Uid       int       `gorm:"column:u_id" json:"u_id,omitempty"`
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt  time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b Collection) TableName() string {
	return "collections"
}

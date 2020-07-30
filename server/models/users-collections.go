package models

import "time"

// 用户 收藏
/*
CREATE TABLE `users-collections` (
`u_id` INT NOT NULL,
`collect_id` INT NOT NULL,
UNIQUE INDEX `u_collect_id` (`u_id`,`collect_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

type UsersCollection struct {
	Uid       int `gorm:"column:u_id" json:"u_id,omitempty"`
	CollectID int `gorm:"column:collect_id" json:"collect_id,omitempty"`
}

func (b UsersCollection) TableName() string {
	return "users-collections"
}

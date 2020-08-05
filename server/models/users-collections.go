package models

import (
	"pika/server/db"
	"time"
)

// 用户 收藏
/*
CREATE TABLE `users_collections` (
`u_id` INT NOT NULL,
`b_id` INT NOT NULL,
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME,
CONSTRAINT `_b_id_111` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`),
CONSTRAINT `_u_id_111` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

type UsersCollection struct {
	Uid       int `gorm:"column:u_id" json:"u_id,omitempty"`
	CollectID int `gorm:"column:collect_id" json:"collect_id,omitempty"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b UsersCollection) TableName() string {
	return "users_collections"
}


// 添加收藏
func UserAddBenziCollection(bid int,uid string) error {
	SQL := "INSERT INTO users_collections (b_id,u_id) VALUES (?,?)"
	return db.SDB.Exec(SQL,bid,uid).Error
}
// 添加收藏
func UnUserAddBenziCollection(bid int,uid string) error {
	SQL := "UPDATE users_collections SET delete_id = now() WHERE b_id = ? AND u_id = ?"
	return db.SDB.Exec(SQL,bid,uid).Error
}
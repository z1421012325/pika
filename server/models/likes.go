package models

import (
	"pika/server/db"
	"time"
)

// 点赞
/*
CREATE TABLE `likes` (
`u_id` INT NOT NULL,
`b_id` INT NOT NULL,
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME,
CONSTRAINT `_b_id_7` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`),
CONSTRAINT `_u_id_7` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000 DEFAULT CHARSET=utf8;

*/
type Like struct {
	Uid      int       `gorm:"column:u_id" json:"u_id,omitempty"`
	Bid      int       `gorm:"column:b_id" json:"b_id,omitempty"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b Like) TableName() string {
	return "likes"
}


// 点赞
func UserLikeBenzi(bid int,uid string) error{
	SQL := "INSERT INTO likes (u_id,b_id) VALUES (?,?)"
	return db.SDB.Exec(SQL,uid,bid).Error
}

// 取消点赞
func CancelUserLikeBenzi(bid int,uid string) error{
	SQL := "UPDATE likes SET delete_id = now() WHERE b_id = ? AND u_id = ?"
	return db.SDB.Exec(SQL,bid,uid).Error
}
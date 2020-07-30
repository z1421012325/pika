package models

import "time"

// 点赞
/*
CREATE TABLE `likes` (
`u_id` INT NOT NULL,
`b_id` INT NOT NULL,
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME,
UNIQUE INDEX `_u_id_7` (`u_id`),
UNIQUE INDEX `_b_id_7` (`b_id`),
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

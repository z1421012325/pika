package models

import "time"

// 评论
/*
CREATE TABLE `comments` (
`c_id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '评论id',
`b_id` INT NOT NULL COMMENT '本子id',
`u_id` INT NOT NULL COMMENT '用户id',
`comment` VARCHAR(150) NOT NULL,
`agree` INT DEFAULT 0,
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME,
UNIQUE INDEX `_c_id_4` (`c_id`),
UNIQUE INDEX `_u_id_4` (`u_id`),
UNIQUE INDEX `_b_id_4` (`b_id`),
CONSTRAINT `_b_id_2` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`),
CONSTRAINT `_u_id_2` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/
type Comment struct {
	Cid      int       `gorm:"column:c_id" json:"c_id,omitempty"`
	Bid      int       `gorm:"column:b_id" json:"b_id,omitempty"`
	Uid      int       `gorm:"column:u_id" json:"u_id,omitempty"`
	Comment  string    `gorm:"column:comment" json:"comment,omitempty"`
	Agree 	 int	   `gorm:"column:agree" json:"agree,omitempty"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b Comment) TableName() string {
	return "comments"
}

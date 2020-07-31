package models

import "time"

// 回复评论

/*
CREATE TABLE `reply_comments` (
`reply_id` INT AUTO_INCREMENT PRIMARY KEY COMMENT '回复id',
`c_id` INT COMMENT '一级(顶级)评论id',
`comment` VARCHAR(150) NOT NULL,
`agree` INT DEFAULT 0,
`u_id` INT NOT NULL COMMENT '用户id',
`reply_u_id` INT NOT NULL COMMENT '被回复评论的用户id',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME,
UNIQUE INDEX `_c_id_5` (`c_id`),
CONSTRAINT `_b_id_3` FOREIGN KEY (`c_id`) REFERENCES `comments` (`c_id`),
CONSTRAINT `_u_id_3` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`),
CONSTRAINT `_reply_u_id` FOREIGN KEY (`reply_u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type ReplyComments struct {
	ReplyId  int       `gorm:"column:reply_id" json:"reply_id,omitempty"`
	Cid      int       `gorm:"column:c_id" json:"c_id,omitempty"`
	Comment  string    `gorm:"column:comment" json:"comment,omitempty"`
	Agree 	 int	   `gorm:"column:agree" json:"agree,omitempty"`
	Uid      int       `gorm:"column:u_id" json:"u_id,omitempty"`
	ReplyUid int       `gorm:"column:reply_u_id" json:"reply_uid,omitempty"`
	CreateAt time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b ReplyComments) TableName() string {
	return "reply_comments"
}

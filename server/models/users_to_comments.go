package models

// 多对多
// 用户   评论

/*
CREATE TABLE `users_to_comments` (
`b_id` INT NOT NULL,
`u_id` INT NOT NULL,
UNIQUE INDEX `b_u_id` (`b_id`,`u_id`),
CONSTRAINT `_b_id_1` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`),
CONSTRAINT `_u_id_1` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type UsersToComment struct {
	Uid int `gorm:"column:u_id" json:"u_id,omitempty"`
	Bid int `gorm:"column:b_id" json:"b_id,omitempty"`
}

func (b UsersToComment) TableName() string {
	return "users_to_comments"
}

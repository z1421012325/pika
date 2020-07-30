package models

// 用户创建tag
// 用户  tag

/*
CREATE TABLE `users_tags` (
`u_id` INT NOT NULL,
`t_id` INT NOT NULL,
CONSTRAINT `_u_id_10` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`),
CONSTRAINT `_t_id_10` FOREIGN KEY (`t_id`) REFERENCES `tags` (`t_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

type UserTag struct {
	Uid int `gorm:"column:u_id" json:"u_id,omitempty"`
	Tid int `gorm:"column:u_id" json:"u_id,omitempty"`
}

func (b UserTag) TableName() string {
	return "users_tags"
}

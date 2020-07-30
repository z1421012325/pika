package models

// tag 本子
/*
CREATE TABLE `tag_benzi` (
`u_id` INT NOT NULL,
`b_id` INT NOT NULL,
CONSTRAINT `_u_id_11` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`),
CONSTRAINT `_b_id_11` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type TagBenzi struct {
	Uid int `gorm:"column:u_id" json:"u_id,omitempty"`
	Bid int `gorm:"column:b_id" json:"b_id,omitempty"`
}

func (b TagBenzi) TableName() string {
	return "tag_benzi"
}

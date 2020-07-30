package models

import "time"

// 观看记录
/*
CREATE TABLE `log_record` (
`u_id` INT NOT NULL,
`b_id` INT NOT NULL,
`chapter_id` INT NOT NULL,
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
UNIQUE INDEX `_u_b_chapter_id` (`b_id`,`u_id`,`chapter_id`),
CONSTRAINT `_b_id_8` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`),
CONSTRAINT `_u_id_8` FOREIGN KEY (`u_id`) REFERENCES `users` (`u_id`)

) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type LogRecord struct {
	Uid       int       `gorm:"column:u_id" json:"u_id,omitempty"`
	Bid       int       `gorm:"column:b_id" json:"b_id,omitempty"`
	ChapterId int       `gorm:"column:chapter_id" json:"chapter_id,omitempty"`
	CreateAt  time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
}

func (b LogRecord) TableName() string {
	return "log_record"
}

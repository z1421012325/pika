package models

import (
	"pika/server/db"
	"time"
)

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


// 根据本子id与用户id去查询用户历史记录章节id,  pass > 并在benzi_img表中查询 之后再次更新历史记录
func RecordQueryAndUpRecord(uid string,bid int,in interface{}) error {

	sql1 := "SET @chapter_id = (SELECT chapter_id FROM log_record WHERE u_id = ? AND b_id = ?)"
	sql2 := "SELECT img_id,b_id,b_url,chapter_id FROM benzi_img WHERE chapter_id = @chapter_id ORDER BY img_id ASC"
	return db.SDB.Exec(sql1,uid,bid).Raw(sql2).Scan(&in).Error
}
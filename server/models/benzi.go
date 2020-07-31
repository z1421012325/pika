package models

import "time"

// 本子表

/*
CREATE TABLE `benzi` (
`b_id` INT AUTO_INCREMENT PRIMARY KEY,
`title` VARCHAR(100) NOT NULL,
`b_cover` VARCHAR(255) NOT NULL,
`author` VARCHAR(20) DEFAULT NULL,
`upload_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间  ',
`updata_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更行时间  ',
`delete_at` DATETIME COMMENT '删除时间  ',
`upload_id` INT NOT NULL COMMENT '上传人id  ',
`delete_id` INT NOT NULL COMMENT '删除id  ',
UNIQUE KEY `_title` (`title`),
UNIQUE KEY `_author` (`author`),
CONSTRAINT `_upload_id` FOREIGN KEY (`upload_id`) REFERENCES `users` (`u_id`),
CONSTRAINT `_delete_id` FOREIGN KEY (`delete_id`) REFERENCES `users` (`u_id`)
) ENGINE=InnoDB AUTO_INCREMENT=10000  DEFAULT CHARSET=utf8;
*/

type BenZi struct {
	Bid        int       `gorm:"column:b_id" json:"b_id,omitempty"`
	Title string			`gorm:"column:title" json:"title,omitempty"`
	BCover string			`gorm:"column:b_cover" json:"b_cover,omitempty"`
	Author string			`gorm:"column:author" json:"author,omitempty"`
	UploadAt time.Time		`gorm:"column:upload_at" json:"upload_at,omitempty"`
	UpdateAt time.Time		`gorm:"column:updata_at" json:"updata_at,omitempty"`
	DeleteAt time.Time		`gorm:"column:delete_at" json:"delete_at,omitempty"`
	UploadId int			`gorm:"column:upload_id" json:"upload_id,omitempty"`
	DeleteId int		`gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b BenZi) TableName() string {
	return "benzi"
}

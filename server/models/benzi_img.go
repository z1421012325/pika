package models

import (
	"errors"
	"fmt"
	"pika/server/db"
	"time"
)

// 本子图片
/*
CREATE TABLE `benzi_img` (
`img_id` int AUTO_INCREMENT PRIMARY KEY,
`b_id` INT NOT NULL,
`b_url` VARCHAR(255) NOT NULL,
`chapter_id` INT DEFAULT 0 COMMENT '章节id',
`chapter_name` VARCHAR(10) DEFAULT "第一话" COMMENT '章节名',
`create_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
`delete_at` DATETIME ,
CONSTRAINT `_b_id` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type BenZiImg struct {
	ImgId       int    `gorm:"column:img_id" json:"img_id,omitempty"`
	Bid         int    `gorm:"column:b_id" json:"b_id,omitempty"`
	BUrl        string `gorm:"column:b_url" json:"b_url,omitempty"`
	ChapterId   int    `gorm:"column:chapter_id" json:"chapter_id,omitempty"`
	ChapterName string `gorm:"column:chapter_name" json:"chapter_name,omitempty"`
	CreateAt time.Time		`gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt time.Time		`gorm:"column:delete_at" json:"delete_at,omitempty"`
}

func (b BenZiImg) TableName() string {
	return "benzi_img"
}


// 根据本子id查询该本子id的章节id
func QueryBenziChapterId(bid int) int {
	var tmp_img BenZiImg
	db.SDB.Where("b_id = ? ORDER BY chapter_id DESC",bid).First(&tmp_img)
	if tmp_img.ChapterId == 0 {
		tmp_img.ChapterId++
	}
	return tmp_img.ChapterId
}


// 上传本子本体图片
func UploadBnezi(bid,uid int,imgUrls []string,chapter string) error {

	if !IsBenziInUserUpload(bid,uid){
		return errors.New("上传人异常")
	}

	chapterInt := QueryBenziChapterId(bid)
	InsertSQL := "INSERT INTO benzi_img (b_id,b_url,chapter_id,chapter_name) VALUES "

	for _,ImgUrl := range imgUrls {
		InsertSQL = InsertSQL + fmt.Sprintf("(%d,%s,%d,%s),",bid,ImgUrl,chapterInt,chapter)
	}

	upBenziTime := "UPDATE benzi SET updata_at = now() WHERE b_id = ?"
	return db.SDB.Exec(InsertSQL).Exec(upBenziTime,bid).Error
}
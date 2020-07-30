package models

// 本子图片
/*
CREATE TABLE `benzi_img` (
`img_id` int AUTO_INCREMENT PRIMARY KEY,
`b_id` INT NOT NULL,
`b_url` VARCHAR(255) NOT NULL,
`chapter_id` INT DEFAULT 0 COMMENT '章节id',
`chapter_name` VARCHAR(10) DEFAULT "第一话" COMMENT '章节名',
CONSTRAINT `_b_id` FOREIGN KEY (`b_id`) REFERENCES `benzi` (`b_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type BenZiImg struct {
	ImgId       int    `gorm:"column:img_id" json:"img_id,omitempty"`
	Bid         int    `gorm:"column:b_id" json:"b_id,omitempty"`
	BUrl        string `gorm:"column:b_url" json:"b_url,omitempty"`
	ChapterId   int    `gorm:"column:chapter_id" json:"chapter_id,omitempty"`
	ChapterName string `gorm:"column:chapter_name" json:"chapter_name,omitempty"`
}

func (b BenZiImg) TableName() string {
	return "benzi"
}

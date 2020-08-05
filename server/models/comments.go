package models

import (
	"pika/server/db"
	"time"
)

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
	CreateAt *time.Time `gorm:"column:create_at" json:"create_at,omitempty"`
	DeleteAt *time.Time `gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b Comment) TableName() string {
	return "comments"
}


// 添加评论
func AddBenziComment(uid string,bid int,comment string) error {
	sql := "INSERT INTO comments (b_id,u_id,comment) VALUES (?,?,?)"
	return db.SDB.Exec(sql,bid,uid,comment).Error
}

// 根据本子id查询评论
func QueryBenziComment(bid int,page,number int64,in interface{}) error {
	sql := "SELECT u.u_id,u.nickname,u.username,u.grade_level,c.* " +
		"FROM " +
		"users AS u JOIN comments AS c ON u.u_id = c.u_id" +
		"WHERE c.b_id = ? AND delete_id IS NULL ORDER BY create_at DESC LIMIT ?,?"

	return  db.SDB.Raw(sql,bid,page*number,number).Scan(&in).Error
}









// 查询用户一级评论
func QueryUserComments(in []Comment,uid string,page,number int64) error {
	return db.SDB.Where("u_id = ? AND delete_id IS NULL ORDER BY create_at DESC LIMIT ?,?",uid,page*number,number).Find(&in).Error
}


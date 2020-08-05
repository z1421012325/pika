package models

import (
	"fmt"
	"pika/server/db"
	"time"
)

// 本子表

/*
CREATE TABLE `benzi` (
`b_id` INT AUTO_INCREMENT PRIMARY KEY,
`title` VARCHAR(100) NOT NULL,
`b_cover` VARCHAR(255) NOT NULL,
`author` VARCHAR(20) DEFAULT NULL,
//`tag` VARCHAR(100) NOT NULL COMMENT '二级tag 以,划分  ',
`upload_at` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '上传时间  ',
`updata_at` DATETIME ON UPDATE CURRENT_TIMESTAMP COMMENT '更行时间 ',
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
	//Tag    string			`gorm:"column:tag" json:"tag,omitempty"`
	UploadAt time.Time		`gorm:"column:upload_at" json:"upload_at,omitempty"`
	UpdateAt time.Time		`gorm:"column:updata_at" json:"updata_at,omitempty"`
	DeleteAt time.Time		`gorm:"column:delete_at" json:"delete_at,omitempty"`
	UploadId int			`gorm:"column:upload_id" json:"upload_id,omitempty"`
	DeleteId int		`gorm:"column:delete_id" json:"delete_id,omitempty"`
}

func (b BenZi) TableName() string {
	return "benzi"
}

// 创建本子仓库
func CreateBenziStorge(uid,title,cover,author string,Types,Tags []string) error {

	benziStorgeSQL := "INSERT INTO benzi (title,b_cover,author,upload_id) VALUES (?,?,?,?)"
	lastInsertIdSQL := "SET @id = LAST_INSERT_ID()"
	TypeAndTagSQL := "INSERT INTO tags (b_id,t_level,tag_name) VALUES "
	//InsertIdSQL := "SELECT @id"

	for _,TypeName := range Types {
		TypeAndTagSQL = TypeAndTagSQL + fmt.Sprintf("(@id,%d,%s),",TYPE_INT,TypeName)
	}

	for _,TagName := range Tags {
		TypeAndTagSQL = TypeAndTagSQL + fmt.Sprintf("(@id,%d,%s),",TAG_INT,TagName)
	}

	//return SqlLink.Exec(TypeAndTagSQL).Exec(InsertIdSQL).Error
	return db.SDB.Exec(benziStorgeSQL,title,cover,author,uid).Exec(lastInsertIdSQL).Exec(TypeAndTagSQL).Error
}

// 删除本子
func DelBenzi(uid string,bid int) error {
	DelBenziSQL := "UPDATE benzi SET delete_at = now(),delete_id = ? WHERE b_id = ? ,upload_id = ?"
	return db.SDB.Exec(DelBenziSQL,uid,bid,uid).Error
}


//  查询本子是否为用户上传
func IsBenziInUserUpload(bid,uid int) bool {
	var benzi BenZi
	db.SDB.Model(benzi).Where("b_id = ?",bid).First(&benzi)
	return benzi.UploadId == uid
}

// 查询本子
func QueryBenzi(bid int,in interface{}) error {

	// todo 在结构体里再次插入结构体并添加tag 能不能再这样查询下映射成功
	/*
		type ClassifyTag struct {
			Tlevel  int  `gorm:"column:t_level"`
			TagName  string `gorm:"column:tag_name"`
		}
		type Benzi struct {
			Bid        int       `gorm:"column:b_id" `
			Title string			`gorm:"column:title" `
			BCover string			`gorm:"column:b_cover" `
			......
		}
		type BenziImg struct {
			......
		}
		type result struct {
			ClassifyTag		`json:"tags"`
			BenziImg		`json:"chapter"`
			Benzi
		}
	 */
	// todo 如果查询异常 多个列查询不了,尝试分多次查询并raw(sql,values...).scan(inteface)
	// 查询用户名称,用户id,本子,作者,标题,上传时间,分类 tag ,章节,历史记录
	// 涉及 tables 有 users ,benzi,bengzi_img,classify_tags
	SQL := "SELECT " +
		"b.b_id,b.title,b.b_cover,b.author,b.updata_at," +
		"(SELECT t_level,tag_name FROM classify_tag WHERE b_id = ? AND delete_at IS NULL) as tags," +
		"(SELECT img_id,chapter_id,chapter_name FROM benzi_img WHERE b_id = ? AND delete_at IS NULL ORDER BY create_at ASC GROUP BY chapter_id) as chapter " +
		"FROM benzi AS b WHERE b.b_id = ? AND b.delete_at IS NULL"

	return db.SDB.Raw(SQL,bid,bid,bid).Scan(&in).Error
	//return db.SDB.Raw(SQL,bid,bid,bid).Scan(in).Error

}










// 查询用户收藏本子    todo  计算该本子总共被收藏了多少
func QueryUserCollectionBenzis(in interface{},uid string,page,number int64) error{
	sql := "SELECT * FROM " +
		"benzi AS b " +
		"JOIN " +
		"users_collections AS uc " +
		"JOIN " +
		"collections AS c " +
		"ON c.collect_id = uc.collect_id " +
		"c.b_id = b.b_id " +
		"WHERE uc.u_id = ? AND c.delete_at IS NULL AND b.delete_at IS NULL " +
		"ORDER BY c.create_at DESC LIMIT ?,?"

	return db.SDB.Raw(sql,uid,page*number,number).Scan(&in).Error
}


func JoinStr(v []string) (r string) {
	for _,i := range v{
		r = r + i
	}
	return
}

package tools

import (
	"os"
	"path"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)



type values struct {
	Get string		`json:"get"`
	Put string		`json:"put"`
	Key string		`json:"key"`
	OSSKey string 	`json:"oss_key"`
}

var (
	Bucket  *oss.Bucket
)


func init()  {

	endpoint := os.Getenv("endpoint")
	accessKeyID := os.Getenv("accessKeyID")
	accessKeySecret := os.Getenv("accessKeySecret")

	// 创建oss连接
	client ,err := oss.New(endpoint,accessKeyID,accessKeySecret)
	if err != nil{
		os.Exit(-1)
	}

	// 得到对oss空间的控制器
	bucketName := os.Getenv("Bucket_Name")
	bucket,err := client.Bucket(bucketName)
	if err != nil{
		os.Exit(-1)
	}
	Bucket = bucket

}

/*
返回预设在oss中的url 和 put上传url 以及 文件名
注意 默认私人链接,无法第三方看到,请到oss的设置中设置反盗链,设置域名或者*(all)
*/
func GetImageToken(files []string) map[int]values {
	// put时 带上header content-type:image/png
	options := []oss.Option{
		oss.ContentType("image/png"),
	}

	tokens := make(map[int]values,len(files))
	// oss仓库名
	stored := os.Getenv("oss_store_touxiang")
	for index,file := range files{

		var value values

		uuid := GetUuid()
		suffix := path.Ext(file)
		key := stored + uuid + suffix

		SignPutURL,err :=Bucket.SignURL(key,oss.HTTPPut,600,options...)
		if err == nil {
			value.Put = SignPutURL
		}else {
			value.Put = ""
		}

		SignGetURL,err := Bucket.SignURL(key,oss.HTTPGet,60*60)
		if err == nil {
			value.Get = SignGetURL
		}else {
			value.Get = ""
		}

		value.Key = file
		value.OSSKey = key

		tokens[index+1] = value
	}
	return tokens
}

func GetVideoToken(filenames string) map[int]values {
	// 格式过滤
	options := []oss.Option{
		oss.ContentType("video/mp4"),
	}

	files := strings.Split(filenames,",")
	tokens := make(map[int]values,len(files))

	stored := os.Getenv("oss_store_kecheng")
	for index,file := range files{
		var value values

		uuid := GetUuid()
		key := stored + uuid

		SignPutURL,err :=Bucket.SignURL(key,oss.HTTPPut,60*60*2,options...)
		if err == nil {
			value.Put = SignPutURL
		}else {
			value.Put = ""
		}


		SignGetURL,err := Bucket.SignURL(key,oss.HTTPGet,60*60*2)
		if err == nil {
			value.Get = SignGetURL
		}else {
			value.Get = ""
		}

		value.Key = file
		value.OSSKey = key

		tokens[index+1] = value
	}
	return tokens
}

func GetImageToOtherToken(filenames string) map[int]values {
	// 格式过滤
	options := []oss.Option{
		oss.ContentType("image/png"),
	}

	files := strings.Split(filenames,",")
	tokens := make(map[int]values,len(files))

	stored := os.Getenv("oss_store_other")
	for index,file := range files{
		var value values

		uuid := GetUuid()
		key := stored + uuid

		SignPutURL,err :=Bucket.SignURL(key,oss.HTTPPut,600,options...)
		if err == nil {
			value.Put = SignPutURL
		}else {
			value.Put = ""
		}


		SignGetURL,err := Bucket.SignURL(key,oss.HTTPGet,60*60)
		if err == nil {
			value.Get = SignGetURL
		}else {
			value.Get = ""
		}

		value.Key = file
		value.OSSKey = key

		tokens[index+1] = value
	}
	return tokens
}


// 补全存储在oss中的路径
func CompletionToOssUrl(ossfile string) (url string)  {
	url,err := Bucket.SignURL(ossfile,oss.HTTPGet,60*60*2)
	if err != nil {
		return ossfile
	}
	return
}




// del
func DeleteFile(filename string) bool{
	err := Bucket.DeleteObject(filename)
	if err != nil {
		return false
	}
	return true
}






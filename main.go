package main

import (
	//_ "pika/config"
	"fmt"
)

// 测试读取配置文件
func main() {
	//fmt.Println(os.Getenv("MYSQL_PASSWORD"))

	maps := make([]map[interface{}]int, 1)

	for i := 0; i < 5; i++ {
		m := make(map[interface{}]int)
		m[i*10] = i
		maps = append(maps, m)
	}
	fmt.Println(maps)

	for index, value := range maps {
		fmt.Println("index", index)
		fmt.Println(value)

		if s, ok := value[40]; ok {
			fmt.Println("------------", s)
		}
	}

}

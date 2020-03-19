package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	//创建csv文件进行数据的写入
	//创建posts.csv文件
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	//创建一个post类型的切片，准备把这些数据写入进去
	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	//创建一个写入器
	writer := csv.NewWriter(csvFile)

	//range来把每一个Post变为字符串切片，其中id要转成字符串
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}

		//用写入器调用Write方法把每个post写进去
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}

	//调用写入器的FLush方法刷新缓冲区
	writer.Flush()

	//csv文件的读取
	//打开一个csv文件
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	//调用NewReader方法创建一个读取器
	reader := csv.NewReader(file)
	//这个值表示每次从文件中读取的字节数，如果不够的话就会中断读取，这里把它设置为-1
	reader.FieldsPerRecord = -1
	//调用ReadAll方法一次读完,这里会返回一组切片组成的切片作为结果,即每条记录的每个字段组成了一个切片，这多组的切片又共同组成一个切片
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	//建立一个Post类型的切片
	var posts []Post

	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		//用一个Post类型的结构体暂时存储一下,用于posts切片的增加
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)

}

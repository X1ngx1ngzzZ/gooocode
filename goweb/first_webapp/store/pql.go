package main

import (
	"database/sql"
	"fmt"
 //因为程序在操作数据库的时候只需要用到database/sql ，而不需要直接使用数据库驱动
 _ "github.com/lib/pq"
)

//这个代码由于编码原因没有运行
//帖子
type Post struct {
	Id　　 int
	Content string
	Author　string
}
//定义指向数据库结构的指针
var Db *sql.DB

//初始化，建立数据库连接，调用OPEN给一个特定的数据库驱动字符串，返回一个sql.DB结构的指针，这里并不会创建连接
func init() {
　　var err error
　　Db, err = sql.Open("postgres", "user=goweb dbname=goweb password=111111 sslmode=disable") 
　 if err != nil {
　　　　panic(err)
　　}
}
//一次获取多篇帖子
func Posts(limit int) (posts []Post, err error) {
	//这条数据库语句返回接口类型
　　rows, err := Db.Query("select id, content, author from posts limit $1",limit)
　　if err != nil {
　　　　return
　　}
//rows接口是个迭代器，重复调用Next方法，什么时候停呢？封装了结构，
for rows.Next() {
	　　　　post := Post{}
	　　　　err = rows.Scan(&post.Id, &post.Content, &post.Author)
	　　　　if err != nil {
	　　　　　　return
	　　　　}
	　　　　posts = append(posts, post)
	　　}
	　　rows.Close()
	　　return
	}
	
	func GetPost(id int) (post Post, err error) { 
	　　post = Post{}
	//sql直接有QueryRow语句
	　　err = Db.QueryRow("select id, content, author from posts where id =$1", id).Scan(&post.Id, &post.Content, &post.Author)
	　　return
	}
	
	func (post *Post) Create() (err error) { 
	//此处为SQL预处理语句模板，用于重复执行指定的SQL语句，这里第一步就直接插入了，这里ID已经自动生成了
	　　statement := "insert into posts (content, author) values ($1, $2)returning id"
	//创建预处理语句，使用sql.DB的prepare方法，创建的接口是stmt的接口　
	stmt, err := Db.Prepare(statement)
	　　if err != nil {
	　　　　return
	　　}
	　　defer stmt.Close()
	//执行重复的SQL语句用stmt，执行一次用sql
	　　err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	　　return
	}
	
	func (post *Post) Update() (err error) {
		//第一个返回值是返回的最终插入id，这里并不关心
	　　_, err = Db.Exec("update posts set content = $2, author = $3 where id =$1", post.Id, post.Content, post.Author)  
	　　return
	}
	
	func (post *Post) Delete() (err error) {
	　　_, err = Db.Exec("delete from posts where id = $1", post.Id) 
	　　return
	}
	
	func main() {
	//创了个新的post
	　　post := Post{Content: "Hello World!", Author: "xxx"}
	　　fmt.Println(post)
	//写入数据库
	　　post.Create()
	   　 fmt.Println(post)  
	//获取上面写回的帖子id的帖子
	　　readPost, _ := GetPost(post.Id)
	  　 fmt.Println(readPost)  
	//修改帖子
	　　readPost.Content = "Bonjour Monde!"
	　　readPost.Author = "Pierre"
	　　readPost.Update()
	//把前10条读出来
	　　posts, _ := Posts()
	   　 fmt.Println(posts)  
	
	//　　readPost.Delete()
	}
package main

import "fmt"

//帖子
type Post struct {
	Id      int
	Content string
	Author  string
}

//建立两个Map来存储帖子，以不同方式
var PostById map[int]*Post
var PostByAuthor map[string]*Post

//存储帖子
func store(post Post) {
	PostById[post.Id] = &post
	PostByAuthor[post.Author] = &post
	//使用Append对Map进行添加
	//PostByAuthor[post.Author]=append(PostByAuthor[Post.Author],&post)
}

func main() {
	PostById = make(map[int]*Post)
	PostByAuthor = make(map[string]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}

	store(post1)
	store(post2)
	store(post3)

	fmt.Println(PostById[1])
	fmt.Println(PostByAuthor["Pedro"])

}

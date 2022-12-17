package core

type Post struct {
	Id      int
	Likes   int
	Owner   *User
	Content string
}

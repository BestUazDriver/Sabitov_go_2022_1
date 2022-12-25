package core

type Post struct {
	Id      int    `bson:"id" json:"id"`
	Likes   int    `bson:"likes" json:"likes"`
	Owner   *User  `bson:"owner" json:"owner"`
	Content string `bson:"content" json:"content"`
}

package database

type User struct {
	UUID        string
	Username    string
	DisplayName string
	PictureURL  string
	CreatedAt   string
}

type Post struct {
	UUID       string
	Caption    string
	ImageURL   string
	CreatedAt  string
	AuthorUUID string
	NLikes     int
	NComments  int
}

type Comment struct {
	UUID      string
	Comment   string
	Timestamp string
	Post      Post
	Author    User
}

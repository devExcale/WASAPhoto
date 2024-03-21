package database

type User struct {
	UUID        string `json:"uuid"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	PictureURL  string `json:"picture_url"`
	NPosts      int    `json:"num_posts"`
	NFollowers  int    `json:"num_followers"`
	NFollowing  int    `json:"num_following"`
	CreatedAt   string `json:"created_at"`
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

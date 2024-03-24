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
	UUID       string `json:"uuid"`
	Caption    string `json:"caption"`
	ImageURL   string `json:"image_url"`
	CreatedAt  string `json:"created_at"`
	AuthorUUID string `json:"author_uuid"`
	NLikes     int    `json:"num_likes"`
	NComments  int    `json:"num_comments"`
}

type Comment struct {
	UUID      string `json:"uuid"`
	Comment   string `json:"comment"`
	Timestamp string `json:"created_at"`
	Post      Post   `json:"post"`
	Author    User   `json:"author"`
}

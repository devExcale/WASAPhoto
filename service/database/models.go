package database

type User struct {
	UUID        string `json:"user_uuid"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	PictureURL  string `json:"picture_url"`
	NPosts      int    `json:"num_posts"`
	NFollowers  int    `json:"num_followers"`
	NFollowing  int    `json:"num_following"`
	CreatedAt   string `json:"created_at"`
}

type Post struct {
	UUID              string `json:"post_uuid"`
	Caption           string `json:"caption"`
	ImageURL          string `json:"image_url"`
	CreatedAt         string `json:"created_at"`
	AuthorUUID        string `json:"author_uuid"`
	AuthorUsername    string `json:"author_username"`
	AuthorDisplayName string `json:"author_display_name"`
	NLikes            int    `json:"num_likes"`
	NComments         int    `json:"num_comments"`
	LoggedUserLiked   bool   `json:"logged_user_liked"`
}

type Comment struct {
	UUID       string `json:"comment_uuid"`
	Comment    string `json:"comment"`
	CreatedAt  string `json:"created_at"`
	PostUUID   string `json:"post_uuid"`
	AuthorUUID string `json:"author_uuid"`
}

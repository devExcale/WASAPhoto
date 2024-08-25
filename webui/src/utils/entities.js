class User {
	constructor(
		uuid = '',
		username = '',
		displayName = '',
		pictureUrl = null,
		memberSince = '',
		numFollowers = 0,
		numFollowing = 0,
		numPosts = 0,
	) {
		this.uuid = uuid
		this.username = username
		this.displayName = displayName
		this.pictureUrl = pictureUrl
		this.memberSince = memberSince
		this.numFollowers = numFollowers
		this.numFollowing = numFollowing
		this.numPosts = numPosts
	}

	static fromResponse(resp) {
		return new User(
			resp.user_uuid,
			resp.username,
			resp.display_name,
			resp.picture_url,
			resp.created_at,
			resp.num_followers,
			resp.num_following,
			resp.num_posts,
		)
	}
}

class Post {
	constructor(
		uuid = '',
		authorUuid = '',
		pictureUrl = null,
		caption = '',
		numLikes = 0,
		numComments = 0,
		timestamp = null,
		loggedUserLiked = false,
	) {
		this.uuid = uuid
		this.authorUuid = authorUuid
		this.pictureUrl = pictureUrl
		this.caption = caption
		this.numLikes = numLikes
		this.numComments = numComments
		this.timestamp = timestamp
		this.loggedUserLiked = loggedUserLiked
	}

	static fromResponse(resp) {
		return new Post(
			resp.post_uuid,
			resp.author_uuid,
			resp.image_url,
			resp.caption,
			resp.num_likes,
			resp.num_comments,
			resp.created_at,
			resp.logged_user_liked,
		)
	}
}

export {User, Post}

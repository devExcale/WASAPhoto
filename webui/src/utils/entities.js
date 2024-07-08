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
		return new User(resp.user_uuid, resp.username, resp.display_name, resp.picture_url, resp.created_at, resp.num_followers, resp.num_following, resp.num_posts,)
	}
}

export {User}

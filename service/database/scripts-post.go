package database

const qSelectPost = `
	SELECT
		post_uuid,
		author_uuid,
		caption,
		image_url,
		num_comments,
		num_likes,
		ts_created
	FROM
		post_full
	WHERE lower(post_uuid) = lower(?)
`

const qSelectPostWithLike = `
	SELECT
		p.post_uuid,
		p.author_uuid,
		p.author_username,
		p.author_display_name,
		p.caption,
		p.image_url,
		p.num_comments,
		p.num_likes,
		p.ts_created,
		EXISTS(
			SELECT 1
			FROM post_like l
			WHERE l.post_uuid = p.post_uuid
				AND lower(l.user_uuid) = lower(?2)
		) AS logged_user_liked
	FROM
		post_full p
	WHERE lower(p.post_uuid) = lower(?1);
`

const qSelectPostsByUser = `
	SELECT
		post_uuid,
		author_uuid,
		author_username,
		author_display_name,
		caption,
		image_url,
		num_comments,
		num_likes,
		ts_created
	FROM
		post_full
	WHERE lower(author_uuid) = lower(?)
	ORDER BY ts_created DESC
`

const qSelectPostsByUserWithLikes = `
	SELECT
		p.post_uuid,
		p.author_uuid,
		p.author_username,
		p.author_display_name,
		p.caption,
		p.image_url,
		p.num_comments,
		p.num_likes,
		p.ts_created,
		EXISTS(
			SELECT 1
			FROM post_like l
			WHERE l.post_uuid = p.post_uuid
				AND lower(l.user_uuid) = lower(?2)
		) AS logged_user_liked
	FROM
		post_full p
	WHERE lower(author_uuid) = lower(?1)
	ORDER BY ts_created DESC
`

const qSelectPostsByFollowed = `
	SELECT
		p.post_uuid,
		p.author_uuid,
		p.author_username,
		p.author_display_name,
		p.caption,
		p.image_url,
		p.num_comments,
		p.num_likes,
		p.ts_created,
		EXISTS(
			SELECT 1
			FROM post_like l
			WHERE l.post_uuid = p.post_uuid
				AND lower(l.user_uuid) = lower(?1)
		) AS logged_user_liked
	FROM
		post_full p
	JOIN
        user_followed f
        ON f.followed_uuid = p.author_uuid
	WHERE lower(f.follower_uuid) = lower(?1)
	ORDER BY p.ts_created DESC
`

const qSelectImage = `
	SELECT image
	FROM post
	WHERE lower(post_uuid) = lower(?)
	LIMIT 1
`

const qUpsertPost = `
	INSERT INTO post(post_uuid, author_uuid, caption, image, ts_created)
	VALUES(lower(?), ?, ?, ?, current_timestamp)
	ON CONFLICT (post_uuid) DO
	UPDATE SET caption = ?
`

const qDeletePost = `
	DELETE FROM post
	WHERE lower(post_uuid) = lower(?)
`

const qDeletePostsByUser = `
	DELETE FROM post
	WHERE lower(author_uuid) = lower(?)
`

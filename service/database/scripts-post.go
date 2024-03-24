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

const qSelectPostsByUser = `
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
	WHERE lower(author_uuid) = lower(?)
	ORDER BY ts_created DESC
`

const qSelectPostsByFollowed = `
	SELECT
		p.post_uuid,
		p.author_uuid,
		p.caption,
		p.image_url,
		p.num_comments,
		p.num_likes,
		p.ts_created
	FROM
		post_full p
	JOIN
        user_followed f
        ON f.followed_uuid = p.author_uuid
	WHERE lower(f.follower_uuid) = lower(?)
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

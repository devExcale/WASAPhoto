package database

const qSelectCommentByUUID = `
	SELECT
		c.comment_uuid,
		c.comment,
		c.post_uuid,
		c.author_uuid,
		c.ts_created,
		u.username,
		u.display_name
	FROM
		post_comment c
	JOIN
		user u
		ON c.author_uuid = u.user_uuid
	WHERE lower(comment_uuid) = lower(?)
	ORDER BY c.ts_created DESC
`

const qSelectCommentsByPost = `
	SELECT
		c.comment_uuid,
		c.comment,
		c.post_uuid,
		c.author_uuid,
		c.ts_created,
		u.username,
		u.display_name
	FROM
		post_comment c
	JOIN
		user u
		ON c.author_uuid = u.user_uuid
	WHERE lower(post_uuid) = lower(?)
	ORDER BY c.ts_created DESC
`

const qUpsertComment = `
	INSERT INTO post_comment(comment_uuid, comment, post_uuid, author_uuid, ts_created)
	VALUES(lower(?), ?, ?, ?, current_timestamp)
	ON CONFLICT (comment_uuid) DO
	UPDATE SET comment = ?
`

const qDeleteComment = `
	DELETE FROM post_comment
	WHERE lower(comment_uuid) = lower(?)
`

const qDeletePostComments = `
	DELETE FROM post_comment
	WHERE lower(comment_uuid) IN (
		SELECT lower(comment_uuid)
		FROM post_comment
		WHERE lower(post_uuid) = lower(?)
	)
`

const qDeleteCommentsByUser = `
	DELETE FROM post_comment
	WHERE lower(author_uuid) = lower(?)
`

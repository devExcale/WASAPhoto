package database

const qSelectCommentsByPost = `
	SELECT
		comment_uuid,
		comment,
		post_uuid,
		author_uuid,
		ts_created
	FROM
		post_comment
	WHERE lower(post_uuid) = lower(?)
	ORDER BY ts_created DESC
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

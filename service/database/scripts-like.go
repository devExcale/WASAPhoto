package database

const qSelectUserHasLikedPost = `
	SELECT
		COUNT(*)
	FROM (
		SELECT 1
		FROM post_like
		WHERE lower(post_uuid) = lower(?)
		  AND lower(user_uuid) = lower(?)
	)`

const qSelectUsersLikedPost = `
	SELECT
		u.user_uuid,
		u.username,
		u.display_name,
		u.picture_url,
		u.ts_created
	FROM
		user u
	JOIN
		post_like l
		ON l.user_uuid = u.user_uuid
	WHERE lower(l.post_uuid) = lower(?)
	ORDER BY ts_created DESC
`

const qInsertLike = `
	INSERT INTO post_like(post_uuid, user_uuid)
	VALUES(lower(?), lower(?))
`

const qDeleteLike = `
	DELETE FROM post_like
	WHERE lower(post_uuid) = lower(?)
	  AND lower(user_uuid) = lower(?)
`

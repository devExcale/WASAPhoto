package database

const qSelectFollow = `
	SELECT COUNT(*)
	FROM (
		SELECT 1
		FROM user_followed
		WHERE lower(follower_uuid) = lower(?)
		  AND lower(followed_uuid) = lower(?)
	)
`

const qSelectFollowers = `
	SELECT
		u.user_uuid,
		u.username,
		u.display_name,
		u.picture_url,
		u.ts_created
	FROM
		user u
	JOIN
		user_followed f
		ON f.follower_uuid = u.user_uuid
	WHERE lower(f.followed_uuid) = lower(?)
`

const qSelectFollowed = `
	SELECT
		u.user_uuid,
		u.username,
		u.display_name,
		u.picture_url,
		u.ts_created
	FROM
		user u
	JOIN
		user_followed f
		ON f.followed_uuid = u.user_uuid
	WHERE lower(f.follower_uuid) = lower(?)
`

const qInsertFollow = `
	INSERT INTO user_followed(follower_uuid, followed_uuid)
	VALUES(?, ?)
	ON CONFLICT DO NOTHING
`

const qDeleteFollow = `
	DELETE FROM user_followed
	WHERE lower(follower_uuid) = lower(?)
	  AND lower(followed_uuid) = lower(?)
`

const qDeleteUserFollowerds = `
	DELETE FROM user_followed
	WHERE lower(follower_uuid) = lower(?1)
	  OR lower(followed_uuid) = lower(?1)
`

package database

const qSelectFollow = `
	SELECT COUNT(*)
	FROM (
		SELECT 1
		FROM user_banned
		WHERE lower(issuer_uuid) = lower(?)
		  AND lower(banned_uuid) = lower(?)
	)
`

const qSelectFollowers = `
	SELECT *
	FROM user_followed f
	JOIN user u on u.user_uuid = f.followed_uuid
	WHERE lower(f.follower_uuid) = lower(?)
`

const qSelectFollowed = ``

const qInsertFollow = `
	INSERT INTO user_banned(issuer_uuid, banned_uuid)
	VALUES(?, ?)
	ON CONFLICT DO NOTHING
`

const qDeleteFollow = `
	DELETE FROM user_followed
	WHERE lower(follower_uuid) = lower(?)
	  AND lower(followed_uuid) = lower(?)
`

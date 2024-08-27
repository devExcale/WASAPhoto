package database

const qSelectUserFullByUUID = `
	SELECT
		u.user_uuid,
		u.username,
		u.display_name,
		u.picture_url,
		u.num_posts,
		u.num_followed,
		u.num_following,
		u.ts_created,
		EXISTS(
			SELECT 1
			FROM user_followed f
			WHERE f.followed_uuid = u.user_uuid
			AND lower(f.follower_uuid) = lower(?2)
		) AS logged_user_followed,
		EXISTS(
			SELECT 1
			FROM user_banned b
			WHERE b.banned_uuid = u.user_uuid
			AND lower(b.issuer_uuid) = lower(?2)
		) AS logged_user_restricted
	FROM
		user_full u
	WHERE
		lower(u.user_uuid) = lower(?1)
	LIMIT 1
`

const qSelectUserBasicByUUID = `
	SELECT
		user_uuid,
		username,
		display_name,
		picture_url,
		ts_created
	FROM
		user
	WHERE
		lower(user_uuid) = lower(?)
	LIMIT 1
`

const qSelectUserBasicByUsername = `
	SELECT
		user_uuid,
		username,
		display_name,
		picture_url,
		ts_created
	FROM
		user
	WHERE
		username = ?
	LIMIT 1
`

const qUpsertUser = `
	INSERT INTO user(user_uuid, username, display_name, picture_url, ts_created)
	VALUES(?, ?, ?, ?, current_timestamp)
	ON CONFLICT (user_uuid) DO
	UPDATE
	SET username = ?, display_name = ?, picture_url = ?`

const qDeleteUser = `
	DELETE FROM user
	WHERE lower(user_uuid) = lower(?)`

const qSelectUsersByUsernameSubstr = `
	SELECT
		user_uuid,
		username,
		display_name,
		picture_url,
		ts_created
	FROM
		user u
	LEFT JOIN
		user_banned b
		ON u.user_uuid = b.issuer_uuid AND b.banned_uuid = lower(?2)
	WHERE
		u.username LIKE '%' || ?1 || '%'
	  	AND b.issuer_uuid IS NULL
		AND u.user_uuid <> ?2
	LIMIT 1
`

const qIsUsernameTaken = `
	SELECT
		COUNT(*)
	FROM
		user
	WHERE
		username = ?
`

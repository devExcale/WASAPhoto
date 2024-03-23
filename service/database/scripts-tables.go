package database

const qCreateTableUser = `
	CREATE TABLE user
	(
		user_uuid    TEXT(36) PRIMARY KEY,
		username     TEXT(20)  NOT NULL UNIQUE,
		display_name TEXT(40)  NOT NULL,
		picture_url  TEXT      NULL,
		ts_created   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	);
`

const qCreateTablePost = `
	CREATE TABLE post
	(
		post_uuid   TEXT(36) PRIMARY KEY,
		author_uuid TEXT(36) NOT NULL,
		caption     TEXT,
		image       BLOB     NOT NULL,
		timestamp   TEXT     NOT NULL DEFAULT current_timestamp,
		FOREIGN KEY (author_uuid) REFERENCES user (user_uuid)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	);
`

const qCreateTableComment = `
	CREATE TABLE post_comment
	(
		comment_uuid TEXT(36) PRIMARY KEY,
		post_uuid    TEXT(36) NOT NULL,
		author_uuid  TEXT(36) NOT NULL,
		comment      TEXT     NOT NULL,
		timestamp    TEXT     NOT NULL DEFAULT current_timestamp,
		FOREIGN KEY (post_uuid) REFERENCES post (post_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
		FOREIGN KEY (author_uuid) REFERENCES user (user_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE
	);
`

const qCreateTableLike = `
	CREATE TABLE post_like
	(
		post_uuid TEXT(36) NOT NULL,
		user_uuid TEXT(36) NOT NULL,
		PRIMARY KEY (post_uuid, user_uuid),
		FOREIGN KEY (post_uuid) REFERENCES post (post_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
		FOREIGN KEY (user_uuid) REFERENCES user (user_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE
	);
`

const qCreateTableFollowedUsers = `
	CREATE TABLE user_followed
	(
		follower_uuid TEXT(36) NOT NULL,
		followed_uuid TEXT(36) NOT NULL,
		PRIMARY KEY (follower_uuid, followed_uuid),
		FOREIGN KEY (follower_uuid) REFERENCES user (user_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
		FOREIGN KEY (followed_uuid) REFERENCES user (user_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE
	);
`

const qCreateTableBannedUsers = `
	CREATE TABLE user_banned
	(
		issuer_uuid TEXT(36) NOT NULL,
		banned_uuid TEXT(36) NOT NULL,
		PRIMARY KEY (issuer_uuid, banned_uuid),
		FOREIGN KEY (issuer_uuid) REFERENCES user (user_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
		FOREIGN KEY (banned_uuid) REFERENCES user (user_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE
	);
`

const qCreateViewUserFull = `
	CREATE VIEW user_full AS
		SELECT
			u.user_uuid,
			u.username,
			u.display_name,
			'TODO' as picture_url,
			COUNT(DISTINCT p.post_uuid) as num_posts,
			COUNT(DISTINCT fd.follower_uuid) as num_followed,
			COUNT(DISTINCT fg.followed_uuid) as num_following,
			u.ts_created
		FROM
			user u
		LEFT JOIN
			post p ON p.author_uuid = u.user_uuid
		LEFT JOIN
			user_followed fd ON fd.followed_uuid = u.user_uuid
		LEFT JOIN
			user_followed fg ON fg.follower_uuid = u.user_uuid
		GROUP BY
			u.user_uuid;
`

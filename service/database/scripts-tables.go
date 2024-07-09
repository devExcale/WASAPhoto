package database

var tables = map[string]string{
	"user":          qCreateTableUser,
	"post":          qCreateTablePost,
	"post_comment":  qCreateTableComment,
	"post_like":     qCreateTableLike,
	"user_followed": qCreateTableFollowedUsers,
	"user_banned":   qCreateTableBannedUsers,
	"user_full":     qCreateViewUserFull,
	"post_full":     qCreateViewPostFull,
}

const qCreateTableUser = `
	CREATE TABLE user
	(
		user_uuid    TEXT(36) PRIMARY KEY,
		username     TEXT(20)  NOT NULL UNIQUE,
		display_name TEXT(40)  NOT NULL,
		picture_url  TEXT      NULL,
		ts_created   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)
`

const qCreateTablePost = `
	CREATE TABLE post
	(
		post_uuid   TEXT(36) PRIMARY KEY,
		author_uuid TEXT(36) NOT NULL,
		caption     TEXT,
		image       BLOB     NOT NULL,
		ts_created   TEXT     NOT NULL DEFAULT current_timestamp,
		FOREIGN KEY (author_uuid) REFERENCES user (user_uuid)
			ON DELETE CASCADE
			ON UPDATE CASCADE
	)
`

const qCreateTableComment = `
	CREATE TABLE post_comment
	(
		comment_uuid TEXT(36) PRIMARY KEY,
		post_uuid    TEXT(36) NOT NULL,
		author_uuid  TEXT(36) NOT NULL,
		comment      TEXT     NOT NULL,
		ts_created    TEXT     NOT NULL DEFAULT current_timestamp,
		FOREIGN KEY (post_uuid) REFERENCES post (post_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE,
		FOREIGN KEY (author_uuid) REFERENCES user (user_uuid)
			ON UPDATE CASCADE
			ON DELETE CASCADE
	)
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
	)
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
	)
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
	)
`

const qCreateViewUserFull = `
	CREATE VIEW user_full AS
		SELECT
			u.user_uuid,
			u.username,
			u.display_name,
			u. picture_url,
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
			u.user_uuid
`

const qCreateViewPostFull = `
	CREATE VIEW post_full AS
		SELECT
			p.post_uuid,
			p.author_uuid,
			p.caption,
			'/users/' || p.author_uuid || '/feed/' || p.post_uuid || '/webp' as image_url,
			COUNT(DISTINCT l.user_uuid) as num_likes,
			COUNT(DISTINCT c.comment_uuid) as num_comments,
			p.ts_created
		FROM
			post p
		LEFT JOIN
			post_like l ON l.post_uuid = p.post_uuid
		LEFT JOIN
			post_comment c ON c.post_uuid = p.post_uuid
		GROUP BY
			p.post_uuid
`

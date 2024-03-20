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

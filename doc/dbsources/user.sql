CREATE TABLE user
(
	user_uuid    TEXT(36) PRIMARY KEY,
	username     TEXT(20)  NOT NULL UNIQUE,
	display_name TEXT(40)  NOT NULL,
	picture_url  TEXT      NULL,
	ts_created   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

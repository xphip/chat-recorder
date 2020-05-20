--
-- SQLITE
--

CREATE TABLE IF NOT EXISTS sys (
	tx_sys INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.sys

	tx_token TEXT

	-- id_autostart INTEGER NOT NULL CHECK (id_autostart IN (0, 1)) -- Inactive, Active
);

CREATE TABLE IF NOT EXISTS message_discord (
	cd_message INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.cd_message
	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

	in_status INTEGER NOT NULL CHECK (in_status IN (0, 1, 2)) DEFAULT 0, -- 0:Created, 1:Edited, 2:Deleted
	cd_message_id INTEGER NOT NULL,

	cd_guild INTEGER NOT NULL,
	tx_guild TEXT NOT NULL,

	cd_channel INTEGER NOT NULL,
	tx_channel TEXT NOT NULL,

	cd_author INTEGER,
	tx_author TEXT,

	tx_message TEXT

);

-- -- // TODO: add ignored users table
-- CREATE TABLE IF NOT EXISTS users_ignored (
-- 	cd_user INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT UNIQUE, -- PK.cd_user
-- 	dt_created INTEGER NOT NULL DEFAULT CURRENT_TIMESTAMP,

-- );
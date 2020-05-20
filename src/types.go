package main

import (
	"database/sql"
)

// DB conn structure
type DB struct {
	Conn 		*sql.DB
}

// SQL constants
const (
	SQL_INSERT = "INSERT INTO %s (%s) VALUES (%s)"

	SQL_CHECK = `SELECT name FROM sqlite_master WHERE type='table' AND name='%s';`

	SQL_CreateMessageDiscord_FIELDS = "in_status, cd_message_id, cd_guild, tx_guild, cd_channel, tx_channel, cd_author, tx_author, tx_message"
	SQL_UpdateMessageDiscord_FIELDS = "in_status, cd_message_id, cd_guild, tx_guild, cd_channel, tx_channel, cd_author, tx_author, tx_message"
	SQL_DeleteMessageDiscord_FIELDS = "in_status, cd_message_id, cd_guild, tx_guild, cd_channel, tx_channel"
)

const (
	SQL_TYPE_CREATED = iota
	SQL_TYPE_UPDATE
	SQL_TYPE_DELETED
)

// DBTable 				`json:"message_discord"`
type MessageDiscord struct {
	ID 			int 	`json:"cd_message"`
	CreatedAt 	int 	`json:"dt_created"`

	Status 		int 	`json:"in_status"`
	MessageID 	string 	`json:"cd_message_id"`

	GuildID 	string 	`json:"cd_guild"`
	GuildName 	string 	`json:"tx_guild"`

	ChannelID 	string 	`json:"cd_channel"`
	ChannelName string 	`json:"tx_channel"`

	AuthorID 	string 	`json:"cd_author"`
	AuthorName 	string 	`json:"tx_author"`

	Message 	string 	`json:"tx_message"`
}

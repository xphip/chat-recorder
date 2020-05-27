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

	SQL_CreateMessageDiscord_Fields = "in_status, cd_message_id, cd_guild, tx_guild, cd_channel, tx_channel, cd_author, tx_author, tx_message"
	SQL_UpdateMessageDiscord_Fields = "in_status, cd_message_id, cd_guild, tx_guild, cd_channel, tx_channel, cd_author, tx_author, tx_message"
	SQL_DeleteMessageDiscord_Fields = "in_status, cd_message_id, cd_guild, tx_guild, cd_channel, tx_channel"

	SQL_CreateMessageDiscordLog_Fields 		= "in_status, tx_raw"
	SQL_CreateMessageDiscordErrorLog_Fields = "in_status, tx_description, tx_raw"
)

const (
	Table_Sys 						= "sys"
	Table_MessageDiscord 			= "message_discord"
	Table_MessageDiscordLog 		= "message_discord_log"
	Table_MessageDiscordErrorLog 	= "message_discord_error_log"
)

const (
	SQL_TYPE_NULL = iota -1
	SQL_TYPE_CREATED
	SQL_TYPE_UPDATE
	SQL_TYPE_DELETED
)

const (
	MsgError_Marshal = "Marshal error"
	MsgError_Unknown = "Unknown error"
	MsgError_Ready = "Ready info error"
	MsgError_Channel = "Channel info error"
	MsgError_Guild = "Guild info error"
	MsgError_Author = "Author info error"
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

type MessageDiscordLog struct {
	ID 			int 	`json:"cd_message"`
	CreatedAt 	int 	`json:"dt_created"`
	Status 		int 	`json:"in_status"`
	Raw 		string 	`json:"tx_raw"`
}

type MessageDiscordErrorLog struct {
	ID 			int 	`json:"cd_message"`
	CreatedAt 	int 	`json:"dt_created"`
	Status 		int 	`json:"in_status"`
	Description string 	`json:"tx_description"`
	Raw 		string 	`json:"tx_raw"`
}

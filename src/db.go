package main

import (
	"fmt"
	"log"
	"os"

	"database/sql"
	_ "github.com/xphip-pack/go-sqlite3"
)

func (db *DB) Open(fileName string) (err error) {

	runScript := false

	if _, er := os.Stat(fileName); os.IsNotExist(er) {
		runScript = true
	}

	db.Conn, err = sql.Open("sqlite3", fileName)

	if runScript {
		log.Println("[LOG] Create tables..")
		_, err = db.Conn.Exec(RawSQL)
	}

	return
}

func (db *DB) Close() (err error) {
	err = db.Conn.Close()
	return 
}

func (db *DB) Check() (err error) {

	rows, err := db.Conn.Query(fmt.Sprintf(SQL_CHECK, Table_MessageDiscord))
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		result := ""

		err = rows.Scan(&result)
		if err != nil {
			return
		}
		return
	}

	return
}

func AddMessageDiscordErrorLog(msg MessageDiscordErrorLog) (err error) {

	table := fmt.Sprintf(`'%d', '%s', '%s'`,
						msg.Status,
						msg.Description,
						msg.Raw,
	)

	sql := fmt.Sprintf(SQL_INSERT, Table_MessageDiscordErrorLog, SQL_CreateMessageDiscordErrorLog_Fields, table)
	_, err = db.Conn.Exec(sql)

	return
}

func AddMessageDiscordLog(msg MessageDiscordLog) (err error) {

	table := fmt.Sprintf(`'%d', '%s'`,
						msg.Status,
						msg.Raw,
	)

	sql := fmt.Sprintf(SQL_INSERT, Table_MessageDiscordLog, SQL_CreateMessageDiscordLog_Fields, table)
	_, err = db.Conn.Exec(sql)

	return
}

func (db *DB) AddMessageDiscord(msg MessageDiscord) (err error) {

	table := fmt.Sprintf(`'%d', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'`,
						SQL_TYPE_CREATED,
						msg.MessageID,
						msg.GuildID,
						msg.GuildName,
						msg.ChannelID,
						msg.ChannelName,
						msg.AuthorID,
						msg.AuthorName,
						msg.Message,
	)

	sql := fmt.Sprintf(SQL_INSERT, Table_MessageDiscord, SQL_CreateMessageDiscord_Fields, table)
	_, err = db.Conn.Exec(sql)

	return
}

func (db *DB) UpdateMessageDiscord(msg MessageDiscord) (err error) {

	table := fmt.Sprintf(`'%d', '%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s'`,
						SQL_TYPE_UPDATE,
						msg.MessageID,
						msg.GuildID,
						msg.GuildName,
						msg.ChannelID,
						msg.ChannelName,
						msg.AuthorID,
						msg.AuthorName,
						msg.Message,
	)

	sql := fmt.Sprintf(SQL_INSERT, Table_MessageDiscord, SQL_UpdateMessageDiscord_Fields, table)
	_, err = db.Conn.Exec(sql)

	return
}

func (db *DB) DeleteMessageDiscord(msg MessageDiscord) (err error) {

	table := fmt.Sprintf(`'%d', '%s', '%s', '%s', '%s', '%s'`,
						SQL_TYPE_DELETED,
						msg.MessageID,
						msg.GuildID,
						msg.GuildName,
						msg.ChannelID,
						msg.ChannelName,
	)

	sql := fmt.Sprintf(SQL_INSERT, Table_MessageDiscord, SQL_DeleteMessageDiscord_Fields, table)
	_, err = db.Conn.Exec(sql)

	return
}
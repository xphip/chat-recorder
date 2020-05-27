package main

import (
	"flag"
	"log"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"encoding/json"

	"github.com/xphip-pack/discordgo"
)

func init() {
	flag.StringVar(&token, "t", "", "Bot Token")
	flag.Parse()
}

var (
	token string
	buffer = make([][]byte, 0)

	DbFileName = "./log.db"
	db DB
)

func main() {

	db.Open(DbFileName)

	err := db.Check()
	if err != nil {
		os.Remove(DbFileName)
		db.Close()
		log.Fatal(err)
		return
	}
	defer db.Close()

	if token == "" {
		log.Println("No token provided.")
		return
	}

	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Println("Error creating Discord session: ", err)
		return
	}

	dg.AddHandler(ready)
	dg.AddHandler(messageCreate)
	dg.AddHandler(messageUpdate)
	dg.AddHandler(messageDelete)
	dg.AddHandler(guildJoin)

	err = dg.Open()
	if err != nil {
		log.Println("Error opening Discord session: ", err)
	}

	log.Println("App started. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func ready(s *discordgo.Session, e *discordgo.Ready) {
	j, err := json.Marshal(e)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_NULL,
			Description: MsgError_Marshal,
			Raw: fmt.Sprintf("%s", j),
		})
	}

	defer AddMessageDiscordErrorLog(MessageDiscordErrorLog{
		Status: SQL_TYPE_NULL,
		Description: MsgError_Ready,
		Raw: fmt.Sprintf("%s", j),
	})

	s.UpdateStatus(0, ":)")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	j, err := json.Marshal(m)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_CREATED,
			Description: MsgError_Marshal,
			Raw: fmt.Sprintf("%s", j),
		})
	}

	defer func() {
		if err := recover(); err != nil {
			AddMessageDiscordErrorLog(MessageDiscordErrorLog{
				Status: SQL_TYPE_CREATED,
				Description: MsgError_Unknown,
				Raw: fmt.Sprintf("%s", j),
			})
		}
	}()

	AddMessageDiscordLog(MessageDiscordLog{
		Status: SQL_TYPE_CREATED,
		Raw: fmt.Sprintf("%s", j),
	})

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_CREATED,
			Description: MsgError_Channel,
			Raw: fmt.Sprintf("%s", j),
		})
		return
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_CREATED,
			Description: MsgError_Guild,
			Raw: fmt.Sprintf("%s", j),
		})
		return
	}

	if m.Author == nil || m.Author.ID == "" || m.Author.Username == "" {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_UPDATE,
			Description: MsgError_Author,
			Raw: fmt.Sprintf("%s", j),
		})
		return
	}

	err = db.AddMessageDiscord(MessageDiscord{
		MessageID: m.ID,

		GuildID: m.GuildID,
		GuildName: g.Name,

		ChannelID: m.ChannelID,
		ChannelName: c.Name,

		AuthorID: m.Author.ID,
		AuthorName: m.Author.Username,

		Message: m.Content,
	})

}

func messageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	j, err := json.Marshal(m)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_UPDATE,
			Description: MsgError_Marshal,
			Raw: fmt.Sprintf("%s", j),
		})
	}

	defer func() {
		if err := recover(); err != nil {
			AddMessageDiscordErrorLog(MessageDiscordErrorLog{
				Status: SQL_TYPE_UPDATE,
				Description: MsgError_Unknown,
				Raw: fmt.Sprintf("%s", j),
			})
		}
	}()

	AddMessageDiscordLog(MessageDiscordLog{
		Status: SQL_TYPE_UPDATE,
		Raw: fmt.Sprintf("%s", j),
	})

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_UPDATE,
			Description: MsgError_Channel,
			Raw: fmt.Sprintf("%s", j),
		})
		return
	}

	g, err := s.State.Guild(m.GuildID)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_UPDATE,
			Description: MsgError_Guild,
			Raw: fmt.Sprintf("%s", j),
		})
		return
	}

	if m.Author == nil || m.Author.ID == "" || m.Author.Username == "" {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_UPDATE,
			Description: MsgError_Author,
			Raw: fmt.Sprintf("%s", j),
		})
		return
	}

	err = db.UpdateMessageDiscord(MessageDiscord{
		MessageID: m.ID,

		GuildID: m.GuildID,
		GuildName: g.Name,

		ChannelID: m.ChannelID,
		ChannelName: c.Name,

		AuthorID: m.Author.ID,
		AuthorName: m.Author.Username,

		Message: m.Content,
	})

}

func messageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {
	j, err := json.Marshal(m)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_DELETED,
			Description: MsgError_Marshal,
			Raw: fmt.Sprintf("%s", j),
		})
	}

	defer func() {
		if err := recover(); err != nil {
			AddMessageDiscordErrorLog(MessageDiscordErrorLog{
				Status: SQL_TYPE_DELETED,
				Description: MsgError_Unknown,
				Raw: fmt.Sprintf("%s", j),
			})
		}
	}()

	AddMessageDiscordLog(MessageDiscordLog{
		Status: SQL_TYPE_DELETED,
		Raw: fmt.Sprintf("%s", j),
	})

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_DELETED,
			Description: MsgError_Channel,
			Raw: fmt.Sprintf("%s", j),
		})
		return
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		AddMessageDiscordErrorLog(MessageDiscordErrorLog{
			Status: SQL_TYPE_DELETED,
			Description: MsgError_Guild,
			Raw: fmt.Sprintf("%s", j),
		})
		return
	}

	err = db.DeleteMessageDiscord(MessageDiscord{
		MessageID: m.ID,

		GuildID: m.GuildID,
		GuildName: g.Name,

		ChannelID: m.ChannelID,
		ChannelName: c.Name,
	})

}

func guildJoin(s *discordgo.Session, event *discordgo.GuildCreate) {

	if event.Guild.Unavailable {
		return
	}

	for _, channel := range event.Guild.Channels {
		if channel.ID == event.Guild.ID {
			// _, _ = s.ChannelMessageSend(channel.ID, "Chat-recorder is up!")
			return
		}
	}
}

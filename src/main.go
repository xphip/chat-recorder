package main

import (
	"flag"
	"log"
	"fmt"
	"os"
	"os/signal"
	"syscall"

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

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, ":)")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		AddMessageDiscordError(MessageDiscordErrorLog{
			Status: SQL_TYPE_CREATED,
			Description: MsgError_Channel,
			Raw: fmt.Sprintf("%#v", m),
		})
		return
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		AddMessageDiscordError(MessageDiscordErrorLog{
			Status: SQL_TYPE_CREATED,
			Description: MsgError_Guild,
			Raw: fmt.Sprintf("%#v", m),
		})
		return
	}

	err = db.AddMessageDiscord(MessageDiscord{
		MessageID: m.ID,

		GuildID: g.ID,
		GuildName: g.Name,

		ChannelID: c.ID,
		ChannelName: c.Name,

		AuthorID: m.Author.ID,
		AuthorName: m.Author.Username,

		Message: m.Content,
	})

}

func messageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		AddMessageDiscordError(MessageDiscordErrorLog{
			Status: SQL_TYPE_CREATED,
			Description: MsgError_Channel,
			Raw: fmt.Sprintf("%#v", m),
		})
		return
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		AddMessageDiscordError(MessageDiscordErrorLog{
			Status: SQL_TYPE_CREATED,
			Description: MsgError_Guild,
			Raw: fmt.Sprintf("%#v", m),
		})
		return
	}

	err = db.UpdateMessageDiscord(MessageDiscord{
		MessageID: m.ID,

		GuildID: g.ID,
		GuildName: g.Name,

		ChannelID: c.ID,
		ChannelName: c.Name,

		AuthorID: m.Author.ID,
		AuthorName: m.Author.Username,

		Message: m.Content,
	})

}

func messageDelete(s *discordgo.Session, m *discordgo.MessageDelete) {

	c, err := s.State.Channel(m.ChannelID)
	if err != nil {
		AddMessageDiscordError(MessageDiscordErrorLog{
			Status: SQL_TYPE_CREATED,
			Description: MsgError_Channel,
			Raw: fmt.Sprintf("%#v", m),
		})
		return
	}

	g, err := s.State.Guild(c.GuildID)
	if err != nil {
		AddMessageDiscordError(MessageDiscordErrorLog{
			Status: SQL_TYPE_CREATED,
			Description: MsgError_Guild,
			Raw: fmt.Sprintf("%#v", m),
		})
		return
	}

	err = db.DeleteMessageDiscord(MessageDiscord{
		MessageID: m.ID,

		GuildID: g.ID,
		GuildName: g.Name,

		ChannelID: c.ID,
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

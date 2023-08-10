package main

import (
	"discord-bot-audit-log/internal/pkg/handlers"
	"flag"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var (
	token        *string
	logChannelId *string
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func init() {
	token = flag.String("token", "", "discord bot token")
	logChannelId = flag.String("log-channel-id", "", "channel id for logs")
	flag.Parse()

	ErrorLogger = log.New(os.Stderr, "Error: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	session, err := discordgo.New("Bot " + *token)
	if err != nil {
		ErrorLogger.Print(err)
		return
	}
	session.Identify.Intents = discordgo.IntentGuildVoiceStates | discordgo.IntentGuildMessages

	handlers.LogChannelID = *logChannelId
	session.AddHandler(handlers.VoiceStateUpdate)

	err = session.Open()
	if err != nil {
		ErrorLogger.Print(err)
		return
	}
	defer session.Close()
	InfoLogger.Print("bot is running")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	InfoLogger.Print("bot stopped")
}

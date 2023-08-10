package handlers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	joinColor = 0x08ff08
	leftColor = 0xff0800
	moveColor = 0xff8000
)

var LogChannelID string
var WarningLogger *log.Logger

func init() {
	WarningLogger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
}

func VoiceStateUpdate(session *discordgo.Session, event *discordgo.VoiceStateUpdate) {
	avatar := event.Member.User.AvatarURL("")
	embed := &discordgo.MessageEmbed{
		Author: &discordgo.MessageEmbedAuthor{
			Name:    event.Member.User.Username,
			IconURL: avatar,
		},
		Timestamp: time.Now().Format(time.RFC3339),
	}

	if oldVoiceState := event.BeforeUpdate; oldVoiceState == nil {
		channel, err := session.Channel(event.ChannelID)
		if err != nil {
			WarningLogger.Println(err)
			return
		}

		channelField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("зашел в канал **%s**", channel.Name),
		}
		embed.Fields = append(embed.Fields, channelField)
		embed.Color = joinColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if event.ChannelID == "" {
		channel, err := session.Channel(oldVoiceState.ChannelID)
		if err != nil {
			WarningLogger.Println(err)
			return
		}

		channelField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("вышел из канала **%s**", channel.Name),
		}
		embed.Fields = append(embed.Fields, channelField)
		embed.Color = leftColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if oldVoiceState.ChannelID != event.ChannelID {
		oldChannel, err := session.Channel(oldVoiceState.ChannelID)
		if err != nil {
			WarningLogger.Println(err)
			return
		}

		newChannel, err := session.Channel(event.ChannelID)
		if err != nil {
			WarningLogger.Println(err)
			return
		}

		channelField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("перешёл из канала **%s** в канал **%s**", oldChannel.Name, newChannel.Name),
		}
		embed.Fields = append(embed.Fields, channelField)
		embed.Color = moveColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	}
}

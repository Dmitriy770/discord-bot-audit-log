package handlers

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/bwmarrin/discordgo"
)

const (
	joinColor       = 0x00b300
	leftColor       = 0xff0800
	moveColor       = 0xff9900
	selfMuteColor   = 0x8a0041
	selfUnmuteColor = 0xea69a6
	selfDeafColor   = 0x8a0041
	selfUndeafColor = 0xea69a6
	muteColor       = 0x42036f
	unmuteColor     = 0xa767d5
	deafColor       = 0x42036f
	undeafColor     = 0xa767d5
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
			Name:    event.Member.User.Username + " : ",
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
		embed.Author.Name += "JOIN"

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
		embed.Author.Name += "LEFT"

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
		embed.Author.Name += "MOVE"

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if !oldVoiceState.Deaf && event.Deaf {
		infoField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("пользователю отключили звук"),
		}
		embed.Fields = append(embed.Fields, infoField)
		embed.Author.Name += "DEAF"
		embed.Color = deafColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if oldVoiceState.Deaf && !event.Deaf {
		infoField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("пользователю включили звук"),
		}
		embed.Fields = append(embed.Fields, infoField)
		embed.Author.Name += "UNDEAF"
		embed.Color = undeafColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if !oldVoiceState.Mute && event.Mute {
		infoField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("пользователю отключили микрофон"),
		}
		embed.Fields = append(embed.Fields, infoField)
		embed.Author.Name += "MUTE"
		embed.Color = muteColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if oldVoiceState.Mute && !event.Mute {
		infoField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("пользователю включили микрофон"),
		}
		embed.Fields = append(embed.Fields, infoField)
		embed.Author.Name += "UNMUTE"
		embed.Color = unmuteColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if !oldVoiceState.SelfDeaf && event.SelfDeaf {
		infoField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("отключил звук"),
		}
		embed.Fields = append(embed.Fields, infoField)
		embed.Author.Name += "SELF DEAF"
		embed.Color = selfDeafColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if oldVoiceState.SelfDeaf && !event.SelfDeaf {
		infoField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("включил звук"),
		}
		embed.Fields = append(embed.Fields, infoField)
		embed.Author.Name += "SELF UNDEAF"
		embed.Color = selfUndeafColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if !oldVoiceState.SelfMute && event.SelfMute {
		infoField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("отключил микрофон"),
		}
		embed.Fields = append(embed.Fields, infoField)
		embed.Author.Name += "SELF MUTE"
		embed.Color = selfMuteColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	} else if oldVoiceState.SelfMute && !event.SelfMute {
		infoField := &discordgo.MessageEmbedField{
			Value: fmt.Sprintf("включил микрофон"),
		}
		embed.Fields = append(embed.Fields, infoField)
		embed.Author.Name += "SELF UNMUTE"
		embed.Color = selfUnmuteColor

		session.ChannelMessageSendEmbed(LogChannelID, embed)
	}
}

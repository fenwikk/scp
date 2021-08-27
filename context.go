package scp

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

type Ctx struct {
	Options     []*discordgo.ApplicationCommandInteractionDataOption
	Session     *discordgo.Session
	Interaction *discordgo.InteractionCreate
	Router      *Router
	Guild       *discordgo.Guild
	Channel     *discordgo.Channel
	User        *discordgo.User
}

func (ctx *Ctx) WaitForResponse() *discordgo.MessageCreate {
	var mc *discordgo.MessageCreate
	interactionDone := false

	removeHandler := ctx.Session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		channel, _ := s.Channel(m.ChannelID)

		correctMessage := m.Author.ID == ctx.User.ID && channel.ID == ctx.Channel.ID
		log.Println(correctMessage)
		if correctMessage {
			interactionDone = true
			mc = m
		}
	})

	for !interactionDone {
	}
	removeHandler()

	return mc
}

// func (ctx *Ctx) WaitForEmojiResponse() *discordgo.MessageCreate {
// 	var ic *discordgo.InteractionCreate
// 	interactionDone := false

// 	removeHandler := ctx.Session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
// 		channel, _ := s.Channel(i.ChannelID)
// 		if i.Interaction.Type == discordgo.intera && channel == ctx.Channel {
// 			interactionDone = true
// 			ic = i
// 		}
// 	})

// 	for !interactionDone {
// 	}
// 	removeHandler()

// 	return ic
// }

func (ctx *Ctx) Respond(itype discordgo.InteractionResponseType, data *discordgo.InteractionResponseData) {
	ctx.Session.InteractionRespond(ctx.Interaction.Interaction, &discordgo.InteractionResponse{
		Type: itype,
		Data: data,
	})
}

func (ctx *Ctx) EditResponse(newrsp *discordgo.WebhookEdit) {
	ctx.Session.InteractionResponseEdit(ctx.Session.State.User.ID, ctx.Interaction.Interaction, newrsp)
}

func (ctx *Ctx) DeleteResponse() {
	ctx.Session.InteractionResponseDelete(ctx.Session.State.User.ID, ctx.Interaction.Interaction)
}

func (ctx *Ctx) RespondFollowup(data *discordgo.WebhookParams, wait bool) *Followup {
	m := ctx.Session.FollowupMessageCreate(ctx.Session.State.User.ID, ctx.Interaction.Interaction, wait, data)

	return &Followup{
		Message: m,
		ctx:     ctx,
	}
}

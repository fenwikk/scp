package scp

import (
	"github.com/bwmarrin/discordgo"
)

type Followup struct {
	Message *discordgo.Message
	ctx     *Ctx
}

func (m *Followup) EditFollowup(newrsp *discordgo.WebhookEdit) {
	m.ctx.Session.FollowupMessageEdit(m.ctx.Session.State.User.ID, m.ctx.Interaction.Interaction, m.Message.ID, newrsp)
}

func (m *Followup) DeleteFollowup(newrsp *discordgo.WebhookEdit) {
	m.ctx.Session.FollowupMessageDelete(m.ctx.Session.State.User.ID, m.ctx.Interaction.Interaction, newrsp)
}

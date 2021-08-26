package scp

import (
	"github.com/bwmarrin/discordgo"
)

type Command struct {
	ID            string
	ApplicationID string
	Type          discordgo.ApplicationCommandType
	Name          string
	Description   string
	Version       string
	Options       []*ApplicationCommandOption
	Handler       func(*Ctx)
}

func (c *Command) ToApplicationCommand() *discordgo.ApplicationCommand {
	ac := &discordgo.ApplicationCommand{
		ID:            c.ID,
		ApplicationID: c.ApplicationID,
		Type:          c.Type,
		Name:          c.Name,
		Description:   c.Description,
		Version:       c.Version,
	}

	for _, o := range c.Options {
		ac.Options = append(ac.Options, o.toDGAC())
	}

	return ac
}

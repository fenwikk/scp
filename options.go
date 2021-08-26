package scp

import "github.com/bwmarrin/discordgo"

type ApplicationCommandOption struct {
	Type        discordgo.ApplicationCommandOptionType `json:"type"`
	Name        string                                 `json:"name"`
	Description string                                 `json:"description,omitempty"`
	// NOTE: This feature was on the API, but at some point developers decided to remove it.
	// So I commented it, until it will be officially on the docs.
	// Default     bool                              `json:"default"`
	Required bool                                        `json:"required"`
	Choices  []*discordgo.ApplicationCommandOptionChoice `json:"choices"`
	Options  []*ApplicationCommandOption                 `json:"options"`
	Handler  func(*Ctx)
}

func (ac *ApplicationCommandOption) toDGAC() *discordgo.ApplicationCommandOption {
	dgac := &discordgo.ApplicationCommandOption{
		Type:        ac.Type,
		Name:        ac.Name,
		Description: ac.Description,
		Required:    ac.Required,
		Choices:     ac.Choices,
	}

	for _, o := range ac.Options {
		dgac.Options = append(dgac.Options, o.toDGAC())
	}

	return dgac
}

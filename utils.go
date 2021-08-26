package scp

import (
	"regexp"
)

var (
	RegexArguments      = regexp.MustCompile("(\"[^\"]+\"|[^\\s]+)")
	RegexUserMention    = regexp.MustCompile("<@!?(\\d+)>")
	RegexRoleMention    = regexp.MustCompile("<@&(\\d+)>")
	RegexChannelMention = regexp.MustCompile("<#(\\d+)>")
	RegexSnowflakeID    = regexp.MustCompile("[\\d]{18}")
)

type IDType int

const (
	UserMentionID IDType = iota
	RoleMentionID
	ChannelMentionID
)

func (ctx *Ctx) StrToID(str string, itype IDType) string {
	switch itype {
	case UserMentionID:
		if RegexUserMention.MatchString(str) {
			return RegexUserMention.FindStringSubmatch(str)[1]
		} else {
			return ""
		}
	case RoleMentionID:
		if RegexRoleMention.MatchString(str) {
			return RegexRoleMention.FindStringSubmatch(str)[1]
		} else {
			return ""
		}
	case ChannelMentionID:
		if RegexChannelMention.MatchString(str) {
			return RegexChannelMention.FindStringSubmatch(str)[1]
		} else {
			return ""
		}
	}

	if RegexSnowflakeID.MatchString(str) {
		return RegexSnowflakeID.FindString(str)
	} else {
		return ""
	}
}

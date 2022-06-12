package wotoStyle

import (
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/tg"
)

type Appender func(s string) styling.StyledTextOption

type wotoStyling struct {
	_value []styling.StyledTextOption
}

type WStyle interface {
	Append(md WStyle) WStyle
	AppendThis(md WStyle) WStyle
	AppendNormal(v string) WStyle
	AppendNormalThis(v string) WStyle
	AppendBold(v string) WStyle
	AppendBoldThis(v string) WStyle
	AppendItalic(v string) WStyle
	AppendItalicThis(v string) WStyle
	AppendMono(v string) WStyle
	AppendMonoThis(v string) WStyle
	AppendSpoiler(v string) WStyle
	AppendSpoilerThis(v string) WStyle
	AppendHyperLink(text, url string) WStyle
	AppendHyperLinkThis(text, url string) WStyle
	AppendMention(text string, id int64) WStyle
	AppendMentionThis(text string, id int64) WStyle
	AppendUserMention(text string, u *tg.User) WStyle
	AppendUserMentionThis(text string, u *tg.User) WStyle
	AppendMentionHash(text string, u, hash int64) WStyle
	AppendMentionHashThis(text string, u, hash int64) WStyle
	AppendMentionUsername(text string, username string) WStyle
	AppendMentionUsernameThis(text string, username string) WStyle

	Normal(v string) WStyle
	Bold(v string) WStyle
	Italic(v string) WStyle
	Mono(v string) WStyle
	Spoiler(v string) WStyle
	Link(text, url string) WStyle
	Mention(text string, id int64) WStyle
	MentionHash(text string, u, hash int64) WStyle
	UserMention(text string, u *tg.User) WStyle
	MentionUsername(text string, username string) WStyle

	El() WStyle
	ElThis() WStyle
	Space() WStyle
	SpaceThis() WStyle
	Tab() WStyle
	TabThis() WStyle
	GetStylingArray() []styling.StyledTextOption
	setValue(v []styling.StyledTextOption)
}

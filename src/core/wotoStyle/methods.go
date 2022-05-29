package wotoStyle

import (
	"github.com/AnimeKaizoku/RestorerRobot/src/core/utils/tgUtils"
	"github.com/gotd/td/telegram/message/styling"
	"github.com/gotd/td/tg"
)

/*
Append(md WStyle) WStyle
	AppendThis(md WStyle) WStyle
	ToString() string
	AppendNormal(v string) WStyle
	AppendNormalThis(v string) WStyle
	AppendBold(v string) WStyle
	AppendBoldThis(v string) WStyle
	AppendItalic(v string) WStyle
	AppendItalicThis(v string) WStyle
	AppendMono(v string) WStyle
	AppendMonoThis(v string) WStyle
	AppendHyperLink(text, url string) WStyle
	AppendHyperLinkThis(text, url string) WStyle
	AppendMention(text string, id int64) WStyle
	AppendMentionThis(text string, id int64) WStyle
	El() WStyle
	ElThis() WStyle
	Space() WStyle
	SpaceThis() WStyle
	Tab() WStyle
	TabThis() WStyle
	GetStylingArray() []styling.StyledTextOption
	getValue() string
	setValue(v string)
*/

func (s *wotoStyling) Append(w WStyle) WStyle {
	newStyle := &wotoStyling{
		_value: s._value,
	}
	return newStyle.appendValue(w.GetStylingArray()...)
}

func (s *wotoStyling) AppendThis(w WStyle) WStyle {
	return s.appendValue(w.GetStylingArray()...)
}

func (s *wotoStyling) AppendByFunc(v string, f Appender) WStyle {
	newStyle := &wotoStyling{
		_value: s._value,
	}
	return newStyle.appendValue(f(v))
}

func (s *wotoStyling) AppendByFuncThis(v string, f Appender) WStyle {
	return s.appendValue(f(v))
}

func (s *wotoStyling) AppendNormal(v string) WStyle {
	return s.AppendByFunc(v, styling.Plain)
}

func (s *wotoStyling) AppendNormalThis(v string) WStyle {
	return s.AppendByFuncThis(v, styling.Plain)
}

func (s *wotoStyling) AppendBold(v string) WStyle {
	return s.AppendByFunc(v, styling.Bold)
}

func (s *wotoStyling) AppendBoldThis(v string) WStyle {
	return s.AppendByFuncThis(v, styling.Bold)
}

func (s *wotoStyling) AppendSpoiler(v string) WStyle {
	return s.AppendByFuncThis(v, styling.Spoiler)
}

func (s *wotoStyling) AppendSpoilerThis(v string) WStyle {
	return s.AppendByFuncThis(v, styling.Spoiler)
}

func (s *wotoStyling) AppendItalic(v string) WStyle {
	return s.AppendByFunc(v, styling.Italic)
}

func (s *wotoStyling) AppendItalicThis(v string) WStyle {
	return s.AppendByFuncThis(v, styling.Italic)
}

func (s *wotoStyling) AppendMono(v string) WStyle {
	return s.AppendByFunc(v, styling.Code)
}

func (s *wotoStyling) AppendMonoThis(v string) WStyle {
	return s.AppendByFuncThis(v, styling.Code)
}

func (s *wotoStyling) AppendHyperLink(text, url string) WStyle {
	return s.AppendByFunc(text, func(s string) styling.StyledTextOption {
		return styling.TextURL(s, url)
	})
}

func (s *wotoStyling) AppendHyperLinkThis(text, url string) WStyle {
	return s.AppendByFuncThis(text, func(s string) styling.StyledTextOption {
		return styling.TextURL(s, url)
	})
}

func (s *wotoStyling) AppendMention(text string, id int64) WStyle {
	return s.AppendByFunc(text, func(s string) styling.StyledTextOption {
		input := tgUtils.GetInputUserFromId(id)
		if input == nil {
			return styling.Code(s)
		}

		return styling.MentionName(s, input)
	})
}

func (s *wotoStyling) AppendMentionThis(text string, id int64) WStyle {
	return s.AppendByFuncThis(text, func(s string) styling.StyledTextOption {
		input := tgUtils.GetInputUserFromId(id)
		if input == nil {
			return styling.Code(s)
		}

		return styling.MentionName(s, input)
	})
}

func (s *wotoStyling) AppendUserMention(text string, u *tg.User) WStyle {
	if u == nil {
		return s.AppendMono(text)
	}
	return s.AppendMentionHash(text, u.ID, u.AccessHash)
}

func (s *wotoStyling) AppendUserMentionThis(text string, u *tg.User) WStyle {
	if u == nil {
		return s.AppendMonoThis(text)
	}
	return s.AppendMentionHashThis(text, u.ID, u.AccessHash)
}

func (s *wotoStyling) AppendMentionHash(text string, u, hash int64) WStyle {
	return s.AppendByFunc(text, func(s string) styling.StyledTextOption {
		return styling.MentionName(s, tgUtils.GetInputUser(u, hash))
	})
}

func (s *wotoStyling) AppendMentionHashThis(text string, u, hash int64) WStyle {
	return s.AppendByFuncThis(text, func(s string) styling.StyledTextOption {
		return styling.MentionName(s, tgUtils.GetInputUser(u, hash))
	})
}

func (s *wotoStyling) AppendMentionUsername(text string, username string) WStyle {
	return s.AppendByFunc(text, func(s string) styling.StyledTextOption {
		u := tgUtils.GetInputUserFromUsername(username)
		if u == nil {
			return styling.Code(s)
		}
		return styling.MentionName(s, u)
	})
}

func (s *wotoStyling) AppendMentionUsernameThis(text string, username string) WStyle {
	return s.AppendByFuncThis(text, func(s string) styling.StyledTextOption {
		u := tgUtils.GetInputUserFromUsername(username)
		if u == nil {
			return styling.Code(s)
		}
		return styling.MentionName(s, u)
	})
}

func (s *wotoStyling) El() WStyle {
	return s.AppendNormal("\n")
}

func (s *wotoStyling) ElThis() WStyle {
	return s.AppendNormalThis("\n")
}

func (s *wotoStyling) Space() WStyle {
	return s.AppendNormal(" ")
}

func (s *wotoStyling) SpaceThis() WStyle {
	return s.AppendNormalThis(" ")
}

func (s *wotoStyling) Tab() WStyle {
	return s.AppendNormal("\t")
}

func (s *wotoStyling) TabThis() WStyle {
	return s.AppendNormalThis("\t")
}

func (s *wotoStyling) GetStylingArray() []styling.StyledTextOption {
	return s._value
}

func (s *wotoStyling) setValue(v []styling.StyledTextOption) {
	s._value = v
}

func (s *wotoStyling) appendValue(v ...styling.StyledTextOption) WStyle {
	s._value = append(s._value, v...)
	return s
}

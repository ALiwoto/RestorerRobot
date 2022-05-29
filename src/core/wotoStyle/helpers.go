package wotoStyle

import "github.com/gotd/td/telegram/message/styling"

func GetNormal(value string) WStyle {
	s := &wotoStyling{}
	return s.AppendNormalThis(value)
}

func GetBold(value string) WStyle {
	s := &wotoStyling{}
	return s.AppendBoldThis(value)
}

func GetSpoiler(value string) WStyle {
	s := &wotoStyling{}
	return s.AppendSpoilerThis(value)
}

func GetItalic(value string) WStyle {
	s := &wotoStyling{}
	return s.AppendItalicThis(value)
}

func GetMono(value string) WStyle {
	s := &wotoStyling{}
	return s.AppendMonoThis(value)
}

func GetHyperLink(text, url string) WStyle {
	s := &wotoStyling{}
	return s.AppendHyperLinkThis(text, url)
}

func ArrayToStyle(v []styling.StyledTextOption) WStyle {
	return &wotoStyling{
		_value: v,
	}
}

func ParamToStyle(v ...styling.StyledTextOption) WStyle {
	return &wotoStyling{
		_value: v,
	}
}

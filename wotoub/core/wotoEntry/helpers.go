package wotoEntry

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ALiwoto/wotoub/wotoub/core/wotoEntry/enteryManager"
	"github.com/ALiwoto/wotoub/wotoub/core/wotoValues"
	"github.com/ALiwoto/wotoub/wotoub/database/sessionDatabase"
	"github.com/gotd/td/tg"
)

func LoadAllHandlers(d *tg.UpdateDispatcher) {
	d.OnNewMessage(NewMessageHandler)
	d.OnNewChannelMessage(NewChannelMessageHandler)
	d.OnNewScheduledMessage(NewScheduledMessageHandler)
	d.OnEditMessage(EditMessageHandler)
}

func MainUpdateEntry(ctx context.Context, u tg.UpdatesClass) error {
	switch v := u.(type) {
	//case *tg.UpdatesTooLong: // updatesTooLong#e317af7e
	//case *tg.UpdateShortMessage: // updateShortMessage#313bc7f8
	//case *tg.UpdateShortChatMessage: // updateShortChatMessage#4d6deea5
	//case *tg.UpdateShort: // updateShort#78d4dec1
	//case *tg.UpdatesCombined: // updatesCombined#725b04c3
	case *tg.Updates: // updates#74ae4240
		if len(v.Users) > 0 {
			cacheUsers(v.Users)
			return nil
		}
		//if len(v.Updates) != 0 {
		//	for _, update := range v.Updates {
		//		handleNewUpdate(update)
		//	}
		//}
		return nil
		//log.Println("this is an update:", v)
	//case *tg.UpdateShortSentMessage: // updateShortSentMessage#9015e101
	default:
		//s := fmt.Sprint(v)
		//if strings.Contains(s, "woto") {
		//	log.Println(s)
		//}
	}

	return nil
}

//  switch v := g.(type) {
//  case *tg.UpdateNewMessage: // updateNewMessage#1f2b0afd
//  case *tg.UpdateMessageID: // updateMessageID#4e90bfd6
//  case *tg.UpdateDeleteMessages: // updateDeleteMessages#a20db0e5
//  case *tg.UpdateUserTyping: // updateUserTyping#c01e857f
//  case *tg.UpdateChatUserTyping: // updateChatUserTyping#83487af0
//  case *tg.UpdateChatParticipants: // updateChatParticipants#7761198
//  case *tg.UpdateUserStatus: // updateUserStatus#e5bdf8de
//  case *tg.UpdateUserName: // updateUserName#c3f202e0
//  case *tg.UpdateUserPhoto: // updateUserPhoto#f227868c
//  case *tg.UpdateNewEncryptedMessage: // updateNewEncryptedMessage#12bcbd9a
//  case *tg.UpdateEncryptedChatTyping: // updateEncryptedChatTyping#1710f156
//  case *tg.UpdateEncryption: // updateEncryption#b4a2e88d
//  case *tg.UpdateEncryptedMessagesRead: // updateEncryptedMessagesRead#38fe25b7
//  case *tg.UpdateChatParticipantAdd: // updateChatParticipantAdd#3dda5451
//  case *tg.UpdateChatParticipantDelete: // updateChatParticipantDelete#e32f3d77
//  case *tg.UpdateDCOptions: // updateDcOptions#8e5e9873
//  case *tg.UpdateNotifySettings: // updateNotifySettings#bec268ef
//  case *tg.UpdateServiceNotification: // updateServiceNotification#ebe46819
//  case *tg.UpdatePrivacy: // updatePrivacy#ee3b272a
//  case *tg.UpdateUserPhone: // updateUserPhone#5492a13
//  case *tg.UpdateReadHistoryInbox: // updateReadHistoryInbox#9c974fdf
//  case *tg.UpdateReadHistoryOutbox: // updateReadHistoryOutbox#2f2f21bf
//  case *tg.UpdateWebPage: // updateWebPage#7f891213
//  case *tg.UpdateReadMessagesContents: // updateReadMessagesContents#68c13933
//  case *tg.UpdateChannelTooLong: // updateChannelTooLong#108d941f
//  case *tg.UpdateChannel: // updateChannel#635b4c09
//  case *tg.UpdateNewChannelMessage: // updateNewChannelMessage#62ba04d9
//  case *tg.UpdateReadChannelInbox: // updateReadChannelInbox#922e6e10
//  case *tg.UpdateDeleteChannelMessages: // updateDeleteChannelMessages#c32d5b12
//  case *tg.UpdateChannelMessageViews: // updateChannelMessageViews#f226ac08
//  case *tg.UpdateChatParticipantAdmin: // updateChatParticipantAdmin#d7ca61a2
//  case *tg.UpdateNewStickerSet: // updateNewStickerSet#688a30aa
//  case *tg.UpdateStickerSetsOrder: // updateStickerSetsOrder#bb2d201
//  case *tg.UpdateStickerSets: // updateStickerSets#43ae3dec
//  case *tg.UpdateSavedGifs: // updateSavedGifs#9375341e
//  case *tg.UpdateBotInlineQuery: // updateBotInlineQuery#496f379c
//  case *tg.UpdateBotInlineSend: // updateBotInlineSend#12f12a07
//  case *tg.UpdateEditChannelMessage: // updateEditChannelMessage#1b3f4df7
//  case *tg.UpdateBotCallbackQuery: // updateBotCallbackQuery#b9cfc48d
//  case *tg.UpdateEditMessage: // updateEditMessage#e40370a3
//  case *tg.UpdateInlineBotCallbackQuery: // updateInlineBotCallbackQuery#691e9052
//  case *tg.UpdateReadChannelOutbox: // updateReadChannelOutbox#b75f99a9
//  case *tg.UpdateDraftMessage: // updateDraftMessage#ee2bb969
//  case *tg.UpdateReadFeaturedStickers: // updateReadFeaturedStickers#571d2742
//  case *tg.UpdateRecentStickers: // updateRecentStickers#9a422c20
//  case *tg.UpdateConfig: // updateConfig#a229dd06
//  case *tg.UpdatePtsChanged: // updatePtsChanged#3354678f
//  case *tg.UpdateChannelWebPage: // updateChannelWebPage#2f2ba99f
//  case *tg.UpdateDialogPinned: // updateDialogPinned#6e6fe51c
//  case *tg.UpdatePinnedDialogs: // updatePinnedDialogs#fa0f3ca2
//  case *tg.UpdateBotWebhookJSON: // updateBotWebhookJSON#8317c0c3
//  case *tg.UpdateBotWebhookJSONQuery: // updateBotWebhookJSONQuery#9b9240a6
//  case *tg.UpdateBotShippingQuery: // updateBotShippingQuery#b5aefd7d
//  case *tg.UpdateBotPrecheckoutQuery: // updateBotPrecheckoutQuery#8caa9a96
//  case *tg.UpdatePhoneCall: // updatePhoneCall#ab0f6b1e
//  case *tg.UpdateLangPackTooLong: // updateLangPackTooLong#46560264
//  case *tg.UpdateLangPack: // updateLangPack#56022f4d
//  case *tg.UpdateFavedStickers: // updateFavedStickers#e511996d
//  case *tg.UpdateChannelReadMessagesContents: // updateChannelReadMessagesContents#44bdd535
//  case *tg.UpdateContactsReset: // updateContactsReset#7084a7be
//  case *tg.UpdateChannelAvailableMessages: // updateChannelAvailableMessages#b23fc698
//  case *tg.UpdateDialogUnreadMark: // updateDialogUnreadMark#e16459c3
//  case *tg.UpdateMessagePoll: // updateMessagePoll#aca1657b
//  case *tg.UpdateChatDefaultBannedRights: // updateChatDefaultBannedRights#54c01850
//  case *tg.UpdateFolderPeers: // updateFolderPeers#19360dc0
//  case *tg.UpdatePeerSettings: // updatePeerSettings#6a7e7366
//  case *tg.UpdatePeerLocated: // updatePeerLocated#b4afcfb0
//  case *tg.UpdateNewScheduledMessage: // updateNewScheduledMessage#39a51dfb
//  case *tg.UpdateDeleteScheduledMessages: // updateDeleteScheduledMessages#90866cee
//  case *tg.UpdateTheme: // updateTheme#8216fba3
//  case *tg.UpdateGeoLiveViewed: // updateGeoLiveViewed#871fb939
//  case *tg.UpdateLoginToken: // updateLoginToken#564fe691
//  case *tg.UpdateMessagePollVote: // updateMessagePollVote#106395c9
//  case *tg.UpdateDialogFilter: // updateDialogFilter#26ffde7d
//  case *tg.UpdateDialogFilterOrder: // updateDialogFilterOrder#a5d72105
//  case *tg.UpdateDialogFilters: // updateDialogFilters#3504914f
//  case *tg.UpdatePhoneCallSignalingData: // updatePhoneCallSignalingData#2661bf09
//  case *tg.UpdateChannelMessageForwards: // updateChannelMessageForwards#d29a27f4
//  case *tg.UpdateReadChannelDiscussionInbox: // updateReadChannelDiscussionInbox#d6b19546
//  case *tg.UpdateReadChannelDiscussionOutbox: // updateReadChannelDiscussionOutbox#695c9e7c
//  case *tg.UpdatePeerBlocked: // updatePeerBlocked#246a4b22
//  case *tg.UpdateChannelUserTyping: // updateChannelUserTyping#8c88c923
//  case *tg.UpdatePinnedMessages: // updatePinnedMessages#ed85eab5
//  case *tg.UpdatePinnedChannelMessages: // updatePinnedChannelMessages#5bb98608
//  case *tg.UpdateChat: // updateChat#f89a6a4e
//  case *tg.UpdateGroupCallParticipants: // updateGroupCallParticipants#f2ebdb4e
//  case *tg.UpdateGroupCall: // updateGroupCall#14b24500
//  case *tg.UpdatePeerHistoryTTL: // updatePeerHistoryTTL#bb9bb9a5
//  case *tg.UpdateChatParticipant: // updateChatParticipant#d087663a
//  case *tg.UpdateChannelParticipant: // updateChannelParticipant#985d3abb
//  case *tg.UpdateBotStopped: // updateBotStopped#c4870a49
//  case *tg.UpdateGroupCallConnection: // updateGroupCallConnection#b783982
//  case *tg.UpdateBotCommands: // updateBotCommands#4d712f2e
//  default: panic(v)
//  }
func AhandleNewUpdate(gUpdate tg.UpdateClass) {
	switch u := gUpdate.(type) {
	case *tg.UpdateNewMessage: // updateNewMessage#1f2b0afd
		if u.Message != nil {
			handleNewMessageOld(u.Message)
		}
		return
	//  case *tg.UpdateMessageID: // updateMessageID#4e90bfd6
	case *tg.UpdateDeleteMessages: // updateDeleteMessages#a20db0e5
		log.Println("delete messages:", u)
	//  case *tg.UpdateUserTyping: // updateUserTyping#c01e857f
	//  case *tg.UpdateChatUserTyping: // updateChatUserTyping#83487af0
	//  case *tg.UpdateChatParticipants: // updateChatParticipants#7761198
	//  case *tg.UpdateUserStatus: // updateUserStatus#e5bdf8de
	//  case *tg.UpdateUserName: // updateUserName#c3f202e0
	//  case *tg.UpdateUserPhoto: // updateUserPhoto#f227868c
	//  case *tg.UpdateNewEncryptedMessage: // updateNewEncryptedMessage#12bcbd9a
	//  case *tg.UpdateEncryptedChatTyping: // updateEncryptedChatTyping#1710f156
	//  case *tg.UpdateEncryption: // updateEncryption#b4a2e88d
	//  case *tg.UpdateEncryptedMessagesRead: // updateEncryptedMessagesRead#38fe25b7
	//  case *tg.UpdateChatParticipantAdd: // updateChatParticipantAdd#3dda5451
	//  case *tg.UpdateChatParticipantDelete: // updateChatParticipantDelete#e32f3d77
	//  case *tg.UpdateDCOptions: // updateDcOptions#8e5e9873
	//  case *tg.UpdateNotifySettings: // updateNotifySettings#bec268ef
	//  case *tg.UpdateServiceNotification: // updateServiceNotification#ebe46819
	//  case *tg.UpdatePrivacy: // updatePrivacy#ee3b272a
	//  case *tg.UpdateUserPhone: // updateUserPhone#5492a13
	//  case *tg.UpdateReadHistoryInbox: // updateReadHistoryInbox#9c974fdf
	//  case *tg.UpdateReadHistoryOutbox: // updateReadHistoryOutbox#2f2f21bf
	//  case *tg.UpdateWebPage: // updateWebPage#7f891213
	//  case *tg.UpdateReadMessagesContents: // updateReadMessagesContents#68c13933
	//  case *tg.UpdateChannelTooLong: // updateChannelTooLong#108d941f
	//  case *tg.UpdateChannel: // updateChannel#635b4c09
	case *tg.UpdateNewChannelMessage: // updateNewChannelMessage#62ba04d9
		if u.Message != nil {
			handleNewMessageOld(u.Message)
		}
		return
	//  case *tg.UpdateReadChannelInbox: // updateReadChannelInbox#922e6e10
	//  case *tg.UpdateDeleteChannelMessages: // updateDeleteChannelMessages#c32d5b12
	//  case *tg.UpdateChannelMessageViews: // updateChannelMessageViews#f226ac08
	//  case *tg.UpdateChatParticipantAdmin: // updateChatParticipantAdmin#d7ca61a2
	//  case *tg.UpdateNewStickerSet: // updateNewStickerSet#688a30aa
	//  case *tg.UpdateStickerSetsOrder: // updateStickerSetsOrder#bb2d201
	//  case *tg.UpdateStickerSets: // updateStickerSets#43ae3dec
	//  case *tg.UpdateSavedGifs: // updateSavedGifs#9375341e
	//  case *tg.UpdateBotInlineQuery: // updateBotInlineQuery#496f379c
	//  case *tg.UpdateBotInlineSend: // updateBotInlineSend#12f12a07
	//  case *tg.UpdateEditChannelMessage: // updateEditChannelMessage#1b3f4df7
	//  case *tg.UpdateBotCallbackQuery: // updateBotCallbackQuery#b9cfc48d
	//  case *tg.UpdateEditMessage: // updateEditMessage#e40370a3
	//  case *tg.UpdateInlineBotCallbackQuery: // updateInlineBotCallbackQuery#691e9052
	//  case *tg.UpdateReadChannelOutbox: // updateReadChannelOutbox#b75f99a9
	//  case *tg.UpdateDraftMessage: // updateDraftMessage#ee2bb969
	//  case *tg.UpdateReadFeaturedStickers: // updateReadFeaturedStickers#571d2742
	//  case *tg.UpdateRecentStickers: // updateRecentStickers#9a422c20
	//  case *tg.UpdateConfig: // updateConfig#a229dd06
	//  case *tg.UpdatePtsChanged: // updatePtsChanged#3354678f
	//  case *tg.UpdateChannelWebPage: // updateChannelWebPage#2f2ba99f
	//  case *tg.UpdateDialogPinned: // updateDialogPinned#6e6fe51c
	//  case *tg.UpdatePinnedDialogs: // updatePinnedDialogs#fa0f3ca2
	//  case *tg.UpdateBotWebhookJSON: // updateBotWebhookJSON#8317c0c3
	//  case *tg.UpdateBotWebhookJSONQuery: // updateBotWebhookJSONQuery#9b9240a6
	//  case *tg.UpdateBotShippingQuery: // updateBotShippingQuery#b5aefd7d
	//  case *tg.UpdateBotPrecheckoutQuery: // updateBotPrecheckoutQuery#8caa9a96
	//  case *tg.UpdatePhoneCall: // updatePhoneCall#ab0f6b1e
	//  case *tg.UpdateLangPackTooLong: // updateLangPackTooLong#46560264
	//  case *tg.UpdateLangPack: // updateLangPack#56022f4d
	//  case *tg.UpdateFavedStickers: // updateFavedStickers#e511996d
	//  case *tg.UpdateChannelReadMessagesContents: // updateChannelReadMessagesContents#44bdd535
	//  case *tg.UpdateContactsReset: // updateContactsReset#7084a7be
	//  case *tg.UpdateChannelAvailableMessages: // updateChannelAvailableMessages#b23fc698
	//  case *tg.UpdateDialogUnreadMark: // updateDialogUnreadMark#e16459c3
	//  case *tg.UpdateMessagePoll: // updateMessagePoll#aca1657b
	//  case *tg.UpdateChatDefaultBannedRights: // updateChatDefaultBannedRights#54c01850
	//  case *tg.UpdateFolderPeers: // updateFolderPeers#19360dc0
	//  case *tg.UpdatePeerSettings: // updatePeerSettings#6a7e7366
	//  case *tg.UpdatePeerLocated: // updatePeerLocated#b4afcfb0
	//  case *tg.UpdateNewScheduledMessage: // updateNewScheduledMessage#39a51dfb
	//  case *tg.UpdateDeleteScheduledMessages: // updateDeleteScheduledMessages#90866cee
	//  case *tg.UpdateTheme: // updateTheme#8216fba3
	//  case *tg.UpdateGeoLiveViewed: // updateGeoLiveViewed#871fb939
	//  case *tg.UpdateLoginToken: // updateLoginToken#564fe691
	//  case *tg.UpdateMessagePollVote: // updateMessagePollVote#106395c9
	//  case *tg.UpdateDialogFilter: // updateDialogFilter#26ffde7d
	//  case *tg.UpdateDialogFilterOrder: // updateDialogFilterOrder#a5d72105
	//  case *tg.UpdateDialogFilters: // updateDialogFilters#3504914f
	//  case *tg.UpdatePhoneCallSignalingData: // updatePhoneCallSignalingData#2661bf09
	//  case *tg.UpdateChannelMessageForwards: // updateChannelMessageForwards#d29a27f4
	//  case *tg.UpdateReadChannelDiscussionInbox: // updateReadChannelDiscussionInbox#d6b19546
	//  case *tg.UpdateReadChannelDiscussionOutbox: // updateReadChannelDiscussionOutbox#695c9e7c
	//  case *tg.UpdatePeerBlocked: // updatePeerBlocked#246a4b22
	//  case *tg.UpdateChannelUserTyping: // updateChannelUserTyping#8c88c923
	//  case *tg.UpdatePinnedMessages: // updatePinnedMessages#ed85eab5
	//  case *tg.UpdatePinnedChannelMessages: // updatePinnedChannelMessages#5bb98608
	//  case *tg.UpdateChat: // updateChat#f89a6a4e
	//  case *tg.UpdateGroupCallParticipants: // updateGroupCallParticipants#f2ebdb4e
	//  case *tg.UpdateGroupCall: // updateGroupCall#14b24500
	//  case *tg.UpdatePeerHistoryTTL: // updatePeerHistoryTTL#bb9bb9a5
	//  case *tg.UpdateChatParticipant: // updateChatParticipant#d087663a
	//  case *tg.UpdateChannelParticipant: // updateChannelParticipant#985d3abb
	//  case *tg.UpdateBotStopped: // updateBotStopped#c4870a49
	//  case *tg.UpdateGroupCallConnection: // updateGroupCallConnection#b783982
	//  case *tg.UpdateBotCommands: // updateBotCommands#4d712f2e
	default:
		s := fmt.Sprint(u)
		if strings.Contains(s, "woto") {
			log.Println(s)
		}
	}
}

func cacheUsers(users []tg.UserClass) {
	for _, u := range users {
		if u != nil {
			cacheUser(u)
		}
	}
}

func cacheUser(u tg.UserClass) {
	user, ok := u.(*tg.User)
	if ok && user.ID != 0 && user.AccessHash != 0 {
		_ = sessionDatabase.SaveTgUser(user)
	}
}

//  switch v := g.(type) {
//  case *tg.MessageEmpty: // messageEmpty#90a6ca84
//  case *tg.Message: // message#85d6cbe2
//  case *tg.MessageService: // messageService#2b085862
//  default: panic(v)
//  }
func handleNewMessageOld(gMessage tg.MessageClass) {
	switch message := gMessage.(type) {
	case *tg.MessageEmpty: // messageEmpty#90a6ca84
		// ignore empty messages
		return
	case *tg.Message: // message#85d6cbe2
		handleNormalMessagesOld(message)
		return
	case *tg.MessageService: // messageService#2b085862
		//handleServiceMessages(message)
		return
	}
}

func handleNewMessage(container *enteryManager.WotoContainer) {
	gMessage := container.OriginMessage
	switch message := gMessage.(type) {
	case *tg.MessageEmpty: // messageEmpty#90a6ca84
		// ignore empty messages
		return
	case *tg.Message: // message#85d6cbe2
		container.Message = message
		handleNormalMessages(container)
		return
	case *tg.MessageService: // messageService#2b085862
		container.ServiceMessage = message
		handleServiceMessages(container)
		return
	}
}

func handleNormalMessagesOld(message *tg.Message) {
	if message.Message != "" {
		if len(wotoValues.EnetryMaster) == 0 {
			return
		}
		//go wotoValues.EnetryMaster.TryToRun(message)
		//log.Println("got message: " + message.Message)
	}
}

func handleNormalMessages(container *enteryManager.WotoContainer) {
	if container.Message != nil && container.Message.Message != "" {
		if len(wotoValues.EnetryMaster) == 0 {
			return
		}
		wotoValues.EnetryMaster.TryToRun(container)
		//log.Println("got message: " + message.Message)
	}
}

func handleServiceMessages(container *enteryManager.WotoContainer) {

}

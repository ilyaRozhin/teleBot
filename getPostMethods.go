package main

import (
	"net/http"
	"strconv"
)

type answerGET[T resultGetTypes] struct {
	Ok          bool   `json:"ok"`
	Result      T      `json:"result"`
	Description string `json:"description"`
}

type resultGetTypes interface {
	[]update | user
}

type message struct {
	MessageId                     int64                         `json:"message_id"`
	MessageThreadId               int64                         `json:"message_thread_id"`
	From                          user                          `json:"from"`
	SenderChat                    chat                          `json:"sender_chat"`
	Date                          int64                         `json:"date"`
	Chat                          chat                          `json:"chat"`
	ForwardFrom                   user                          `json:"forward_from"`
	ForwardFromChat               chat                          `json:"forward_from_chat"`
	ForwardFromMessageId          int64                         `json:"forward_from_message_id"`
	ForwardSignature              string                        `json:"forward_signature"`
	ForwardSenderName             string                        `json:"forward_sender_name"`
	ForwardDate                   int64                         `json:"forward_date"`
	IsTopicMessage                bool                          `json:"is_topic_message"`
	IsAutomaticForward            bool                          `json:"is_automatic_forward"`
	ReplyToMessage                *message                      `json:"reply_to_message"` // Надо подумать как считывать без проблем
	ViaBot                        user                          `json:"via_bot"`
	EditDate                      int64                         `json:"edit_date"`
	HasProtectedContent           bool                          `json:"has_protected_content"`
	MediaGroupId                  string                        `json:"media_group_id"`
	AuthorSignature               string                        `json:"author_signature"`
	Text                          string                        `json:"text"`
	Entities                      []messageEntity               `json:"entities"`
	Animation                     animation                     `json:"animation"`
	Audio                         audio                         `json:"audio"`
	Document                      document                      `json:"document"`
	Photo                         []photoSize                   `json:"photo"`
	Sticker                       sticker                       `json:"sticker"`
	Video                         video                         `json:"video"`
	VideoNote                     videoNote                     `json:"video_note"`
	Voice                         voice                         `json:"voice"`
	Caption                       string                        `json:"caption"`
	CaptionEntities               []messageEntity               `json:"caption_entities"`
	HasMediaSpoiler               bool                          `json:"has_media_spoiler"`
	Contact                       contact                       `json:"contact"`
	Dice                          dice                          `json:"dice"`
	Game                          game                          `json:"game"`
	Poll                          poll                          `json:"poll"`
	Venue                         venue                         `json:"venue"`
	Location                      location                      `json:"location"`
	NewChatMembers                []user                        `json:"new_chat_members"`
	LeftChatMember                user                          `json:"left_chat_member"`
	NewChatTitle                  string                        `json:"new_chat_title"`
	NewChatPhoto                  []photoSize                   `json:"new_chat_photo"`
	DeleteChatPhoto               bool                          `json:"delete_chat_photo"`
	GroupChatCreated              bool                          `json:"group_chat_created"`
	SupergroupChatCreated         bool                          `json:"supergroup_chat_created"`
	MessageAutoDeleteTimerChanged messageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed"`
	MigrateToChatId               int64                         `json:"migrate_to_chat_id"`
	MigrateFromChatId             int64                         `json:"migrate_from_chat_id"`
	PinnedMessage                 *message                      `json:"pinned_message"` // Надо подумать как считывать без проблем
	Invoice                       invoice                       `json:"invoice"`
	SuccessfulPayment             successfulPayment             `json:"successful_payment"`
	UserShared                    userShared                    `json:"user_shared"`
	ChatShared                    chatShared                    `json:"chat_shared"`
	ConnectedWebsite              string                        `json:"connected_website"`
	WriteAccessAllowed            writeAccessAllowed            `json:"write_access_allowed"`
	PassportData                  passportData                  `json:"passport_data"`
	ProximityAlertTriggered       proximityAlertTriggered       `json:"proximity_alert_triggered"`
	ForumTopicCreated             forumTopicCreated             `json:"forum_topic_created"`
	ForumTopicEdited              forumTopicEdited              `json:"forum_topic_edited"`
	ForumTopicClosed              forumTopicClosed              `json:"forum_topic_closed"`
	ForumTopicReopened            forumTopicReopened            `json:"forum_topic_reopened"`
	GeneralForumTopicHidden       generalForumTopicHidden       `json:"general_forum_topic_hidden"`
	GeneralForumTopicUnhidden     generalForumTopicUnhidden     `json:"general_forum_topic_unhidden"`
	VideoChatScheduled            videoChatScheduled            `json:"video_chat_scheduled"`
	VideoChatStarted              videoChatStarted              `json:"video_chat_started"`
	VideoChatEnded                videoChatEnded                `json:"video_chat_ended"`
	VideoChatParticipantsInvited  videoChatParticipantsInvited  `json:"video_chat_participants_invited   "`
	WebAppData                    webAppData                    `json:"web_app_data"`
	ReplyMarkup                   inlineKeyboardMarkup          `json:"reply_markup"`
}

type videoChatEnded struct {
	Users []user `json:"users"`
}

type videoChatParticipantsInvited struct {
	Users []user `json:"users"`
}

type webAppData struct {
	Data       string `json:"data"`
	ButtonText string `json:"button_text"`
}

type inlineKeyboardMarkup struct {
	InlineKeyboard []inlineKeyboardButton `json:"inline_keyboard"`
}

type inlineKeyboardButton struct {
	Text                         string                      `json:"text"`
	Url                          string                      `json:"url"`
	CallbackData                 string                      `json:"callback_data"`
	WebApp                       webAppInfo                  `json:"web_app"`
	LoginUrl                     loginUrl                    `json:"login_url"`
	SwitchInlineQuery            string                      `json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string                      `json:"switch_inline_query_current_chat	String"`
	SwitchInlineQueryChosenChat  switchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat"`
	CallbackGame                 callBackGame                `json:"callback_game"`
	Pay                          bool                        `json:"pay"`
}

type webAppInfo struct {
	Url string `json:"url"`
}

type loginUrl struct {
	Url                string `json:"url"`
	ForwardText        string `json:"forward_text"`
	BotUsername        string `json:"bot_username"`
	RequestWriteAccess bool   `json:"request_write_access"`
}

type switchInlineQueryChosenChat struct {
	Query             string `json:"query"`
	AllowUserChats    bool   `json:"allow_user_chats"`
	AllowBotChats     bool   `json:"allow_bot_chats"`
	AllowGroupChats   bool   `json:"allow_group_chats"`
	AllowChannelChats bool   `json:"allow_channel_chats"`
}

type callBackGame struct {
	// Смотрим что присылает потом уже собственно что-то напишем
}

type videoChatStarted struct {
	Duration int64 `json:"duration	Integer"`
}

type videoChatScheduled struct {
	StartDate int64 `json:"start_date	Integer"`
}

type generalForumTopicUnhidden struct {
	//Надо пробовать отправлять запросы и смотреть что возвращает
}

type generalForumTopicHidden struct {
	//Надо пробовать отправлять запросы и смотреть что возвращает
}

type forumTopicReopened struct {
	//Надо пробовать отправлять запросы и смотреть что возвращает
}

type forumTopicClosed struct {
	//Надо пробовать отправлять запросы и смотреть что возвращает
}

type forumTopicEdited struct {
	Name              string `json:"name"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type forumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int64  `json:"icon_color"`
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type writeAccessAllowed struct {
	WebAppName string `json:"web_app_name"`
}

type passportData struct {
	Data        []encryptedPassportElement `json:"data"`
	Credentials encryptedCredentials       `json:"credentials	EncryptedCredentials"`
}

type encryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data"`
	PhoneNumber string         `json:"phone_number"`
	Email       string         `json:"email"`
	Files       []passportFile `json:"files"`
	FrontSide   passportFile   `json:"front_side"`
	ReverseSide passportFile   `json:"reverse_side"`
	Selfie      passportFile   `json:"selfie"`
	Translation []passportFile `json:"translation"`
	Hash        string         `json:"hash"`
}

type encryptedCredentials struct {
	Data   string `json:"data	"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

type passportFile struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FileDate     int64  `json:"file_date"`
}

type proximityAlertTriggered struct {
	Traveler user  `json:"traveler"`
	Watcher  user  `json:"watcher"`
	Distance int64 `json:"distance"`
}

type chatShared struct {
	RequestId int64 `json:"request_id"`
	ChatId    int64 `json:"chat_id"`
}

type userShared struct {
	RequestId int64 `json:"request_id"`
	UserId    int64 `json:"user_id"`
}

type successfulPayment struct {
	Currency                string    `json:"currency"`
	TotalAmount             int64     `json:"total_amount"`
	InvoicePayload          string    `json:"invoice_payload"`
	ShippingOptionId        string    `json:"shipping_option_id"`
	OrderInfo               orderInfo `json:"order_info"`
	TelegramPaymentChargeId string    `json:"telegram_payment_charge_id"`
	ProviderPaymentChargeId string    `json:"provider_payment_charge_id"`
}

type orderInfo struct {
	Name            string          `json:"name"`
	PhoneNumber     string          `json:"phone_number"`
	Email           string          `json:"email"`
	ShippingAddress shippingAddress `json:"shipping_address"`
}

type shippingAddress struct {
	CountryCode string `json:"country_code"`
	State       string `json:"state"`
	City        string `json:"city"`
	StreetLine1 string `json:"street_line1"`
	StreetLine2 string `json:"street_line2"`
	PostCode    string `json:"post_code"`
}

type invoice struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	StartParameter string `json:"start_parameter"`
	Currency       string `json:"currency"`
	TotalAmount    int64  `json:"total_amount"`
}

type messageAutoDeleteTimerChanged struct {
	MessageAutoDeleteTime int64 `json:"message_auto_delete_time"`
}

type dice struct {
	Emoji string `json:"emoji"`
	Value int64  `json:"value"`
}

type game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []photoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []messageEntity `json:"text_entities"`
	Animation    animation       `json:"animation"`
}

type venue struct {
	Location        location `json:"location"`
	Title           string   `json:"title"`
	Address         string   `json:"address"`
	FoursquareId    string   `json:"foursquare_id"`
	FoursquareType  string   `json:"foursquare_type"`
	GooglePlaceId   string   `json:"google_place_id"`
	GooglePlaceType string   `json:"google_place_type"`
}

type poll struct {
	Id                    string          `json:"id"`
	Question              string          `json:"question	String"`
	Options               []poolOption    `json:"options"`
	TotalVoterCount       int64           `json:"total_voter_count"`
	IsClosed              bool            `json:"is_closed"`
	IsAnonymous           bool            `json:"is_anonymous"`
	Type                  string          `json:"type"`
	AllowsMultipleAnswers bool            `json:"allows_multiple_answers"`
	CorrectOptionId       int64           `json:"correct_option_id"`
	Explanation           string          `json:"explanation"`
	ExplanationEntities   []messageEntity `json:"explanation_entities"`
	OpenPeriod            int64           `json:"open_period"`
	CloseDate             int64           `json:"close_date"`
}

type poolOption struct {
	Text       string `json:"text"`
	VoterCount int64  `json:"voter_count"`
}

type audio struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Duration     int64     `json:"duration"`
	Performer    string    `json:"performer"`
	Title        string    `json:"title"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int64     `json:"file_size"`
	ThumbNail    photoSize `json:"thumbnail"`
}

type animation struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Width        int64     `json:"width"`
	Height       int64     `json:"height"`
	Duration     int64     `json:"duration"`
	Thumbnail    photoSize `json:"thumbnail"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int64     `json:"file_size"`
}

type photoSize struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
	FileSize     int64  `json:"file_size"`
}

type messageEntity struct {
	Type          string `json:"type String"`
	Offset        int64  `json:"offset"`
	Length        int64  `json:"length"`
	Url           string `json:"url"`
	User          user   `json:"user"`
	Language      string `json:"language"`
	CustomEmojiId string `json:"custom_emoji_id"`
}

type document struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Thumbnail    photoSize `json:"thumbnail"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int64     `json:"file_size"`
}

type sticker struct {
	FileId           string       `json:"file_id"`
	FileUniqueId     string       `json:"file_unique_id"`
	Type             string       `json:"type"`
	Width            int64        `json:"width"`
	Height           int64        `json:"height"`
	IsAnimated       bool         `json:"is_animated"`
	IsVideo          bool         `json:"is_video"`
	Thumbnail        photoSize    `json:"thumbnail"`
	Emoji            string       `json:"emoji"`
	SetName          string       `json:"set_name"`
	PremiumAnimation file         `json:"premium_animation"`
	MaskPosition     maskPosition `json:"mask_position	MaskPosition"`
	CustomEmojiId    string       `json:"custom_emoji_id"`
	NeedsRepainting  bool         `json:"needs_repainting"`
	FileSize         int64        `json:"file_size"`
}

type video struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Width        int64     `json:"width"`
	Height       int64     `json:"height"`
	Duration     int64     `json:"duration"`
	Thumbnail    photoSize `json:"thumbnail"`
	FileName     string    `json:"file_name"`
	MimeType     string    `json:"mime_type"`
	FileSize     int64     `json:"file_size"`
}

type videoNote struct {
	FileId       string    `json:"file_id"`
	FileUniqueId string    `json:"file_unique_id"`
	Length       int64     `json:"length"`
	Duration     int64     `json:"duration"`
	ThumbNail    photoSize `json:"thumbnail"`
	FileSize     int64     `json:"file_size"`
}

type voice struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	Duration     int64  `json:"duration"`
	MimeType     string `json:"mime_type"`
	FileSize     int64  `json:"file_size"`
}

type contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserId      int64  `json:"user_id"`
	Vcard       string `json:"vcard"`
}

// Надо почитать что с отдачей и принятем файлов
type file struct {
	FileId       string `json:"file_id"`
	FileUniqueId string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	FilePath     string `json:"file_path"`
}

type maskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

type chat struct {
	ID                                 int64           `json:"id"`
	Type                               string          `json:"type"`
	Title                              string          `json:"title"`
	UserName                           string          `json:"username"`
	FirstName                          string          `json:"first_name"`
	LastName                           string          `json:"last_name"`
	IsForum                            bool            `json:"is_forum"`
	Photo                              chatPhoto       `json:"photo"`
	ActiveUsernames                    []string        `json:"active_usernames"`
	EmojiStatusCustomEmojiId           string          `json:"emoji_status_custom_emoji_id"`
	Bio                                string          `json:"bio"`
	HasPrivateForwards                 bool            `json:"has_private_forwards"`
	HasRestrictedVoiceAndVideoMessages bool            `json:"has_restricted_voice_and_video_messages"`
	JointToSendMessages                bool            `json:"join_to_send_messages"`
	JoinByRequest                      bool            `json:"join_by_request"`
	Description                        string          `json:"description"`
	InviteLink                         string          `json:"invite_link"`
	PinnedMessage                      *message        `json:"pinned_message"`
	Permissions                        chatPermissions `json:"permissions"`
	SlowModeDelay                      int64           `json:"slow_mode_delay"`
	MessageAutoDeleteTime              int64           `json:"message_auto_delete_time"`
	HasAggressiveAntiSpamEnabled       bool            `json:"has_aggressive_anti_spam_enabled"`
	HasHiddenMembers                   bool            `json:"has_hidden_members"`
	HasProtectedContent                bool            `json:"has_protected_content"`
	StickerSetName                     string          `json:"sticker_set_name"`
	CanSetStickerSet                   bool            `json:"can_set_sticker_set"`
	LinkedChatId                       int64           `json:"linked_chat_id"`
	Location                           chatLocation    `json:"location"`
}

type chatPermissions struct {
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendAudios         bool `json:"can_send_audios"`
	CanSendDocuments      bool `json:"can_send_documents"`
	CanSendPhotos         bool `json:"can_send_photos"`
	CanSendVideos         bool `json:"can_send_videos"`
	CanSendVideoNotes     bool `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes"`
	CanSendPools          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`
	CanManageTopics       bool `json:"can_manage_topics"`
}

type chatPhoto struct {
	SmallFileId       string `json:"small_file_id"`
	SmallFileUniqueId string `json:"small_file_unique_id"`
	BigFileId         string `json:"big_file_id"`
	BitFileUniqueId   string `json:"big_file_unique_id"`
}

type chatLocation struct {
	Location location `json:"location"`
	Address  string   `json:"address"`
}

type location struct {
	Longitude            float64 `json:"longitude"`
	Latitude             float64 `json:"latitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy"`
	LivePeriod           int64   `json:"live_period"`
	Heading              int64   `json:"heading"`
	ProximityAlertRadius int64   `json:"proximity_alert_radius"`
}

type user struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	UserName                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	IsPremium               bool   `json:"is_premium"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportInlineQueries    bool   `json:"supports_inline_queries"`
}

type inlineQuery struct {
	Id       string   `json:"id	"`
	From     user     `json:"from"`
	Query    string   `json:"query"`
	Offset   string   `json:"offset"`
	ChatType string   `json:"chat_type"`
	Location location `json:"location"`
}

type chosenInlineResult struct {
	ResultId        string   `json:"result_id"`
	From            user     `json:"from"`
	Location        location `json:"location"`
	InlineMessageId string   `json:"inline_message_id"`
	Query           string   `json:"query"`
}

type callbackQuery struct {
	Id              string  `json:"id"`
	From            user    `json:"from"`
	Message         message `json:"message"`
	InlineMessageId string  `json:"inline_message_id	String"`
	ChatInstance    string  `json:"chat_instance"`
	Data            string  `json:"data"`
	GameShortName   string  `json:"game_short_name"`
}

type shippingQuery struct {
	Id              string          `json:"id"`
	From            user            `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress shippingAddress `json:"shipping_address"`
}

type preCheckoutQuery struct {
	Id               string    `json:"id"`
	From             user      `json:"from"`
	Currency         string    `json:"currency"`
	TotalAmount      int64     `json:"total_amount"`
	InvoicePayload   string    `json:"invoice_payload"`
	ShippingOptionId string    `json:"shipping_option_id"`
	OrderInfo        orderInfo `json:"order_info"`
}

type pollAnswer struct {
	PollId    string  `json:"poll_id"`
	User      user    `json:"user"`
	OptionIds []int64 `json:"option_ids"`
}

type chatJoinRequest struct {
	Chat       chat           `json:"chat"`
	From       user           `json:"from"`
	UserChatId int64          `json:"user_chat_id"`
	Date       int64          `json:"date"`
	Bio        string         `json:"bio"`
	InviteLink chatInviteLink `json:"invite_link"`
}

type chatMemberUpdate struct {
	Chat                    chat                   `json:"chat"`
	From                    user                   `json:"from"`
	Date                    int64                  `json:"date"`
	BufferOldChatMember     map[string]interface{} `json:"old_chat_member"`
	BufferNewChatMember     map[string]interface{} `json:"new_chat_member"`
	InviteLine              chatInviteLink         `json:"invite_link"`
	ViaChatFolderInviteLink bool                   `json:"via_chat_folder_invite_link"`
	OldChatMember           chatMember
	NewChatMember           chatMember
}

type chatMember struct {
	ChaMemberOwner          chatMemberOwner
	ChatMemberAdministrator chatMemberAdministrator
	ChatMemberMember        chatMemberMember
	ChatMemberRestricted    chatMemberRestricted
	ChatMemberLeft          chatMemberLeft
	ChatMemberBanned        chatMemberBanned
}

type chatMemberOwner struct {
	Status      string `json:"status"`
	User        user   `json:"user"`
	IsAnonymous bool   `json:"is_anonymous"`
	CustomTitle string `json:"custom_title"`
}

type chatMemberAdministrator struct {
	Status              string `json:"status"`
	User                user   `json:"user"`
	CanBeEdited         bool   `json:"can_be_edited"`
	IsAnonymous         bool   `json:"is_anonymous"`
	CanManageChat       bool   `json:"can_manage_chat"`
	CanDeleteMessages   bool   `json:"can_delete_messages"`
	CanManageVideoChats bool   `json:"can_manage_video_chats"`
	CanRestrictMembers  bool   `json:"can_restrict_members"`
	CanPromoteMembers   bool   `json:"can_promote_members"`
	CanChangeInfo       bool   `json:"can_change_info"`
	CanInviteUsers      bool   `json:"can_invite_users"`
	CanPostMessages     bool   `json:"can_post_messages"`
	CanEditMessages     bool   `json:"can_edit_messages"`
	CanPinMessages      bool   `json:"can_pin_messages"`
	CanManageTopics     bool   `json:"can_manage_topics"`
	CustomTitle         string `json:"custom_title"`
}

type chatMemberMember struct {
	Status string `json:"status"`
	User   user   `json:"user"`
}

type chatMemberRestricted struct {
	Status                string `json:"status"`
	User                  user   `json:"user"`
	IsMember              bool   `json:"is_member"`
	CanSendMessages       bool   `json:"can_send_messages"`
	CanSendAudios         bool   `json:"can_send_audios"`
	CanSendDocuments      bool   `json:"can_send_documents"`
	CanSendPhotos         bool   `json:"can_send_photos"`
	CanSendVideos         bool   `json:"can_send_videos"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes"`
	CanSendPolls          bool   `json:"can_send_polls"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews"`
	CanChangeInfo         bool   `json:"can_change_info"`
	CanInviteUsers        bool   `json:"can_invite_users"`
	CanPinMessages        bool   `json:"can_pin_messages"`
	CanManageTopics       bool   `json:"can_manage_topics"`
	UntilDate             int64  `json:"until_date"`
}

type chatMemberLeft struct {
	Status string `json:"status"`
	User   user   `json:"user"`
}

type chatMemberBanned struct {
	Status    string `json:"status"`
	User      user   `json:"user"`
	UntilDate int64  `json:"until_date"`
}

type chatInviteLink struct {
	InviteLink              string `json:"invite_link"`
	Creator                 user   `json:"creator"`
	CreatesJoinRequest      bool   `json:"creates_join_request"`
	IsPrimary               bool   `json:"is_primary"`
	IsRevoked               bool   `json:"is_revoked"`
	Name                    string `json:"name"`
	ExpireDate              int64  `json:"expire_date"`
	MemberLimit             int64  `json:"member_limit"`
	PendingJoinRequestCount int64  `json:"pending_join_request_count"`
}

// sendToPerson функция отправки ответа пользователю
func (u *update) sendToPerson(text string, token string) {
	chatID := strconv.FormatInt(u.Message.Chat.ID, 10)
	str := "/sendMessage?chat_id=" + chatID + "&text=" + text
	http.Get(urlApi + token + str)
}

// answerPOST - хранит данные возвращаемые методом POST()
type answerPOST struct {
}

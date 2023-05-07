package main

// update представление входящих обновлений.
// NOTE: В любом заданном обновлении может присутствовать
// не более одного необязательного параметра
type update struct {
	UpdateId          int64              `json:"update_id"`
	Message           message            `json:"message"`              // Optional
	EditedMessage     message            `json:"edited_message"`       // Optional
	ChanelPost        message            `json:"channel_post"`         // Optional
	EditedChanelPost  message            `json:"edited_channel_post"`  // Optional
	InlineQuery       inlineQuery        `json:"inline_query"`         // Optional
	ChoseInlineResult chosenInlineResult `json:"chosen_inline_result"` // Optional
	CallbackQuery     callbackQuery      `json:"callback_query"`       // Optional
	ShippingQuery     shippingQuery      `json:"shipping_query"`       // Optional
	PreCheckoutQuery  preCheckoutQuery   `json:"pre_checkout_query"`   // Optional
	Poll              poll               `json:"poll"`                 // Optional
	PollAnswer        pollAnswer         `json:"poll_answer"`          // Optional
	MyChatMember      chatMemberUpdate   `json:"my_chat_member"`       // Optional
	ChatMemberUpdate  chatMemberUpdate   `json:"chat_member"`          // Optional
	ChatJoinRequest   chatJoinRequest    `json:"chat_join_request"`    // Optional
}

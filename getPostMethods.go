package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

// answerGET - сохраняет все данные полученные с помощью метода GET()
type answerGET[T resultTypes] struct {
	Ok     bool `json:"ok"`
	Result T    `json:"result"`
}

// resultTypes - возможные типы возвращаемых значений
// методами GET() необходимые для описания AnswerGET
type resultTypes interface {
	[]updateInfo | telegramPerson
}

// getInfo метод структуры AnswerGET для обработки
// входящей от метода/команды информации в формате JSON
func (a *answerGET[T]) getInfo(command, token string) {
	resp, err := http.Get(urlApi + token + command)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		panic(err)
	}
	err = json.Unmarshal(body, a)
}

// answerPOST - хранит данные возвращаемые методом POST()
type answerPOST struct {
}

// message - структура сообщения, исходящего из телеграм бота
type message struct {
	MessageId int64          `json:"update_id"`
	FromWho   telegramPerson `json:"from"`
	Chat      telegramChat   `json:"chat"`
	Date      int64          `json:"date"`
	Text      string         `json:"text"`
}

// telegramChat - структура для хранения данных о чате
type telegramChat struct {
	Id        int64  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"username"`
	Type      string `json:"type"`
}

// telegramPerson - структура любого пользователя телеграм
type telegramPerson struct {
	Id                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bot"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name"`
	UserName                string `json:"username"`
	LanguageCode            string `json:"language_code"`
	IsPremium               bool   `json:"is_premium"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu"`
	CanJoinGroups           bool   `json:"can_join_groups"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries"`
}

// updateInfo - структура для хранения информации полученной от метода getUpdate
type updateInfo struct {
	UpdateId int64   `json:"update_id"`
	Message  message `json:"message"`
}

// sendToPerson функция отправки ответа пользователю
func (u *updateInfo) sendToPerson(text string, token string) {
	chatID := strconv.FormatInt(u.Message.Chat.Id, 10)
	str := "/sendMessage?chat_id=" + chatID + "&text=" + text
	http.Get(urlApi + token + str)
}

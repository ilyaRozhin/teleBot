package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const str = "https://api.telegram.org/bot"

// resultTypes - возможные типы возвращаемых значений методами GET()
type resultTypes interface {
	[]updateInfo | telegramPerson
}

// AnswerGET - сохраняет все данные полученные с помощью метода GET()
type AnswerGET[T resultTypes] struct {
	Ok     bool `json:"ok"`
	Result T    `json:"result"`
	hash   string
}

// getInfo функция обработки входящей от команды информации в формате JSON
func (a *AnswerGET[T]) getInfo(command, hash string) {
	a.hash = hash
	resp, err := http.Get(str + hash + command)
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

// updateInfo - структура для хранения информации полученной от метода getUpdate
type updateInfo struct {
	UpdateId int64   `json:"update_id"`
	Message  message `json:"message"`
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

func (a *AnswerGET[T]) sendToPerson(str string, upd updateInfo) {
	chatID := strconv.FormatInt(upd.Message.Chat.Id, 10)
	a.getInfo("/sendMessage?chat_id="+chatID+"&text="+str, a.hash)
}

func main() {
	var botToken = "5622217484:AAEuo1G5nDoKwKWPNW2SyyXYihvTRb0b6F4"
	var obj1 AnswerGET[telegramPerson]
	obj1.getInfo("/getMe", botToken)
	var obj AnswerGET[[]updateInfo]
	obj.getInfo("/getUpdates", botToken)
	if cap(obj.Result) == 0 {
		fmt.Println("No,Updates!")
	}
	for _, value := range obj.Result {
		if value.Message.Text == "/start" {
			obj.sendToPerson("Hello, "+value.Message.FromWho.FirstName+" how are you?", value)
		}
	}
}

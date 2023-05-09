package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type communication struct {
	dialogMass []dialog
	uniqUsers  []user
}

type dialog struct {
	firstPerson  user
	secondPerson user
	firstChatId  int64
	secondChatId int64
}

type TelegramBot struct {
	Token         string
	UrlServer     string
	Update        []update
	Data          string
	Communication communication
}

func (t *TelegramBot) getUpdates(param paramGetUpdates) {
	var requestLine = t.UrlServer + "bot" + t.Token + "/getUpdates"
	var safeBuffer answerGET[[]update]
	var notAlone = false
	if param.offset != 0 || param.limit != 0 || param.timeout != 0 || len(param.allowedUpdates) != 0 {
		requestLine += "?"
		if param.offset != 0 {
			requestLine += "offset="
			requestLine += strconv.FormatInt(param.offset, 10)
			notAlone = true
		}
		if param.limit != 0 {
			if notAlone {
				requestLine += "&"
			}
			requestLine += "limit="
			requestLine += strconv.FormatInt(param.limit, 10)
			notAlone = true
		}
		if param.timeout != 0 {
			if notAlone {
				requestLine += "&"
			}
			requestLine += "timeout="
			requestLine += strconv.FormatInt(param.timeout, 10)
			notAlone = true
		}
		if len(param.allowedUpdates) != 0 {
			if notAlone {
				requestLine += "&"
			}
			requestLine += "allowed_updates="
			byteMass, err := json.Marshal(param.allowedUpdates)
			if err != nil {
				panic("Error in getUpdates - 'allowed_updates Marshal Error'")
			}
			requestLine += string(byteMass)
		}
	}
	//fmt.Println(requestLine)
	err := json.Unmarshal(get(requestLine), &safeBuffer)
	if err != nil {
		panic(err)
	}
	if safeBuffer.Ok {
		t.Update = safeBuffer.Result
	} else {
		fmt.Println(safeBuffer.Description)
	}
}

func (t *TelegramBot) sendMessage(param paramSendMessages) {
	var questionFlag = false
	var ampersandFlag = false
	var saferBuffer answerGET[message]
	var requestLine = t.UrlServer + "bot" + t.Token + "/sendMessage"
	if param.chatId != 0 || param.messageThreadId != 0 || param.text != "" {
		requestLine += "?"
		questionFlag = true
		if param.chatId != 0 {
			requestLine += "chat_id="
			requestLine += strconv.FormatInt(param.chatId, 10)
			ampersandFlag = true
		}
		if param.messageThreadId != 0 {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "message_thread_id="
			requestLine += strconv.FormatInt(param.messageThreadId, 10)
			ampersandFlag = true
		}
		if param.text != "" {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "text="
			requestLine += param.text
			ampersandFlag = true
		}

	}
	if param.parseMode != "" || len(param.entities) != 0 || param.disableWebPagePreview != false {
		if !questionFlag {
			requestLine += "?"
			questionFlag = true
		}
		if param.parseMode != "" {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "parse_mode="
			requestLine += param.parseMode
			ampersandFlag = true
		}
		if len(param.entities) != 0 {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "entities="
			byteMass, err := json.Marshal(param.entities)
			if err != nil {
				panic("Error in sendMessage - 'allowed_updates Marshal Error'")
			}
			requestLine += string(byteMass)
			ampersandFlag = true
		}
		if param.disableWebPagePreview != false {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "disable_web_page_preview="
			requestLine += strconv.FormatBool(param.disableWebPagePreview)
			ampersandFlag = true
		}
	}
	if param.disableNotification != false || param.protectContent != false || param.replyToMessageId != 0 {
		if !questionFlag {
			requestLine += "?"
			questionFlag = true
		}
		if param.disableNotification != false {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "disable_notification="
			requestLine += strconv.FormatBool(param.disableNotification)
			ampersandFlag = true
		}
		if param.protectContent != false {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "protect_content="
			requestLine += strconv.FormatBool(param.protectContent)
			ampersandFlag = true

		}
		if param.replyToMessageId != 0 {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "reply_to_message_id="
			requestLine += strconv.FormatInt(param.replyToMessageId, 10)
			ampersandFlag = true
		}
	}
	if param.allowSendingWithoutReply != false || param.replyMarkup != "" {
		if !questionFlag {
			requestLine += "?"
			questionFlag = true
		}
		if param.allowSendingWithoutReply != false {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "allowed_sending_without_reply="
			requestLine += strconv.FormatBool(param.allowSendingWithoutReply)
			ampersandFlag = true
		}
		if param.replyMarkup != "" {
			if ampersandFlag {
				requestLine += "&"
			}
			requestLine += "reply_markup="
			//byteMass, err := json.Marshal(param.replyMarkup)
			//if err != nil {
			//	panic("Error in sendMessage - 'allowed_updates Marshal Error'")
			//}
			requestLine += param.replyMarkup //string(byteMass)
			//ampersandFlag = true
		}
	}
	err := json.Unmarshal(get(requestLine), &saferBuffer)
	if err != nil {
		panic(err)
	}
	if saferBuffer.Ok {
		// Подумать что сделать с принимаемым сообщением
	} else {
		fmt.Println(saferBuffer.Description + "Салам попалам")
		fmt.Println(requestLine)
	}
	//fmt.Println(requestLine)
}

func (t *TelegramBot) setWebhook(url string, certificate inputFile, ipAddress string, maxConnections int, allowedUpdates []string, dropPendingUpdates bool, secretToken string) {
	//Необходимо разработать после реализации кнопок
}

func get(request string) []byte {
	resp, err := http.Get(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		panic(err)
	}
	return body
}

func post(request string) []byte {
	var body io.Reader

	resp, err := http.Post(request, "", body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	byteMass, err2 := io.ReadAll(resp.Body)
	if err2 != nil {
		panic(err2)
	}
	//fmt.Println(string(byteMass))

	return byteMass
}

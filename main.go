package main

import (
	"encoding/json"
	"fmt"
)

// str - строка для хранения начала url
const urlApi = "https://api.telegram.org/"

// Запоминаем %0A - символ перехода на новую строку
func commandHandler(t *TelegramBot) {
	var param paramSendMessages
	param.chatId = t.Update[0].Message.Chat.ID
	switch t.Update[0].Message.Text {
	case "/start":
		param.text = "Привет, @" + t.Update[0].Message.From.UserName + " ты здесь новенький, набирай /help чтобы войти в курс дел."
		t.sendMessage(param)
	case "/end":
		param.text = "Жаль, @" + t.Update[0].Message.From.UserName + ", но иногда приходится прощаться. Рад буду пообщаться снова!"
		t.sendMessage(param)
	case "/help":
		param.text = "*Как пользоваться этим ботом\\?*%0A" +
			"%0A/search чтобы начать поиск нового котика по представленному описанию\\." +
			"%0A/newchat если тебе стало скучно и хочется пообщаться с любителями котят\\." +
			"%0A/endchat новый друг уже нашел вас в соц\\. сетях и потребность в чате отпала не беда\\." +
			"%0A/aboutme вдруг тебе стало интересно кто же сделал такого крутого бота \\. "
		param.parseMode = "MarkdownV2"
		t.sendMessage(param)
	case "/aboutme":
		param.text = "*Кто создатель\\?*%0A" +
			"%0AМеня зовут Илья и я занимаюсь Go\\-разработкой\\, пока что я новичек\\, но усердно работаю над всем чего не знаю\\. " +
			"Если вам захотелось со мной связаться\\, то здесь я оставляю свои контакты\\: " +
			"@KuZy\\_i \\- телеграмм\\."
		param.parseMode = "MarkdownV2"
		t.sendMessage(param)
	case "/search":
		inlineKeyboardHandler(&param, "/search")
		t.sendMessage(param)
	case "/endchat":
		inlineKeyboardHandler(&param, "/endchat")
		t.sendMessage(param)
	case "/startchat":
		inlineKeyboardHandler(&param, "/startchat")
		t.sendMessage(param)
	}
}

func menuButtonConstruct(markup *inlineKeyboardMarkup, rows int, lines int) {
	for i := 0; i < rows; i++ {
		firstRow := make([]inlineKeyboardButton, lines)
		markup.InlineKeyboard = append(markup.InlineKeyboard, firstRow)
	}
}

func inlineKeyboardHandler(param *paramSendMessages, str string) {
	var keyboardMarkup inlineKeyboardMarkup
	switch str {
	case "/search":
		param.text = "*Важные данные о котиках\\.\\.\\.*%0A%0AВыбери необходимое действие с описанием твоего котика\\. " +
			"Не забывай\\, что ты можешь посмотреть его либо изменить\\. Как только описание будет готов можно приступить к поискам" +
			" нажав клавишу *Начать*\\."
		param.parseMode = "MarkdownV2"
		menuButtonConstruct(&keyboardMarkup, 3, 1)
		keyboardMarkup.InlineKeyboard[0][0].Text = "Начать😼"
		keyboardMarkup.InlineKeyboard[0][0].CallbackData = "start_search"
		keyboardMarkup.InlineKeyboard[1][0].Text = "Изменить🧐"
		keyboardMarkup.InlineKeyboard[1][0].CallbackData = "edit"
		keyboardMarkup.InlineKeyboard[2][0].Text = "Посмотреть👁"
		keyboardMarkup.InlineKeyboard[2][0].CallbackData = "look"
		byteMass, err := json.Marshal(keyboardMarkup)
		if err != nil {
			panic(err)
		}
		param.replyMarkup = string(byteMass)
	case "/endchat":
		param.text = "*Чат завершен\\, как тебе собеседник\\?*%0A%0A❤️ \\- безупречный любитель котов\\.%0A🤌 \\- знаток в смешных котах\\." +
			"%0A👎 \\- совершенный профан\\."
		param.parseMode = "MarkdownV2"
		menuButtonConstruct(&keyboardMarkup, 1, 3)
		keyboardMarkup.InlineKeyboard[0][0].Text = "❤️"
		keyboardMarkup.InlineKeyboard[0][0].CallbackData = "love"
		keyboardMarkup.InlineKeyboard[0][1].Text = "🤌"
		keyboardMarkup.InlineKeyboard[0][1].CallbackData = "good"
		keyboardMarkup.InlineKeyboard[0][2].Text = "👎"
		keyboardMarkup.InlineKeyboard[0][2].CallbackData = "bad"
		byteMass, err := json.Marshal(keyboardMarkup)
		if err != nil {
			panic(err)
		}
		param.replyMarkup = string(byteMass)
	case "/startchat":
		param.text = "*Кого ищем\\?*%0A%0AМадам али Мусьё?"
		param.parseMode = "MarkdownV2"
		menuButtonConstruct(&keyboardMarkup, 1, 2)
		keyboardMarkup.InlineKeyboard[0][0].Text = "Мадам"
		keyboardMarkup.InlineKeyboard[0][0].CallbackData = "woman"
		keyboardMarkup.InlineKeyboard[0][1].Text = "Мусьё"
		keyboardMarkup.InlineKeyboard[0][1].CallbackData = "man"
		byteMass, err := json.Marshal(keyboardMarkup)
		if err != nil {
			panic(err)
		}
		param.replyMarkup = string(byteMass)
	}
}

func handlerStartEnd(t *TelegramBot) {
	commandHandler(t) //Обрабатываем команду
}

func callbackKeyboardExecutor(t *TelegramBot) {
	var requestLine = t.UrlServer + "bot" + t.Token + "/answerCallbackQuery"
	var param paramSendMessages
	var keyboard inlineKeyboardMarkup
	switch t.Update[0].CallbackQuery.Data {
	case "edit":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Загружаем прошлое описание ..."
		param.text = "*Хотите изменить\\?*"
		if t.Data != "" {
			param.text += "%0A%0AВаше прошлое описание\\: " + t.Data + "\\."
		} else {
			param.text += "%0A%0AСтарого описания не существует\\, введите новые данные для поиска\\."
		}
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID // (КОСТЫЛЬ) Огромный знак вопроса насчет использования вместо Id чата пользователя ????
		fmt.Println(param.chatId)
		param.protectContent = true
		menuButtonConstruct(&keyboard, 1, 2)
		keyboard.InlineKeyboard[0][0].Text = "Да"
		keyboard.InlineKeyboard[0][0].CallbackData = "yes"
		keyboard.InlineKeyboard[0][1].Text = "Нет"
		keyboard.InlineKeyboard[0][1].CallbackData = "no"
		byteMass, err := json.Marshal(keyboard)
		if err != nil {
			panic(err)
		}
		param.replyMarkup = string(byteMass)
		t.sendMessage(param)
	case "start_search":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Начали поиск ..."
	case "look":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Загружаем прошлое описание ..."
		if t.Data != "" {
			param.text += "%0A%0A*Ваше прошлое описание\\:* " + t.Data + "\\."
		} else {
			param.text += "%0A%0A*Старого описания не существует\\, введите новые данные для поиска\\.*"
		}
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID
		t.sendMessage(param)
	case "love":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Спасибо за ваш отзыв!"
	case "good":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Спасибо за ваш отзыв!"
	case "bad":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Спасибо за ваш отзыв!"
	case "woman":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Уже начали поиск собеседника ..."
	case "man":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Уже начали поиск собеседника ..."
	}
	fmt.Println(string(post(requestLine)))
}

func main() {
	var botToken = "5622217484:AAHqr_GhpMfbfSs38Kc4ybsr85CMTh61wGE"
	var tBot TelegramBot
	var param paramGetUpdates

	tBot.Token = botToken
	tBot.UrlServer = urlApi

	param.limit = 0
	param.timeout = 0
	param.allowedUpdates = []string{}

	for true {
		tBot.getUpdates(param)
		if len(tBot.Update) != 0 {
			param.offset = tBot.Update[len(tBot.Update)-1].UpdateId
			handlerStartEnd(&tBot)
			if tBot.Update[0].CallbackQuery.Data != "" {
				callbackKeyboardExecutor(&tBot)
			}
			param.offset++
		}

	}
}

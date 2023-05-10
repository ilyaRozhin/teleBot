package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// str - строка для хранения начала url
const urlApi = "https://api.telegram.org/"

// Запоминаем %0A - символ перехода на новую строку
func commandHandler(t *TelegramBot, command string, chatID int64) {
	var param paramSendMessages
	param.chatId = chatID
	switch command {
	case "/start":
		param.text = "Привет, @" + t.Update[0].Message.From.UserName + " ты здесь новенький, набирай /help чтобы войти в курс дел."
		t.sendMessage(param)
	case "/end":
		param.text = "Жаль, @" + t.Update[0].Message.From.UserName + ", но иногда приходится прощаться. Рад буду пообщаться снова!"
		t.sendMessage(param)
	case "/help":
		param.text = "*Как пользоваться этим ботом\\?*%0A" +
			"%0A*/search* чтобы начать поиск нового котика по представленному описанию\\." +
			"%0A*/startchat* если тебе стало скучно и хочется пообщаться с любителями котят\\." +
			"%0A*/endchat* новый друг уже нашел вас в соц\\. сетях и потребность в чате отпала\\." +
			"%0A*/aboutme* вдруг тебе стало интересно кто же сделал такого крутого бота\\. " +
			"%0A*/menu* мое главное меню, со всеми доступными функциями\\."
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
	case "/menu":
		mainMenu(&param)
		t.sendMessage(param)
	}
}

func menuButtonConstruct(markup *inlineKeyboardMarkup, rows int, lines int) {
	for i := 0; i < rows; i++ {
		firstRow := make([]inlineKeyboardButton, lines)
		markup.InlineKeyboard = append(markup.InlineKeyboard, firstRow)
	}
}

func mainMenu(param *paramSendMessages) {
	var keyboardMarkup inlineKeyboardMarkup
	param.text = "*Меню*%0A%0AВыберите какое\\-нибудь действие\\! \\(Лучше всего что\\-то с котиками\\)\\. Если нужна помощь, то набери /help\\."
	param.parseMode = "MarkdownV2"
	menuButtonConstruct(&keyboardMarkup, 2, 2)
	keyboardMarkup.InlineKeyboard[0][0].Text = "Поиск"
	keyboardMarkup.InlineKeyboard[0][0].CallbackData = "search"
	keyboardMarkup.InlineKeyboard[1][0].Text = "Помощь"
	keyboardMarkup.InlineKeyboard[1][0].CallbackData = "help"
	keyboardMarkup.InlineKeyboard[0][1].Text = "Общение"
	keyboardMarkup.InlineKeyboard[0][1].CallbackData = "communicate"
	keyboardMarkup.InlineKeyboard[1][1].Text = "О создателе"
	keyboardMarkup.InlineKeyboard[1][1].CallbackData = "about_me"
	byteMass, err := json.Marshal(keyboardMarkup)
	if err != nil {
		panic(err)
	}
	param.replyMarkup = string(byteMass)
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
		param.text = "*Кого ищем\\?*%0A%0AТы можешь найти себе новых друзей, с которыми можно будет делиться котами, неплохая идея не так ли\\?"
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
	commandHandler(t, t.Update[0].Message.Text, t.Update[0].Message.Chat.ID) //Обрабатываем команду
}

func convertForRequest(str string) string {
	var result string
	for _, value := range str {
		if value == ' ' {
			result += "%20"
		} else {
			result += string(value)
		}
	}
	return string(result)
}

func (t *TelegramBot) sendParamMessage(str string) {
	var param paramSendMessages
	param.text = str
	param.parseMode = "MarkdownV2"
	param.chatId = t.Update[0].Message.Chat.ID
	t.sendMessage(param)
}

func (t *TelegramBot) sendPhoto(url string) bool {
	byteMass := post(t.UrlServer + "bot" + t.Token + "/sendPhoto?" + "chat_id=" + strconv.FormatInt(t.Update[0].CallbackQuery.From.ID, 10) +
		"&photo=" + url)
	//fmt.Println(t.UrlServer + "bot" + t.Token + "/sendPhoto?" + "chat_id=" + strconv.FormatInt(t.Update[0].CallbackQuery.From.ID, 10) +
	//	"&photo=" + url)
	fmt.Println(string(byteMass))
	var sConvFirst = &answerSET[bool]{}
	var sConvSecond = &answerSET[message]{}
	err := json.Unmarshal(byteMass, sConvFirst)
	err2 := json.Unmarshal(byteMass, sConvSecond)
	if err != nil && err2 != nil {
		panic(err)
	}
	//fmt.Println(*sConvFirst, *sConvSecond)
	if sConvFirst.Ok || sConvSecond.Ok {
		return true
	}
	return false
}

func callbackKeyboardExecutor(t *TelegramBot) {
	var requestLine = t.UrlServer + "bot" + t.Token + "/answerCallbackQuery"
	var param paramSendMessages
	var keyboard inlineKeyboardMarkup
	var newPhoto []unsplashPhoto
	switch t.Update[0].CallbackQuery.Data {
	case "look":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Загружаем прошлое описание ..."
		if t.Data != "" {
			param.text += "%0A%0A_*Ваше прошлое описание\\: " + t.Data + "\\.*_"
		} else {
			param.text += "%0A%0A_*Старого описания не существует\\, введите новые данные для поиска\\.*_"
		}
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID
		t.sendMessage(param)
		if t.Data != "" {
			mainMenu(&param)
			t.sendMessage(param)
			break
		}
		fallthrough
	case "edit":
		if requestLine[len(requestLine)-1] != '.' {
			requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Загружаем прошлое описание ..."
		}
		param.text = "*Хотите изменить\\?*"
		if t.Data != "" {
			param.text += "%0A%0AВаше прошлое описание\\: " + t.Data + "\\."
		} else {
			param.text += "%0A%0AСтарого описания не существует\\, введите новые данные для поиска\\."
		}
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID // (КОСТЫЛЬ) Огромный знак вопроса насчет использования вместо Id чата Id пользователя ????
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
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id
		if t.Data == "" {
			t.sendMessage(paramSendMessages{chatId: t.Update[0].CallbackQuery.From.ID, parseMode: "MarkdownV2", text: "_*Вы еще не ввели описание котика\\. Попробуйте нажать кнопку Изменить\\.*_"})
		} else {
			t.Data = convertForRequest(t.Data)
			byteMass := get("https://api.unsplash.com/photos/random?query=" + t.Data + "&count=1&content_filter=high" + "&client_id=6b_vt-YJ4ntrxcK6-CxNx1uUS_glNcBdn-LD1Bd2sXU")
			err := json.Unmarshal(byteMass, &newPhoto)
			if err != nil {
				panic(err)
			}
			//fmt.Println(newPhoto[0].Links.Download)
			if t.sendPhoto(newPhoto[0].Links.Download) {
				param.text = "*Повторим еще разочек\\? Или ты уже устал\\?*"
				requestLine += "&text=Картинка найдена!"
			} else {
				param.text = "*Произошла ошибка\\(такое бывает\\), повторить поиск\\?*"
				requestLine += "&text=Что-то пошло не так!"
			}
			param.chatId = t.Update[0].CallbackQuery.From.ID
			param.parseMode = "MarkdownV2"
			menuButtonConstruct(&keyboard, 1, 2)
			keyboard.InlineKeyboard[0][0].Text = "Да"
			keyboard.InlineKeyboard[0][0].CallbackData = "yes_cat"
			keyboard.InlineKeyboard[0][1].Text = "Нет"
			keyboard.InlineKeyboard[0][1].CallbackData = "no_cat"
			byteMass2, err2 := json.Marshal(keyboard)
			if err2 != nil {
				panic(err2)
			}
			param.replyMarkup = string(byteMass2)
			t.sendMessage(param)
		}
	case "yes_cat":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id
		t.Data = convertForRequest(t.Data)
		byteMass := get("https://api.unsplash.com/photos/random?query=" + t.Data + "&count=1&content_filter=high" + "&client_id=6b_vt-YJ4ntrxcK6-CxNx1uUS_glNcBdn-LD1Bd2sXU")
		err := json.Unmarshal(byteMass, &newPhoto)
		if err != nil {
			panic(err)
		}
		//fmt.Println(newPhoto[0].Links.Download)
		if t.sendPhoto(newPhoto[0].Links.Download) {
			param.text = "*Повторим еще разочек\\? Или ты уже устал\\?*"
			requestLine += "&text=Картинка найдена!"
		} else {
			param.text = "*Произошла ошибка\\(такое бывает\\), повторить поиск\\?*"
			requestLine += "&text=Что-то пошло не так!"
		}
		param.chatId = t.Update[0].CallbackQuery.From.ID
		param.parseMode = "MarkdownV2"
		menuButtonConstruct(&keyboard, 1, 2)
		keyboard.InlineKeyboard[0][0].Text = "Да"
		keyboard.InlineKeyboard[0][0].CallbackData = "yes_cat"
		keyboard.InlineKeyboard[0][1].Text = "Нет"
		keyboard.InlineKeyboard[0][1].CallbackData = "no_cat"
		byteMass2, err2 := json.Marshal(keyboard)
		if err2 != nil {
			panic(err2)
		}
		param.replyMarkup = string(byteMass2)
		t.sendMessage(param)
	case "no_cat":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Переходим в главное меню..."
		param.chatId = t.Update[0].CallbackQuery.From.ID
		mainMenu(&param)
		t.sendMessage(param)
	case "love":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Спасибо за ваш отзыв!"
		param.text = "*Не работает\\!*" + "%0A%0AПрошу прощения, но данный сервис еще не разработан до конца\\." +
			" Если у вас есть интересные идеи пишите сюда\\: @KuZy\\_i\\. Чтобы перейти в главное меню тыкните сюда /menu\\."
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID
		t.sendMessage(param)
		//mainMenu(&param)
		//t.sendMessage(param)
	case "good":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Спасибо за ваш отзыв!"
		param.text = "*Не работает\\!*" + "%0A%0AПрошу прощения, но данный сервис еще не разработан до конца\\." +
			" Если у вас есть интересные идеи пишите сюда\\: @KuZy\\_i\\. Чтобы перейти в главное меню тыкните сюда /menu\\."
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID
		t.sendMessage(param)
		//mainMenu(&param)
		//t.sendMessage(param)
	case "bad":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Спасибо за ваш отзыв!"
		param.text = "*Не работает\\!*" + "%0A%0AПрошу прощения, но данный сервис еще не разработан до конца\\." +
			" Если у вас есть интересные идеи пишите сюда\\: @KuZy\\_i\\. Чтобы перейти в главное меню тыкните сюда /menu\\."
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID
		t.sendMessage(param)
		//mainMenu(&param)
		//t.sendMessage(param)
	case "woman":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Уже начали поиск собеседника ..."
		param.text = "*Не работает\\!*" + "%0A%0AПрошу прощения, но данный сервис еще не разработан до конца\\." +
			" Если у вас есть интересные идеи пишите сюда\\: @KuZy\\_i\\. Чтобы перейти в главное меню тыкните сюда /menu\\."
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID
		t.sendMessage(param)
		//mainMenu(&param)
		//t.sendMessage(param)
		//t.LastButtonCommand = "woman"
	case "man":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Уже начали поиск собеседника ..."
		param.text = "*Не работает\\!*" + "%0A%0AПрошу прощения, но данный сервис еще не разработан до конца\\." +
			" Если у вас есть интересные идеи пишите сюда\\: @KuZy\\_i\\. Чтобы перейти в главное меню тыкните сюда /menu\\."
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID
		t.sendMessage(param)
		//mainMenu(&param)
		//t.sendMessage(param)
	case "yes":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Ждёмс ..."
		param.text = "_*Введите новое описание\\:*_"
		param.parseMode = "MarkdownV2"
		param.chatId = t.Update[0].CallbackQuery.From.ID
		t.Data = "nil"
		t.sendMessage(param)
	case "no":
		requestLine += "?callback_query_id=" + t.Update[0].CallbackQuery.Id + "&text=Переходим в главное меню..."
		param.chatId = t.Update[0].CallbackQuery.From.ID
		mainMenu(&param)
		t.sendMessage(param)
	case "search":
		commandHandler(t, "/search", t.Update[0].CallbackQuery.From.ID)
	case "help":
		commandHandler(t, "/help", t.Update[0].CallbackQuery.From.ID)
	case "communicate":
		commandHandler(t, "/startchat", t.Update[0].CallbackQuery.From.ID)
	case "aboutme":
		commandHandler(t, "/aboutme", t.Update[0].CallbackQuery.From.ID)
	}
	post(requestLine)
}

func main() {
	var botToken = "5622217484:AAG5JMG14W2b7k-_l4lY-GGzIg6gJbouJ5k"
	var tBot TelegramBot
	var paramMessage paramSendMessages
	tBot.Token = botToken
	var param paramGetUpdates

	tBot.UrlServer = urlApi

	var lastCommand string

	param.limit = 0
	param.timeout = 0
	param.allowedUpdates = []string{}

	for true {
		tBot.getUpdates(param)
		if len(tBot.Update) != 0 {
			param.offset = tBot.Update[len(tBot.Update)-1].UpdateId
			if tBot.Data == "nil" {
				tBot.Data = tBot.Update[0].Message.Text
				tBot.sendParamMessage("_*Новое описание успешно сохранено\\!\\!\\!*_")
				paramMessage.chatId = tBot.Update[0].Message.Chat.ID
				mainMenu(&paramMessage)
				tBot.sendMessage(paramMessage)
			}
			handlerStartEnd(&tBot)
			lastCommand = tBot.Update[0].Message.Text
			if tBot.Update[0].CallbackQuery.Data != "" {
				callbackKeyboardExecutor(&tBot)
			}
			fmt.Println(lastCommand)
			//fmt.Println(tBot.Data)
			param.offset++
		}
	}
}

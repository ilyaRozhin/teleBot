package main

import "fmt"

// str - строка для хранения начала url
const urlApi = "https://api.telegram.org/bot"

type TelegramBot struct {
	botToken    string
	updateInfo  answerGET[[]Update]
	meInfo      answerGET[telegramPerson]
	postInfo    answerPOST
	answerQuery queryBotTypes[[]Update]
}

// initBot инициализация экземпляра структуры TelegramBot
func (t *TelegramBot) initBot(token string) {
	t.botToken = token
	t.meInfo.getInfo("/getMe", t.botToken)
	t.updateInfo.getInfo("/getUpdates", t.botToken)
}

func (t *TelegramBot) updateAnswerQuery() bool {
	var oldEntries = t.updateInfo.Result
	t.updateInfo.getInfo("/getUpdates", t.botToken)
	fmt.Println(oldEntries)
	/*
		fmt.Println(t.updateInfo.Result)
		if flagUpdate, newEntries := updateQ(oldEntries, t.updateInfo.Result); flagUpdate {
			t.answerQuery.initQuery(newEntries)
			return true
		}
	*/
	return false
}

func (t *TelegramBot) handlerUpdates() {
	var size = t.answerQuery.length
	var bufferNode = t.answerQuery.start
	for i := 0; i < size; i++ {
		switch bufferNode.info.Message.Text {
		case "/start":
			bufferNode.info.sendToPerson("Привет,"+bufferNode.info.Message.FromWho.UserName+". Это телеграм бот,"+t.meInfo.Result.UserName, t.botToken)
		case "/end":
			bufferNode.info.sendToPerson("Прощай,"+bufferNode.info.Message.FromWho.UserName+"!", t.botToken)
		case "Привет":
			bufferNode.info.sendToPerson("Салам,"+bufferNode.info.Message.FromWho.UserName+"!", t.botToken)
		}
		bufferNode = bufferNode.next
		t.answerQuery.deleteFirst()
	}
}

func main() {
	var botToken = "5622217484:AAEuo1G5nDoKwKWPNW2SyyXYihvTRb0b6F4"
	var firstBot TelegramBot
	firstBot.initBot(botToken)

	for true {
		if firstBot.updateAnswerQuery() {
			firstBot.handlerUpdates()
		}
	}
}

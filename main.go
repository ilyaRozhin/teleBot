package main

// str - строка для хранения начала url
const urlApi = "https://api.telegram.org/bot"

type TelegramBot struct {
	botToken    string
	updateInfo  answerGET[[]updateInfo]
	meInfo      answerGET[telegramPerson]
	postInfo    answerPOST
	answerQuery queryBotTypes[[]updateInfo]
}

// initBot инициализация экземпляра структуры TelegramBot
func (t *TelegramBot) initBot(token string) {
	t.botToken = token
	t.meInfo.getInfo("/getMe", t.botToken)
	t.updateInfo.getInfo("/getUpdates", t.botToken)
	t.answerQuery.initQuery(t.updateInfo.Result)
}

func main() {
	var botToken = "5622217484:AAEuo1G5nDoKwKWPNW2SyyXYihvTRb0b6F4"
	var firstBot TelegramBot
	firstBot.initBot(botToken)
	firstBot.answerQuery.end.info.sendToPerson("Здарова как дела?", firstBot.botToken)
}

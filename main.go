package main

import "fmt"

// str - строка для хранения начала url
const urlApi = "https://api.telegram.org/"

func main() {
	var botToken = "5622217484:AAHqr_GhpMfbfSs38Kc4ybsr85CMTh61wGE"
	var tBot TelegramBot
	tBot.Token = botToken
	tBot.UrlServer = urlApi
	//Разобраться с передачей JSON массива
	//NOTE(08/05/23): оказалось все передаю правильно
	var i int64
	i = 202666147
	for true {
		tBot.getUpdates(i, 0, 0, []string{})
		if len(tBot.Update) != 0 {
			fmt.Println(tBot.Update[0].Message.Text, tBot.Update[0].UpdateId)
			i = tBot.Update[len(tBot.Update)-1].UpdateId
			i++
		}
	}
}

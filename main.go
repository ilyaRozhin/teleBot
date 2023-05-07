package main

// str - строка для хранения начала url
const urlApi = "https://api.telegram.org/bot"

func main() {
	var botToken = "5622217484:AAEuo1G5nDoKwKWPNW2SyyXYihvTRb0b6F4"
	var tBot TelegramBot
	tBot.Token = botToken
	tBot.UrlServer = urlApi
	//Разобраться с передачей JSON массива
	tBot.getUpdates(1, 1, 2, []string{"\"message\"", "\"edited_channel_post\""})
	//fmt.Println(tBot.Update)
}

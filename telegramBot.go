package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type TelegramBot struct {
	Token     string
	UrlServer string
	Update    []update
}

func (t *TelegramBot) getUpdates(offset int64, limit int64, timeout int64, allowedUpdates []string) {
	var requestLine = t.UrlServer + t.Token + "/getUpdates"
	var safeBuffer answerGET[[]update]
	var notAlone = false
	if offset != 0 || limit != 0 || timeout != 0 || len(allowedUpdates) != 0 {
		requestLine += "?"
		if offset != 0 {
			requestLine += "offset="
			requestLine += strconv.FormatInt(offset, 10)
			notAlone = true
		}
		if limit != 0 {
			if notAlone {
				requestLine += "&"
			}
			requestLine += "limit="
			requestLine += strconv.FormatInt(limit, 10)
			notAlone = true
		}
		if timeout != 0 {
			if notAlone {
				requestLine += "&"
			}
			requestLine += "timeout="
			requestLine += strconv.FormatInt(timeout, 10)
			notAlone = true
		}
		if len(allowedUpdates) != 0 {
			if notAlone {
				requestLine += "&"
			}
			requestLine += "allowed_updates="
			requestLine += "["
			requestLine += allowedUpdates[0]
			for i := 1; i < len(allowedUpdates); i++ {
				requestLine += "," + allowedUpdates[i]
			}
			requestLine += "]"
		}
	}
	fmt.Println(requestLine)
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

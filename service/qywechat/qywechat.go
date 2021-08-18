package qywechat

import (
	"encoding/json"
	"fmt"
	"gitee.com/dtapps/go-library/service/qywechat/config"
	"gitee.com/dtapps/go-library/service/qywechat/message"
	"gitee.com/dtapps/go-library/utils/djson"
	"io/ioutil"
	"net/http"
	"strings"
)

type QyBot struct {
	Key string
}

type response struct {
	Errcode   int64  `json:"errcode"`
	Errmsg    string `json:"errmsg"`
	Type      string `json:"type"`
	MediaId   string `json:"media_id"`
	CreatedAt string `json:"created_at"`
}

func (bot *QyBot) Send(msg message.Message) (response, error) {
	var response response
	qyUrl := fmt.Sprintf("%s?key=%s", config.Api, bot.Key)
	toString, err := djson.MarshalToString(msg)
	if err != nil {
		return response, err
	}
	resp, e := http.Post(qyUrl, "application/json", strings.NewReader(toString))
	if e != nil {
		return response, e
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	e = json.Unmarshal(body, &response)
	if e != nil {
		return response, e
	}
	return response, nil
}
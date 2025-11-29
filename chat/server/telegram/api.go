package telegram

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	pb "server/chat"
)

const (
	token    = "XXX"
	chat_id  = "XXX"
	topic_id = 0 // integer
)

type Message struct {
	ChatID          string `json:"chat_id"`
	MessageThreadID int    `json:"message_thread_id,omitempty"`
	Text            string `json:"text"`
	// parseMode : default, Markdown, HTML
	ParseMode             string `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool   `json:"disable_web_page_preview"`
	DisableNotification   bool   `json:"disable_notification"`
	ReplyToMessageID      *int32 `json:"reply_to_message_id,omitempty"`
}

func SendMessage(in *pb.ChatMsg) error {
	// text
	text, err := json.Marshal(in)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	message := Message{}
	message.ChatID = chat_id
	message.MessageThreadID = topic_id
	message.Text = string(text)
	// message.ParseMode = "Markdown"
	// message.ParseMode = "HTML"
	message.DisableWebPagePreview = false
	message.DisableNotification = false

	// payload := strings.NewReader("{\"text\":\"hello world\",\"disable_web_page_preview\":false,\"disable_notification\":false,\"reply_to_message_id\":null,\"chat_id\":\"6161672218\"}")
	// req, _ := http.NewRequest("POST", url, payload)
	payload, err := json.Marshal(message)
	if err != nil {
		log.Print(err)
		return err
	}
	url := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", token)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))

	req.Header.Add("accept", "application/json")
	req.Header.Add("User-Agent", "Telegram Bot SDK - (https://github.com/irazasyed/telegram-bot-sdk)")
	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print(err)
		return err
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Print(err)
		return err
	}

	log.Println(string(body))
	return nil
}

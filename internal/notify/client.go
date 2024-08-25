package notify

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/shoshtari/divar-notifier/internal/configs"
)

type Notifier interface {
	SendMessage(text string) (int, error)
	EditMessage(messageID int, text string) error
}

type NotifierImp struct {
	config configs.SectionNotifier
	client http.Client
}

// sendRequest send http request with body req to url and decode response to res
func (n NotifierImp) sendRequest(url string, req any, res any) error {

	reqBody, err := json.Marshal(req)
	if err != nil {
		return errors.Wrap(err, "cant marshal request")
	}

	httpRes, err := n.client.Post(url, "Application/Json", bytes.NewReader(reqBody))
	if err != nil {
		return errors.Wrap(err, "can't send request")
	}

	if httpRes.StatusCode != http.StatusOK {
		return errors.New("result status is not 200")
	}

	resData, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return errors.Wrap(err, "can't read data")
	}

	err = json.Unmarshal(resData, &res)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal response")
	}
	return nil
}

func (n NotifierImp) SendMessage(text string) (int, error) {
	url := fmt.Sprintf("%v/%v/sendMessage", n.config.BaseUrl, n.config.Token)
	req := SendMessageRequest{
		ChatID: n.config.Target,
		Text:   text,
	}

	var res SendMessageResponse

	err := n.sendRequest(url, req, &res)
	if err != nil {
		return -1, errors.Wrap(err, "can't send request")
	}

	if !res.Ok {
		return -1, errors.New(fmt.Sprint("res is not ok, res data is", res))
	}

	return res.Result.MessageID, nil
}

func (n NotifierImp) EditMessage(messageID int, text string) error {
	url := fmt.Sprintf("%v/%v/editMessageText", n.config.BaseUrl, n.config.Token)
	req := EditMessageRequest{
		ChatID:    n.config.Target,
		MessageID: messageID,
		Text:      text,
	}

	var res SendMessageResponse // since we only need ok field, I didn't define new struct

	err := n.sendRequest(url, req, &res)
	if err != nil {
		return errors.Wrap(err, "can't send request")
	}

	if !res.Ok {
		return errors.New(fmt.Sprint("res is not ok, res data is", res))
	}

	return nil
}

func NewNotifier(config configs.SectionNotifier) Notifier {
	return NotifierImp{
		config: config,
	}
}

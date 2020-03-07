package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

const url = "https://hooks.slack.com/services/"
const ContentType = "application/json"
var _token = ""
var _httpClient = http.DefaultClient

type (
	Field struct {
		Title string `json:"title"`
		Value string `json:"value"`
		Short bool   `json:"short"`
	}
	Message struct {
		Text    string  `json:"text"`
		Pretext string  `json:"pretext"`
		Color   string  `json:"color"`
		Fields  []Field `json:"fields"`
	}
	Builder struct {
		Message *Message
	}
)

func init() {
	_token = os.Getenv("SLACK_TOKEN")
	if len(_token) == 0 {
		fmt.Printf("slack env var = `SLACK_TOKEN`")
	}
}

func SetToken(token string) {
	_token = token
}

func NewBuilder() *Builder {
	return &Builder{
		Message: &Message{},
	}
}

func (b *Builder) Text(text string) *Builder {
	b.Message.Text = text
	return b
}

func (b *Builder) Pretext(text string) *Builder {
	b.Message.Pretext = text
	return b
}

func (b *Builder) Color(color string) *Builder {
	b.Message.Color = color
	return b
}

func (b *Builder) AddField(title, value string) *Builder {
	if len(b.Message.Fields) == 0 {
		b.Message.Fields = []Field{{
			Title: title,
			Value: value,
		}}
	} else {
		b.Message.Fields = append(b.Message.Fields, Field{
			Title: title,
			Value: value,
		})
	}
	return b
}

func (b Builder) Send() error {

	if b.Message == nil {
		return errors.New("send message is nil")
	}

	data, err := json.Marshal(b.Message)

	if err != nil {
		return err
	}

	req, err := _httpClient.Post(url+_token, ContentType, bytes.NewReader(data))
	if err != nil {
		return err
	}

	if req.StatusCode != http.StatusOK {
		return fmt.Errorf("[slack] response code = %d, message = %s", req.StatusCode, req.Status)
	}
	return nil
}

// 发出错误代码？
func SendError(err error) {
	err = errors.WithStack(err)
	b := NewBuilder().Pretext("Error Message").Text(fmt.Sprintf("%+v", err))
	_ = b.Send()
}
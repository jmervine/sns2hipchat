package hipchat

import (
	//"fmt"
	"github.com/jmervine/hipchat-sns-relay/Godeps/_workspace/src/github.com/tbruyelle/hipchat-go/hipchat"
	"github.com/jmervine/hipchat-sns-relay/config"
	//"net/http"
	"net/url"
)

//const HIPCHAT_ENDPOINT = "https://%s/v1/rooms/message"

//var cfg *config.Config

type Hipchat struct {
	client  *hipchat.Client
	request *hipchat.NotificationRequest
	room    string
	//room    *hipchat.Room
}

func New(cfg *config.Config) (*Hipchat, error) {
	//cfg = c
	h := Hipchat{}

	h.client = hipchat.NewClient(cfg.Token)
	//h.endpoint = fmt.Sprintf(HIPCHAT_ENDPOINT, cfg.Host)

	u, err := url.Parse(cfg.Host)
	if err != nil {
		return nil, err
	}
	h.client.BaseURL = u

	h.request = &hipchat.NotificationRequest{
		Color:         cfg.Color,
		Notify:        cfg.Notify,
		MessageFormat: cfg.MessageFormat,
	}

	//r, _, err := h.client.Room.Get(cfg.RoomID)
	//if err != nil {
	//return nil, err
	//}
	//h.room = r
	h.room = cfg.RoomID

	return &h, nil
}

func (h *Hipchat) Post(msg string) error {
	_, err := h.client.Room.Notification(h.room, h.request)
	return err
}

//type Hipchat struct {
//endpoint string
//params   url.Values
//auth     url.Values
//}

//func New(c *config.Config) *Hipchat {
//cfg = c
//h := Hipchat{}

//h.endpoint = fmt.Sprintf(HIPCHAT_ENDPOINT, cfg.Host)

//h.params = url.Values{
//"room_id": {cfg.RoomID},
//"from":    {cfg.From},
//}

//if cfg.Notify {
//h.params.Set("notify", "1")
//}

//if cfg.Color != "" {
//h.params.Add("color", cfg.Color)
//}

//if cfg.MessageFormat != "" {
//h.params.Add("message_format", cfg.MessageFormat)
//}

//h.auth = url.Values{
//"auth_token": {cfg.Token},
//}

//return &h
//}

//func (h *Hipchat) Post(msg string) error {
//h.params.Add("message", msg)

//hipchat, err := url.Parse(h.endpoint)
//if err != nil {
//return err
//}
//hipchat.RawQuery = h.auth.Encode()

//if cfg.Debug {
//fmt.Printf("at=hipchat hipchat=%+v\n", hipchat)
//}

//resp, err := http.PostForm(hipchat.String(), h.params)
//if err != nil {
//return err
//}

//defer resp.Body.Close()

//if resp.StatusCode != 200 {
//return fmt.Errorf("HipChat error with status %d", resp.StatusCode)
//}

//return nil
//}

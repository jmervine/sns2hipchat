package hipchat

import (
	"fmt"
	"github.com/jmervine/hipchat-sns-relay/config"
	"net/http"
	"net/url"
)

const HIPCHAT_ENDPOINT = "https://%s/v1/rooms/message"

type Hipchat struct {
	endpoint string
	params   url.Values
	auth     url.Values
}

func New(cfg *config.Config) *Hipchat {
	h := Hipchat{}

	h.endpoint = fmt.Sprintf(HIPCHAT_ENDPOINT, cfg.Host)

	h.params = url.Values{
		"room_id": {cfg.RoomID},
		"from":    {cfg.From},
	}

	if cfg.Notify {
		h.params.Set("notify", "1")
	}

	if cfg.Color != "" {
		h.params.Add("color", cfg.Color)
	}

	if cfg.MessageFormat != "" {
		h.params.Add("message_format", cfg.MessageFormat)
	}

	h.auth = url.Values{
		"auth_token": {cfg.Token},
	}

	return &h
}

func (h *Hipchat) Post(msg string) error {
	h.params.Add("message", msg)

	hipchat, err := url.Parse(h.endpoint)
	if err != nil {
		return err
	}
	hipchat.RawQuery = h.auth.Encode()

	resp, err := http.PostForm(hipchat.String(), h.params)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("hipchat returned %d, expected 200", resp.StatusCode)
	}

	return nil
}

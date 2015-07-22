package hipchat

/* TODO:
 * ----
 * I really should learn out to use interfaces properly. I have a strong feeling
 * that the way I'm handling v1 vs. v2 (especially in server/server.go) could
 * be greatly simplified via an interface. I did a little reading and tried a few
 * things, but couldn't get it to work and have time restrictions for using this
 * codebase at the moment. Hence me leaving it the way it is.
 */

import (
	"fmt"
	"github.com/jmervine/sns2hipchat/Godeps/_workspace/src/github.com/tbruyelle/hipchat-go/hipchat"
	"github.com/jmervine/sns2hipchat/config"
	"net/http"
	"net/url"
)

const HIPCHAT_V1_ENDPOINT = "https://%s/v1/rooms/message"
const HIPCHAT_V2_ENDPOINT = "https://%s/v2/"

// these can persist
var v1 *HipchatV1
var v2 *HipchatV2

type HipchatV2 struct {
	client  *hipchat.Client
	request *hipchat.NotificationRequest
}

// NewV2 should only be called once, but support being called with
// each request with minimal overhead.
func NewV2(cfg *config.Config) (*HipchatV2, error) {
	if v2 == nil {
		v2 = &HipchatV2{}
		v2.client = hipchat.NewClient(cfg.Token)
		u, err := url.Parse(fmt.Sprintf(HIPCHAT_V2_ENDPOINT, cfg.Host))
		if err != nil {
			return nil, err
		}
		v2.client.BaseURL = u

		v2.request = &hipchat.NotificationRequest{
			Color:         cfg.Color,
			Notify:        cfg.Notify,
			MessageFormat: cfg.MessageFormat,
		}
	} else {
		v2.request.Message = ""
		v2.client.Room = nil
	}

	return v2, nil
}

func (h *HipchatV2) Post(room, msg string) error {
	h.request.Message = msg
	_, err := h.client.Room.Notification(room, h.request)
	return err
}

type HipchatV1 struct {
	endpoint string
	client   *url.URL
	request  url.Values
	auth     url.Values
	debug    bool
}

// NewV1 should only be called once, but support being called with
// each request with minimal overhead.
func NewV1(cfg *config.Config) (*HipchatV1, error) {
	if v1 == nil {
		v1 = &HipchatV1{}

		v1.endpoint = fmt.Sprintf(HIPCHAT_V1_ENDPOINT, cfg.Host)
		v1.request = url.Values{
			"from": {cfg.From},
		}

		if cfg.Notify {
			v1.request.Set("notify", "1")
		}

		if cfg.Color != "" {
			v1.request.Add("color", cfg.Color)
		}

		if cfg.MessageFormat != "" {
			v1.request.Add("message_format", cfg.MessageFormat)
		}

		v1.auth = url.Values{
			"auth_token": {cfg.Token},
		}

		v1.debug = cfg.Debug

		var err error
		v1.client, err = url.Parse(v1.endpoint)
		if err != nil {
			return nil, err
		}
		v1.client.RawQuery = v1.auth.Encode()

	} else {
		v1.request.Del("room_id")
		v1.request.Del("message")
	}

	return v1, nil
}

func (h *HipchatV1) Post(room, msg string) error {
	h.request.Add("room_id", room)
	h.request.Add("message", msg)

	if h.debug {
		fmt.Printf("at=client client=%+v\n", h.client)
	}

	resp, err := http.PostForm(h.client.String(), h.request)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("client error with status %d", resp.StatusCode)
	}

	return nil
}

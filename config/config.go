package config

import (
	"fmt"
	"strconv"

	"github.com/jmervine/sns2hipchat/Godeps/_workspace/src/gopkg.in/codegangsta/cli.v1"
)

// slightly modified version of
// https://github.com/codegangsta/cli/blob/v1.2.0/help.go#L13
var AppHelpTemplate = `Name:
    {{.Name}} - {{.Usage}}

Usage:
    {{.Name}} [args...]

Version:
    {{.Version}}

Options:
    {{range .Flags}}{{.}}
    {{end}}
`

type Config struct {
	Token             string
	Addr              string
	Host              string
	Test              bool
	Debug             bool
	HipchatAPIVersion int // 1 || 2
	Formatter         string

	// HipchatMessageRequest
	// https://www.hipchat.com/docs/api/method/rooms/message
	// room id or name
	Rooms []string
	From  string
	//Message string

	// html or text, default is html
	MessageFormat string

	// ping people?
	Notify bool

	// yellow (default) / red / green / purple / gray / random
	Color string
}

func Parse(args []string) (cfg *Config) {
	// use custom help template
	cli.AppHelpTemplate = AppHelpTemplate

	app := cli.NewApp()

	app.Version = "1.0.0"
	app.Name = "sns2hipchat"
	app.Usage = "Simple AWS/SNS HTTP{S} endpoint relay to HipChat"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "addr, a",
			Value:  "",
			Usage:  "http listener address",
			EnvVar: "ADDR",
		},
		cli.StringFlag{
			Name:   "port, p",
			Value:  "3000",
			Usage:  "listener port",
			EnvVar: "PORT",
		},
		cli.StringFlag{
			Name:   "token, t",
			Value:  "",
			Usage:  "[required] hipchat api token",
			EnvVar: "HIPCHAT_TOKEN",
		},
		cli.StringSliceFlag{
			Name:   "room, r",
			Value:  &cli.StringSlice{},
			Usage:  "[required] target hipchat room",
			EnvVar: "HIPCHAT_ROOMS",
		},
		cli.StringFlag{
			Name:   "from, f",
			Value:  "Amazon SNS",
			Usage:  "displayed hipchat sender",
			EnvVar: "HIPCHAT_FROM",
		},
		cli.StringFlag{
			Name:   "format",
			Value:  "html",
			Usage:  "hipchat html or text format",
			EnvVar: "HIPCHAT_FORMAT",
		},
		cli.StringFlag{
			Name:   "formatter, F",
			Value:  "raw",
			Usage:  "hipchat message formatter [raw|basic|alarm]",
			EnvVar: "HIPCHAT_FORMATTER",
		},
		cli.StringFlag{
			Name:   "notify, n",
			Value:  "true",
			Usage:  "ping people in hipchat",
			EnvVar: "HIPCHAT_NOTIFY",
		},
		cli.StringFlag{
			Name:   "color, c",
			Value:  "yellow",
			Usage:  "hipchat message color",
			EnvVar: "HIPCHAT_COLOR",
		},
		cli.StringFlag{
			Name:   "host, H",
			Value:  "api.hipchat.com",
			Usage:  "target hipchat api host",
			EnvVar: "HIPCHAT_HOST",
		},
		cli.BoolFlag{
			Name:   "debug",
			Usage:  "enable debug logging",
			EnvVar: "DEBUG",
		},
		cli.BoolFlag{
			Name:   "test",
			Usage:  "enable /test and disable posting to hipchat",
			EnvVar: "TEST",
		},
		cli.IntFlag{
			Name:   "api, A",
			Value:  1,
			Usage:  "hipchat api version",
			EnvVar: "HIPCHAT_API_VERSION",
		},
	}

	app.Action = func(c *cli.Context) {
		var notify bool
		if ok, err := strconv.ParseBool(c.String("notify")); err != nil {
			fmt.Printf("Invalid notify value: %s\n\n", c.String("notify"))
			return
		} else {
			notify = ok
		}

		token := c.String("token")
		if token == "" {
			fmt.Printf("Hipchat Token Required. See '--help' for details.")
			return
		}

		var rooms []string
		rooms = c.StringSlice("room")

		cfg = &Config{
			Token: token,
			Addr: fmt.Sprintf("%s:%s",
				c.String("addr"),
				c.String("port")),
			Rooms:             rooms,
			From:              c.String("from"),
			Formatter:         c.String("formatter"),
			MessageFormat:     c.String("format"),
			Notify:            notify,
			Color:             c.String("color"),
			Host:              c.String("host"),
			Debug:             c.Bool("debug"),
			Test:              c.Bool("test"),
			HipchatAPIVersion: c.Int("api"),
		}
	}

	app.Run(args)

	return cfg
}

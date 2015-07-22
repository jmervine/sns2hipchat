# hipchat-sns-rela

**SNS HTTP{S} Endpoint Relay to HipChat**

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Usage

```
go get -v github.com/jmervine/hipchat-sns-relay
$ hipchat-sns-relay --help
```


## Heroku Deployment

"Deploy to Heroku" button on the top of this README file is recommended way to set up this app to [Heroku](https://www.heroku.com/).

```bash
heroku create
```

If you're not familiar with using Go on Heroku, see [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go)

## Configuration

To make this app up and work properly, you need to setup the following environment variables:

```bash
# required
heroku config:add HIPCHAT_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
heroku config:add HIPCHAT_ROOM="target room"

# optional (defaults listed)
heroku config:add HIPCHAT_FROM="Amazon SNS"
heroku config:add HIPCHAT_FORMAT=html
heroku config:add HIPCHAT_NOTIFY=true
heroku config:add HIPCHAT_COLOR=yellow
heroku config:add HIPCHAT_HOST=api.hipchat.com
```

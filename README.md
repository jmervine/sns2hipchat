## (AWS) sns2hipchat [![Build Status](https://travis-ci.org/jmervine/sns2hipchat.svg?branch=master)](https://travis-ci.org/jmervine/sns2hipchat)

> #### Simple AWS/SNS HTTP{S} endpoint relay to HipChat

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Usage

```
go get -v github.com/jmervine/sns2hipchat
$ sns2hipchat --help
```


## Heroku Deployment

"Deploy to Heroku" button on the top of this README file is recommended way to set up this app to [Heroku](https://www.heroku.com/).

```bash
heroku create
```

If you're not familiar with using Go on Heroku, see [Getting Started with Go on Heroku](https://devcenter.heroku.com/articles/getting-started-with-go)

#### Configuration

To make this app up and work properly, you need to setup the following environment variables:

> See `--help` or `app.json` for details.

**Required**

```bash
heroku config:add HIPCHAT_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
# The API auth token for your HipChat account

heroku config:add HIPCHAT_ROOM="target room"
# The HipChat room to send messages to
```

**Optional** (defaults listed)

```bash
heroku config:add HIPCHAT_FROM="Amazon SNS"
# The HipChat sender messages come from
# - v1 only, v2 will show the key owner

heroku config:add HIPCHAT_FORMAT=html
# The HipChat message format

heroku config:add HIPCHAT_NOTIFY=true
# Tell HipChat to notify people on sending messages

heroku config:add HIPCHAT_COLOR=yellow
# The HipChat message color

heroku config:add HIPCHAT_HOST=api.hipchat.com
# The HipChat target host

heroku config:add HIPCHAT_API_VERSION=1
# The HipChat API Version

heroku config:add DEBUG=false
# Verbose logging for debugging SNS and hipchat errors
```

## Development

**Testing with Docker**

```bash
$ go get -v github.com/jmervine/dockme
$ dockme -C Dockme.test.yml
```


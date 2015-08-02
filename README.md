## (AWS) sns2hipchat [![Build Status](https://travis-ci.org/jmervine/sns2hipchat.svg?branch=master)](https://travis-ci.org/jmervine/sns2hipchat)

> #### Simple AWS/SNS HTTP{S} endpoint relay to HipChat

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy)

## Usage

```
go get -v github.com/jmervine/sns2hipchat
$ sns2hipchat --help
```

## Setting up SNS subscript

1. Create SNS Topic (via AWS console)
2. Subscribe SNS Topic to your endpoint (via AWS console)
    1. Select topic
    2. Choose "Subscribe to topic"
    3. Choose "Protocol" (`http` or `https` accordingly)
    4. Enter endpoint url, e.g.:
        * `http://mysnsendpoint.herokuapp.com/`
        * Or with rooms: `http://mysnsendpoint.herokuapp.com/?room=1&room=2`
3. Confirm subscription:
    1. Select topic (in AWS Console)
    2. Choose "Request Confirmation"
    3. Choose "Confirm Subscription"
    4. View logs for confirmation url (in terminal), e.g.:
        * `heroku logs -a mysnsendpoint | grep "SubscribeURL detected:"`
    5. Paste in subscription url from logs

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
```

**Optional** (defaults listed)

```bash
heroku config:add HIPCHAT_ROOMs=<ROOMID>,<ROOMID>
# The default HipChat room(s) to send messages to

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

# running units
$ dockme go test ./...

# running in test mode
$ dockme

# then hit http://localhost:3000/test
```


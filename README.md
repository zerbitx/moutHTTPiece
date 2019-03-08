# moutHTTPiece 

A simple little server to reply with whatever you'd like when you want to test what you're client will do, and don't care about the real server(s).

### Installation (requires go)
```
export GOBIN=`go env GOPATH`/bin
go get -u github.com/zerbitx/mouthttpiece 
```

### Running

From `mouthttpiece --help`

```bash
Usage of mouthttpiece:
  -contentType string
        Content type header to be returned (default "application/json")
  -echo
        Reply with the request it was given
  -port int
        Default port (default 8000)
  -reply string
        What would you like the server to reply with for EVERY Request: default is json of {"status": "OK"}
  -statusCode int
        What status code to reply with by default (default 200)
```

So, if you want a server that always replies with a 418 with "I'm a little...well you know" on port 8312

`echoserver -port 8312 -statusCode 418 -reply="I'm a little...well you know"` 

### These options are also configurable by environment variables

- PORT          (port to start on)
- REPLY_DEFAULT (string of the standard response body)
- ECHO_RESPONSE (true|false)
- STATUS_CODE   (int : status code to return by default)
- CONTENT_TYPE  (string : default Content-Type header)

### Optionally change per request with the following query params

- responseBody
- statusCode
- contentType

For example 
`curl -v http://localhost:8312?statusCode=502&contentType=application/whatever&responseBody=great`

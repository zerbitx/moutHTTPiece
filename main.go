package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"

	"github.com/kelseyhightower/envconfig"
)

const (
	statusCodeQueryParam = "statusCode"
	responseQueryParam   = "responseBody"
	contentTypeParam     = "contentType"
)

type responseBits struct {
	reply       string
	statusCode  int
	contentType string
}

type Env struct {
	Port              int    `envconfig:"PORT" default:"8000"`
	ReplyDefault      string `envconfig:"REPLY_DEFAULT", default:"{"status": "OK"}"`
	Echo              bool   `envconfig:"ECHO_RESPONSE" default:"false"`
	StatusCodeDefault int    `envconfig:"STATUS_CODE" default:"200"`
	ContentType       string `envconfig:"CONTENT_TYPE" default:"application/json"`
}

func main() {
	var (
		replyDefault       string
		echo               bool
		port               int
		statusCodeDefault  int
		contentTypeDefault string
	)

	e := Env{}
	envconfig.Process("", &e)

	flag.StringVar(&replyDefault, "reply", e.ReplyDefault, `What would you like the server to reply with for EVERY Request: default is json of {"status": "OK"}`)
	flag.BoolVar(&echo, "echo", e.Echo, `Reply with the request it was given`)
	flag.IntVar(&port, "port", e.Port, "Default port")
	flag.IntVar(&statusCodeDefault, "statusCode", e.StatusCodeDefault, "What status code to reply with by default")
	flag.StringVar(&contentTypeDefault, "contentType", e.ContentType, "Content type header to be returned")

	flag.Parse()

	serveMux := http.NewServeMux()

	serveMux.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		bytes, err := httputil.DumpRequest(req, true)

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
			return
		}

		// Echo the incoming request
		fmt.Println(string(bytes))

		rb := responseBits{
			reply:       replyDefault,
			statusCode:  statusCodeDefault,
			contentType: contentTypeDefault,
		}

		err = setupResponse(&rb, req.URL.Query())
		if err != nil {
			fmt.Println(err)
		}
		rw.Header().Set("Content-Type", rb.contentType)

		rw.WriteHeader(rb.statusCode)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Println(err)
		}

		defer func() {
			if err := req.Body.Close(); err != nil {
				fmt.Println(err)
			}
		}()

		if echo {
			rw.Write(bytes)
			return
		}

		rw.Write([]byte(rb.reply))
	}))

	fmt.Printf("Listening on %d\n", port)
	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", port), serveMux))
}

func setupResponse(config *responseBits, query url.Values) error {
	if qv := query[statusCodeQueryParam]; len(qv) == 1 {
		statusCode, err := strconv.Atoi(qv[0])

		config.statusCode = statusCode
		if err != nil {
			return err
		}
	}

	if qv := query[responseQueryParam]; len(qv) == 1 {
		config.reply = qv[0]
	}

	if qv := query[contentTypeParam]; len(qv) == 1 {
		config.contentType = qv[0]
	}

	return nil
}

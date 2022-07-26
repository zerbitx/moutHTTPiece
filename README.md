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


```mermaid
graph TB %% comments
  %% Entity[Text]
  ID-1[Node 1]
  ID-2>Node 2]
  ID-3(Node 3 <br> text)
  %% Entity--Entity
  ID-1---ID-2
  ID-1 --> ID-3
  %% Entity--Text--Entity
  ID-2--Link between 2 and 3---ID-3
  ID-3-->|Action from 3 to 1|ID-1
  ID-3 -- "Action from 3 to 2. p/w: '_-!#$%^&*+=?,\'" --> ID-2
  ID-4 
  %% Complex cases
  A[Hard edge] -->|Link text| B(Round edge)
  ID-1---ID-2(Text)
  B --> C{Text}
  C -->|One| D[Text]
  A(A) --> B(B)
  C[/C/] --> D>D]
  %% class/classDef
  classDef blue fill:#08f,stroke:#fff;
  class ID-1 blue
  class ID-1,ID-2 red
  %% click
  click ID-1 "https://github.com" "Tooltip text" %% comments
  click ID-2 alert "Tooltip for a callback"
  %% subgraph
  subgraph A subgraph
    ID-4{Node 4}
    ID-5((fa:fa-spinner))
    ID-6["Node 6 (same #quot;shape#quot;)"]
    ID-4-.->ID-5
    ID-5 -. Action from 5 to 4 .-> ID-4
    ID-5==>ID-6
    ID-6 == Action from 6 to 5 ==> ID-5
  end
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

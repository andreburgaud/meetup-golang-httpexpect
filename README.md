# GoMN Meetup - httpexpect

## Folders Description

* `1_gotesting`: Golang tests (no server involved) of an Add function using table driven data and subtests

* `2_go_web_servers`: Different approaches of serving HTTP requests using the Golang builtin `net/http` package
  * server1: Define a HTTP Handler and use ServeHTTP
  * server2: Convert a function to HTTP handler using `http.HandlerFunc`
  * server3: Convert a function implicitly to `http.HandlerFunc`
  * server4: Use `http.HandleFunc` to register functions as HTTP handlers

* `3_httptest`: Examples using `net/http/httptest`

* `4_httpexpect`: Examples using `httpexpect`
  * `add_server`: testing of the Add function served via HTTP
  * `fastapi`: testing of a FastAPI server using Golang and `httpexpect`
  * `self_contained`: server and test in the same file
  * `separate_processes`: test a server started as a separate process

* `slides`: Presentation slides in PDF format


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

## FastAPI Example

FastAPI is a Python high performance web framework compatible with Python 3.6+

To execute the FastAPI server in directory `4_httpexpect/fastapi/web` you first need to prepare a Python virtual environment. The instructions below assume that you have Python 3 installed:

```
$ python3 -m venv .venv
$ . .venv/bin/activate
$ pip install fastapi[all] # only once
$ uvicorn main:app
```

The installation of fastapi is done only once. Subsequently, the folloing should be sufficient to activate the environment and start the web server:

```
$ . .venv/bin/activate
$ uvicorn main:app
```

The `Makefile` included in the same folder shows other options to start the server.

After starting the server, you can point your browser to http://localhost:8000/docs to display the Swagger UI.

To test using `httpexpect`, in the directory `4_httpexpect/fastapi/test`, execute the following command:

```
$ go test -v
```

The test folder is a Go mod project. If you experience any error related to the `httpexpect`, fetch the package manually with the followng command, and execute the test again:

```
$ go get github.com/gavv/httpexpect/v2
```

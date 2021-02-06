# TIL some HTTP!

## Redirect
There is a HTTP status code called [301 Moved Permanently](https://en.wikipedia.org/wiki/HTTP_301), which, when returned with a Location header will automatically be redirected in browsers.
Start the server in the Go program, and go to [localhost:8080/redirect] in your browser.
It will automatically redirect to google.com.
curl will also do this if you use it with the `--location` option.

[![301 Moved Permanently](https://http.cat/301 "301 Moved Permanently")](https://http.cat/301)

## File downloads
The Content-Type and Content-Disposition headers can be used to tell a browser to download a file instead of displaying it (see `content` below).
Start the server and go to [localhost:8080/content] in your browser.
It should trigger a file download, and the file will be named `data.csv`.

## Go code

```go
package main

import (
	"context"
	"log"
	"net/http"
	"os"
)

var interruptChannel = make(chan os.Signal, 1)

func redirect(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Location", "https://google.com")
	w.WriteHeader(http.StatusMovedPermanently)
}

func content(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/csv; charset=UTF-8")
	w.Header().Set("Content-Disposition", `attachment; filename="data.csv"`)
	w.Write([]byte(`a,b
1,2
`))
}

func main() {
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/content", content)
	server := http.Server{Addr: ":8080"}
	go server.ListenAndServe()
	<-interruptChannel
	server.Shutdown(context.Background())
}
```

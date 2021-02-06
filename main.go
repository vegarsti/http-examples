package main

import (
	"context"
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

package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("안녕. 에디킴"))
	})

	http.ListenAndServe(":8088", nil)
}

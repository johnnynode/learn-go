package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // 此处 http.NewServeMux() 是较底层的东西，相当于http.ListenAndServe中的handler
    mux.Handle("/", &myHandler{})

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {}

func (* myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: " + r.URL.String()) // 默认访问 localhost:8080 会输出 URL: /
}

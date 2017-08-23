package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux() // 此处 http.NewServeMux() 是较底层的东西，相当于http.ListenAndServe中的handler
    mux.Handle("/", &myHandler{})
	mux.HandleFunc("/hello", sayHello)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {}

func (* myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: " + r.URL.String()) // 默认访问 localhost:8080 会输出 URL: /
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HelloWorld version 2") // 备注即使没有用到r 也需要传递上
}
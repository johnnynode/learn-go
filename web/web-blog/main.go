package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// StripPrefix 是去除前缀，然后 http.Dir 去定位要访问的目录
	mux.Handle("/static/", http.StripPrefix("/static/",
		http.FileServer(http.Dir(wd))))

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {}

func (* myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "URL: " + r.URL.String()) // 默认访问 localhost:8080 会输出 URL: /
}
package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

var mux map[string] func(w http.ResponseWriter, r *http.Request) // 定义一个map

func main() {
	server := http.Server{
		Addr: ":8080",
		Handler : &myHandler{},
		ReadTimeout: 5 * time.Second,
	}

	mux = make(map[string] func(w http.ResponseWriter, r *http.Request))
	mux["/hello"] = sayHello
	mux["/bye"] = sayBye


	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

type myHandler struct {}

func (* myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 使用 map 的 ok pattern
	if h,ok := mux[r.URL.String()]; ok {
		h(w, r) // 直接转发
		return
	}
	io.WriteString(w, "URL: " + r.URL.String()) // 默认访问 localhost:8080 会输出 URL: /
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HelloWorld version sayHello") // 备注即使没有用到r 也需要传递上
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "HelloWorld version sayBye") // 备注即使没有用到r 也需要传递上
}
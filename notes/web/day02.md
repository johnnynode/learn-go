### __Web 学习笔记__

- beego 控制器中的方法向客户端返回字符串 ：
 ```
 h.Ctx.WriteString("Hello World")

 ```
 运行：$ `go run yourApp.go`
 打开浏览器：localhost:8080
 显示HelloWorld

- 自定义实现http功能
 ```
  func main() {
  	// 设置路由
  	http.HandleFunc("/", sayHello)
  	err := http.ListenAndServe(":8080", nil)
  	if err!=nil{
  		log.Fatal(err)
  	}
  }

  func sayHello(w http.ResponseWriter, r *http.Request) {
  	io.WriteString(w, "Hello World version1")
  }

 ```

- 使用 http.NewServeMux 来处理路由
 ```
  func main() {
  	mux := http.NewServeMux() // 此处 http.NewServeMux() 是较底层的东西
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

 ```

-

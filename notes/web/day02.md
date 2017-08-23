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


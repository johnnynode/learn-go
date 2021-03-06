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
  var mux map[string] func(w http.ResponseWriter, r *http.Request) // 定义一个map

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

 ```

- 使用更底层的 http.Server 来处理请求示例
 ```
  func main() {
  	server := http.Server{
  		Addr: ":8080",
  		Handler : &myHandler{},
  		ReadTimeout: 5 * time.Second,
  	}

  	err := server.ListenAndServe()
  	if err != nil {
  		log.Fatal(err)
  	}
  }

  type myHandler struct {}

  func (* myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  	io.WriteString(w, "URL: " + r.URL.String()) // 默认访问 localhost:8080 会输出 URL: /
  }

 ```

- 通过map来保存注册的handler, 再进行底层ServeHTTP进行转发 [效率最高的方法]
 ```
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

 ```

 - 创建一个静态文件服务器
  ```
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

  ```

- 使用beego中的bee工具初始化项目
 * 命令：$ `go get github.com/beego/bee` , $ `bee new myApp`,  $ `cd myApp` $ `bee run myApp`
 * 项目跑起, 默认端口8080

- 获取 beego 配置文件信息的方法
  * 注意 go 在1.8 版本中未找到使用默认参数来获取配置文件的方法 如下
    1. $ `beego.AppName`
    2. $ `strconv.Itoa(beego.HttpPort)`
    3. $ `beego.RunMode`
    4. 以上的API不存在
 ```
  appName := beego.AppConfig.String("appname")
  httpPort := beego.AppConfig.String("httpport")
  runMode := beego.AppConfig.String("runmode")

  // 刷新 http;//localhost:8080 后台打印出：appname:  myApp  httpport:  8080  runmode:  dev
  fmt.Println("appname: ", appName, " httpport: ", httpPort, " runmode: ", runMode)

 ```

- 设置beego的日志级别
 ```
  	beego.SetLevel(beego.LevelInformational)
  	beego.Trace("trace 1") // 未被输出，因为trace级别小于info级别
  	beego.Info("info 1")
  	beego.Trace("trace 2") // 未被输出，因为trace级别小于info级别
  	beego.Info("info 2")

 ```
 以上只输出了info的日志，而trace级别低于所设置的级别，所以未被输出
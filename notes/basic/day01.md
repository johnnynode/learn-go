### Go 是什么

> Go 是一门并发支持、垃圾回收的编译型系统编程语言，旨在创造一门具有在静态编译语言的高性能和动态语言的高效开发之间拥有良好平衡点的一门编程语言。

### Go 主要特点

- 类型安全和内存安全
- 以非常直观和极低代价的方案实现高并发
- 高效的垃圾回收机制
- 快速编译（同时解决C语言中头文件太多的问题）
- 为多核计算机提供性能提升的方案
- UTF-8编码支持

### Go 的价值
- [Go at Google: Language Design in the Service of Software Engineering](https://talks.golang.org/2012/splash.article)
- [Go 在谷歌：以软件工程为目的的语言设计](http://studygolang.com/resources/69)

### Go 的世纪应用
- Youtube （谷歌）
- 七牛云储存及其旗下网盘服务 （Q盘）
- 爱好者开发的Go论坛及博客
- 已在服务端使用Go开发的著名企业：谷歌、盛大、七牛、360等
- 其他海量开源项目：go-wiki、GoDoc、Go Language Resources

### Go 语言的环境变量相关配置
- 安装好Go并配置好环境变量 (此处省略步骤, easy!)
- 打开检测命令`go env` 找到配置的GOPATH所在目录
 * bin (存放变异后生成的可执行文件)
 * pkg (存放编译后生成的包文件)
 * src (存放项目源码)

### Go 常用命令
- $`go get`     : 获取远程包(需提前安装git或hg)
- $`go run`     : 直接运行程序
- $`go build`   : 测试编译，检查是否有编译错误
- $`go fmt`     : 格式化源码
- $`go install` : 编译包文件并编译整个程序
- $`go test`    : 运行测试文件
- $`go doc`     : 查看文档 (会在本地建立起一个文档服务器)

### 推荐编辑器
- goland
- vscode
- sublime
- atom
- ...



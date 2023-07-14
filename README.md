# Gout Web Framework(Gin_demo)  

**Gout支持功能:** 

- 支持GET、POST、PUT、DELETE等HTTP请求
- 简单的log、recovery中间件
- 封装了Context上下文
- Trie树实现路由映射(支持动态路由匹配)
- 分组路由

## 入门
### 先决条件  
- **[Go] (https://go.dev/)**: 最新版本  

### 安装Gout
```
$ go get -u github.com/notacommonperson/gin_demo
```
### 运行Gout 
首先你需要导入 Gout 包来使用 Gout，一个最简单的例子如下“example.go”：
```go
package main

import (
  "net/http"

  "github.com/notacommonperson/gin_demo"
)

func main() {
  r := gout.Default()
  r.GET("/ping", func(c *gout.Context) {
    c.JSON(http.StatusOK, gout.H{
      "msg": "pong",
    })
  })
  r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
```
之后使用命令运行示例
```
# run example.go and visit 0.0.0.0:8080/ping on browser
$ go run example.go
```



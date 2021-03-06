# tinylogger
一个轻量到极致的日志记录器。 自带一个简单的 Info 与 Error 的记录器, 输出到标准输出. 当然也可以自定义。

主要适用的场景应该是在一些小辅助工具的编写上。

![](https://img.shields.io/badge/golang-v0.0.1-blue.svg)
![](https://img.shields.io/github/license/zjxpcyc/tinylogger.svg)

## 安装与使用

**安装**
```golang
// x.y.z 为版本号
go get github.com/zjxpcyc/tinylogger@vx.y.x
```

**使用**
```golang

// NewLogger 不填默认使用 os.StdOut、os.StdErr 输出, 也就是输出到屏幕上
var l tinylogger.LogService = tinylogger.NewLogger()

// 如果传入 io.Writer , 那么会写入到指定的地方
// 下面是示例写入文件
f, _ := os.OpenFile("test.log", os.O_RDWR|os.O_CREATE, 0755)
defer f.close()
l2 := tinylogger.NewLogger(f)

// 调用方式
// Info
l.Info("这是一段普通的日志")

// Error
l2.Error("这是一段有错误的日志")
```

调用成功一般会写入类似如下的信息
```bash
[2018-12-03 17:17:36] [INFO] [xxxx.go:46] 这是一段普通的日志 
[2018-12-03 17:17:36] [ERR] [xxxx.go:47] 这是一段有错误的日志
```
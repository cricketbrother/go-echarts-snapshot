# go-echarts-snapshot

一个用来将echarts html文件转换成图片的Go模块

### 准备工作

1.安装nodejs并配置环境变量

2.按照[https://pptr.dev/](https://pptr.dev/)安装puppeteer

### 安装

```shell
# go.mod

require github.com/cricketbrother/go-echarts-snapshot
```

### 用法

```golang
package main

import "github.com/cricketbrother/go-echarts-snapshot/snapshot"

func main() {
	conf := snapshot.SetConfig(
		"/path/to/node_modules/puppeteer",
		"/path/to/chrome",
		"canvas", 0, 0)
	err := snapshot.Make(conf, "file:///path/to/echarts.html", "/path/to/echarts.png")
	if err != nil {
		fmt.Println(err.Error())
	}
}
```
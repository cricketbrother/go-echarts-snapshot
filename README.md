# go-echarts-snapshot

A Go module for converting echarts html file into image

### Preparation
1.Install nodejs and configure environment variables

2.Follow [https://pptr.dev/](https://pptr.dev/) to install puppeteer

### Installation
```shell
# go.mod

require github.com/cricketbrother/go-echarts-snapshot
```

### Usage
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
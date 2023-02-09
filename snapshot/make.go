package snapshot

import (
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
)

type Config struct {
	PuppeteerPath string
	BrowserPath   string
	Renderer      string
	Width         int
	Height        int
}

func (c *Config) Init(puppeteerPath string, browserPath string, renderer string, width int, height int) {
	c.PuppeteerPath = puppeteerPath
	c.BrowserPath = browserPath
	c.Renderer = renderer
	c.Width = width
	c.Height = height
}

// SetConfig 设置Puppeteer路径和Browser路径
//
// puppeteerPath: Puppeteer路径
//
// browserPath: Browser路径
//
// renderer: 渲染模式(canvas或svg)
//
// width: 页面宽度
//
// height: 页面高度
func SetConfig(puppeteerPath string, browserPath string, renderer string, width int, height int) *Config {
	if renderer != "svg" {
		renderer = "canvas"
	}
	c := new(Config)
	c.Init(puppeteerPath, browserPath, renderer, width, height)
	return c
}

// Make 将HTML文件转换为图片
//
// conf: 配置信息
//
// src: 本地HTML（file:///）或在线网页（http(s)://）
//
// dst: 图片路径
func Make(conf *Config, src string, dst string) error {
	fmt.Println(runtime.Caller(0))
	var cmd *exec.Cmd
	if conf.Width == 0 && conf.Height == 0 {
		cmd = exec.Command("node", "snapshot.js", conf.PuppeteerPath, conf.BrowserPath, conf.Renderer, src, dst)
	} else {
		cmd = exec.Command("node", "snapshot.js", conf.PuppeteerPath, conf.BrowserPath, conf.Renderer, src, dst, strconv.Itoa(conf.Width), strconv.Itoa(conf.Height))
	}
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

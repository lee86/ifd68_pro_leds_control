package main

import (
	"flag"
	"fmt"
	"github.com/getlantern/systray"
	"os"
	"os/signal"
	"syscall"
)

var productIDGD = []byte{0x002d, 0x002c}

const (
	vendorIDGD uint16 = 0x31d6
)

var testIs = flag.Bool("t", false, "测试开关")
var sig, sigAu, sigAua, sigAuaa = make(chan bool), make(chan bool), make(chan bool), make(chan bool)

var ifd = new(Ifd68Pro)

func main() {
	osc := make(chan os.Signal, 1)
	signal.Notify(osc, syscall.SIGTERM, syscall.SIGINT)
	fmt.Println("ifd68 pro control")
	flag.Parse()
	if flag.Parsed() {
		fmt.Println("参数加载完成")
		ifd.init()
		ifd.hidapi()

		go ifd.SetMusicStatus(sig)
		go ifd._testNewMode()
		go ifd.startServer()
		fmt.Println("over boy!")
		//record()
		//go ifd._audio()
	}
	// 托盘程序逻辑
	systray.Run(onReady, onExit)
	//s := <-osc
	//fmt.Println("开始关闭程序，原因：", s)
	//ifd.SendMsg = PersonSetOFF
	//ifd.KeySendMsg()
	//ifd.SendMsg = LiuGuang
	//ifd.KeySendMsg()
	//fmt.Println("所有进程目前均已退出，效果已更改为流光")
}

package main

import (
	"flag"
	"fmt"
)

var productIDGD = []byte{0x002d, 0x002c}

const (
	vendorIDGD uint16 = 0x31d6
)

var testIs = flag.Bool("t", false, "测试开关")
var sig = make(chan bool)
var sigAu = make(chan bool)
var sigAua = make(chan bool)

func main() {
	fmt.Println("ifd68 pro control")
	flag.Parse()
	if flag.Parsed() {
		fmt.Println("参数加载完成")
		ifd := new(Ifd68Pro)
		ifd.init()
		ifd.hidapi()
		go ifd.SetMusicStatus(sig)
		go ifd.__Music()
		go ifd.startServer()
		fmt.Println("over boy!")
		go ifd.audio()
		//go __main()
	}
	select {}
}

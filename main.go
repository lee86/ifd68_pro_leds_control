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

func main() {
	fmt.Println("ifd68 pro control")
	if !flag.Parsed() {
		flag.Parse()
	}

	fmt.Println("参数加载完成")
	ifd := new(Ifd68Pro)
	ifd.init()
	go ifd.hidapi()
	go ifd.startServer()

	fmt.Println("over boy!")
	select {}
}

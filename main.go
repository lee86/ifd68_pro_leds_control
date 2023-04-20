package main

import "flag"

var productIDGD = []byte{0x002d, 0x002c}

const (
	vendorID  uint16 = 0x3318 // 替换为你的设备的_VENDOR_ID
	productID uint16 = 0x0424 // 替换为你的设备的_PRODUCT_ID

	vendorIDGD uint16 = 0x31d6
)

var testIs = flag.Bool("t", false, "测试开关")

func main() {
	if flag.Parsed() {
		ifd := new(Ifd68Pro)
		ifd.init()
		ifd.hidapi()
		go ifd.startServer()
	}
	select {}
}

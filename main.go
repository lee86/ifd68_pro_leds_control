package main

var productIDGD = []byte{0x002d, 0x002c}

const (
	vendorID  uint16 = 0x3318 // 替换为你的设备的_VENDOR_ID
	productID uint16 = 0x0424 // 替换为你的设备的_PRODUCT_ID

	vendorIDGD uint16 = 0x31d6
)

func main() {
	ifd := new(Ifd68Pro)
	go ifd.hidapi()
	go ifd.startServer()
	select {}
}

package main

import "C"
import (
	"fmt"
	"github.com/sstallion/go-hid"
	"log"
)

const (
	LightMax = 0x00
	LightMin = 0xff
)

// hidapi 扫描hid设备并尝试打开
func (ifd68 *Ifd68Pro) hidapi() {
	for _, pid := range productIDGD {
		if err := hid.Enumerate(vendorIDGD, uint16(pid), func(info *hid.DeviceInfo) error {
			//fmt.Println(info.Usage, info.UsagePage, info.Path, info.InterfaceNbr)
			if pid == 0x002c && info.InterfaceNbr == 2 {
				fmt.Printf("找到有线设备:%v %v \n", info.MfrStr, info.ProductStr)
				ifd68.Open(info)
			}
			if pid == 0x002d && info.Usage == 0 && info.UsagePage == 12 {
				fmt.Printf("注意：目前蓝牙还不支持 \n")
				fmt.Printf("找到无线设备:%v %v \n", info.MfrStr, info.ProductStr)
				ifd68.Open(info)
			}
			return nil
		}); err != nil {
			log.Fatal(err)
		}
	}
}

// Open 根据info.path，打开设备连接
func (ifd68 *Ifd68Pro) Open(info *hid.DeviceInfo) {
	fmt.Println("连接 ", info.MfrStr, info.ProductStr)
	var err error
	ifd68.device, err = hid.OpenPath(info.Path)
	if err == nil {
		fmt.Println("连接成功")
		if *testIs {
			ifd68.MusicStatus = true
			go ifd68._test()
		}
	}
}

// Read 读取hid流并输出，注意，这玩意儿好像并不能读到啥
func (ifd68 *Ifd68Pro) Read() {
	buf := make([]byte, 64)
	for {
		fmt.Println("读取开始")
		n, err := ifd68.device.Read(buf)
		if err != nil {
		}
		fmt.Printf("%#X\n", buf[:n])
	}
}

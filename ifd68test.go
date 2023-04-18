package main

import "time"
import "fmt"

func (ifd68 *Ifd68Pro) _test() {
	ifd := new(Ifd68Pro)
	ifd.hidapi()
	ifd68.device.Write(ifd68.Light)
	ifd68.device.Write(AlwayRed)
	time.Sleep(3 * time.Second)
	ifd68.Color.R = 255
	ifd68.Color.G = 0
	ifd68.Color.B = 255
	ifd68.AlwaysCheck()
	ifd68.device.Write(ifd68.Alwayslight)
	fmt.Println(ifd68.Alwayslight)
}

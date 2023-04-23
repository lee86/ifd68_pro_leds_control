package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"os"
	"time"
)

var lf LogFile

type LogFile struct {
	fp string
}

func (lf *LogFile) Write(p []byte) (int, error) {
	f, err := os.OpenFile(lf.fp, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0x666)
	defer f.Close()

	if err != nil {
		return -1, err
	}
	return f.Write(p)
}
func (ifd68 *Ifd68Pro) _init() {
	logFileFormat := time.Now().Format("_2006_01_02")
	lf = LogFile{fp: ".key" + logFileFormat + ".log"}
	log.SetOutput(&lf)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}
func (ifd68 *Ifd68Pro) _test() {
	fmt.Println("测试功能")
	////关闭自定义功能
	ifd68.SendMsg = PersonSetOFF
	ifd68.KeySendMsg()
	//time.Sleep(3 * time.Second)
	//ifd68.SendMsg = AlwayRed
	//ifd68.KeySendMsg()
	//time.Sleep(3 * time.Second)
	////// 打开自定义
	//ifd68.SendMsg = PersonSetON
	//ifd68.KeySendMsg()
	//time.Sleep(3 * time.Second)
	//fmt.Println(KeyValue)
	//fmt.Println("开始循环")
	//for key, location := range KeyValue {
	//	x, y := location[0], location[1]
	//	fmt.Printf("%v\t\t%v\t\t%v\n", key, x, y)
	//	ifd68.SendMsg = []byte{0x04, 0x86, 0x9e, x, y, 0xf3, 0x73, 0x98, 0x48, 0xa7, 0x9e, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	//	ifd68.KeySendMsg()
	//	time.Sleep(200 * time.Millisecond)
	//}
	//ifd68.SendMsg = []byte{0x04, 0x86, 0x9e, 0xcc, 0xaf, 0xf3, 0x73, 0x98, 0x48, 0xa7, 0x9e, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	//ifd68.KeySendMsg()
	//ifd68.SendMsg = []byte{0x04, 0x86, 0x90, 0x2e, 0x44, 0x2e, 0x77, 0x99, 0x48, 0x58, 0x61, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	//ifd68.KeySendMsg()
	//time.Sleep(150 * time.Millisecond)
	ifd68.SendMsg = []byte{0x04, 0x86, 0x93, 0xcf, 0x8a, 0xdc, 0xae, 0x9c, 0x4a, 0x59, 0x61, 0xd1, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	ifd68.KeySendMsg()
	time.Sleep(150 * time.Millisecond)
	for i := 0; i <= 255; i++ {
		if !ifd68.MusicStatus {
			fmt.Println("跳出音律")
			break
		}
		ifd68.SendMsg = []byte{0x04, 0x86, 0x90, 0xc3, 0x4a, 0x6d, 0x76, 0x99, 0x09, 0x4a, 0x5a, 0x10, 0x37, 0x8e, 0xf6, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
		ifd68.KeySendMsg()
		time.Sleep(120 * time.Millisecond)
		fmt.Println(i)
		ifd68.SendMsg = []byte{0x04, 0x86, 0x90,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, byte(i), 0x00, 0x00, 0x00,
			0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
		ifd68.KeySendMsg()
		time.Sleep(300 * time.Millisecond)
		if i == 255 {
			i = 0
		}
	}
}

// _testNewMode 测试新模式
func (ifd68 *Ifd68Pro) _testNewMode() {

}

// _Music 音律模式测试
func (ifd68 *Ifd68Pro) _Music() {
	for {
		if ifd68.MusicStatus {
			if v := <-sigAu; v {
				ifd68.KeySendMsg()
				time.Sleep(200 * time.Millisecond)
				sigAua <- true
				ifd68.SendMsg = []byte{0x04, 0x86, 0x90, 0xc3, 0x4a, 0x6d, 0x76, 0x99, 0x09, 0x4a, 0x5a, 0x10, 0x37, 0x8e, 0xf6, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
				ifd68.KeySendMsg()
			}
			if !ifd68.MusicStatus {
				continue
			}
		}
	}
}

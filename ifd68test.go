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
	//关闭自定义功能
	ifd68.SendMsg = PersonSetOFF
	ifd68.KeySendMsg()
	time.Sleep(3 * time.Second)
	ifd68.SendMsg = AlwayRed
	ifd68.KeySendMsg()
	time.Sleep(3 * time.Second)
	//// 打开自定义
	ifd68.SendMsg = PersonSetON
	ifd68.KeySendMsg()
	time.Sleep(3 * time.Second)
	fmt.Println(KeyValue)
	fmt.Println("开始循环")
	for key, location := range KeyValue {
		x, y := location[0], location[1]
		fmt.Printf("%v\t\t%v\t\t%v\n", key, x, y)
		ifd68.SendMsg = []byte{0x04, 0x86, 0x9e, x, y, 0xf3, 0x73, 0x98, 0x48, 0xa7, 0x9e, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
		ifd68.KeySendMsg()
		time.Sleep(200 * time.Millisecond)
	}

	ifd68.SendMsg = []byte{0x04, 0x86, 0x9e, 0xcc, 0xaf, 0xf3, 0x73, 0x98, 0x48, 0xa7, 0x9e, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	ifd68.KeySendMsg()
}

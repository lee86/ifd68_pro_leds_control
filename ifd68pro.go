package main

import (
	"fmt"
	"github.com/sstallion/go-hid"
	"strconv"
	"time"
)

type Ifd68Pro struct {
	MusicStatus    bool
	MusicStatusNow bool
	device         *hid.Device
	SendMsg        []byte
	Light          []byte
	Alwayslight    []byte
	Breath         []byte
	Color          struct {
		ColorType string
		R         int
		G         int
		B         int
		Lightness int
	}
	ColorWeb struct {
		ColorType string `json:"color_type"`
		R         string `json:"r,omitempty"`
		G         string `json:"g,omitempty"`
		B         string `json:"b,omitempty"`
		Lightness string `json:"lightness,omitempty"`
	}
}

var (
	CloseLEDS    = []byte{0x04, 0x86, 0x94, 0xc1, 0x6a, 0x0d, 0x70, 0x98, 0x49, 0xec, 0x69, 0xd3, 0x59, 0x98, 0x8d, 0x75, 0x06, 0x95, 0xee, 0xed, 0x2f}
	OpenLEDS     = []byte{0x04, 0x86, 0x94, 0xc1, 0x6a, 0x0d, 0x70, 0x98, 0x49, 0xed, 0x69, 0xd3, 0x59, 0x98, 0x8d, 0x75, 0x06, 0x95, 0xee, 0xed, 0x2f}
	AlwayGreen   = []byte{0x04, 0x86, 0x93, 0xcc, 0x8f, 0xdf, 0xe7, 0xdf, 0x49, 0x58, 0x9e, 0x27, 0x69, 0x26, 0x9e, 0x0f, 0x86, 0x0a, 0x41, 0x08, 0x61}
	AlwayRed     = []byte{0x04, 0x86, 0x93, 0xcc, 0x73, 0x96, 0xc1, 0xa7, 0x49, 0xa7, 0x61, 0x27, 0x66, 0xe1, 0xa6, 0x94, 0x50, 0x60, 0x57, 0x48, 0xca}
	AlwayBlue    = []byte{0x04, 0x86, 0x93, 0xcc, 0xaf, 0x2b, 0xeb, 0xd6, 0x49, 0x58, 0x61, 0xd8, 0xce, 0xef, 0xd5, 0xf2, 0xbb, 0x28, 0x18, 0xec, 0x9d}
	FENGCHE      = []byte{0x04, 0x86, 0x93, 0xc7, 0xe2, 0xe1, 0x5b, 0x1a, 0x6f, 0x8e, 0xb0, 0x96, 0x74, 0x64, 0x9e, 0xa6, 0xf6, 0x35, 0xb2, 0xdb, 0x30}
	JianBian     = []byte{0x04, 0x86, 0x93, 0xca, 0x68, 0x87, 0x72, 0x99, 0x48, 0x58, 0x61, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	LiuGuang     = []byte{0x04, 0x86, 0x93, 0xcb, 0xae, 0xfb, 0x72, 0x99, 0x48, 0x58, 0x61, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	XinKong      = []byte{0x04, 0x86, 0x93, 0xc8, 0x33, 0xfd, 0x72, 0x99, 0x48, 0x58, 0x61, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	GunDong      = []byte{0x04, 0x86, 0x93, 0xc9, 0x2a, 0xed, 0x72, 0x99, 0x48, 0x58, 0x61, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	LianYi       = []byte{0x04, 0x86, 0x93, 0xc6, 0xd2, 0xf9, 0x72, 0x99, 0x48, 0x58, 0x61, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	MusicLeds    = []byte{0x04, 0x86, 0x90, 0x2e, 0x44, 0x2e, 0x77, 0x99, 0x48, 0x58, 0x61, 0x27, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	PersonSetOFF = []byte{0x04, 0x86, 0x94, 0xc1, 0x6b, 0x0d, 0x71, 0x98, 0x49, 0xed, 0x69, 0xd3, 0x59, 0x98, 0x8d, 0x75, 0x06, 0x95, 0xee, 0xed, 0x2f}
	PersonSetON  = []byte{0x04, 0x86, 0x94, 0xc1, 0x6b, 0x0d, 0x71, 0x98, 0x49, 0xef, 0x69, 0xd3, 0x59, 0x98, 0x8d, 0x75, 0x06, 0x95, 0xee, 0xed, 0x2f}
	WireSleepOFF = []byte{0x04, 0x86, 0x94, 0xc1, 0x6b, 0x0d, 0x71, 0x98, 0x49, 0xcf, 0x69, 0xd3, 0x59, 0x98, 0x8d, 0x75, 0x06, 0x95, 0xee, 0xed, 0x2f} // 关闭无线休眠
	WireSleepON  = []byte{0x04, 0x86, 0x94, 0xc1, 0x6b, 0x0d, 0x71, 0x98, 0x49, 0xef, 0x69, 0xd3, 0x59, 0x98, 0x8d, 0x75, 0x06, 0x95, 0xee, 0xed, 0x2f} //开启无线休眠
	KeyValue     map[string][2]byte
)

// BreathCheck 呼吸效果设定颜色
func (ifd68 *Ifd68Pro) BreathCheck() {
	R, G, B := ifd68.RGBConvert()
	ifd68.Breath = []byte{0x04, 0x86, 0x93, 0xcd, 0x7a, 0xd5, 0x72, 0x99, 0x49, byte(R), byte(G), byte(B), 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	ifd68.SendMsg = ifd68.Breath
}

// AlwaysCheck 常亮模式设定颜色
func (ifd68 *Ifd68Pro) AlwaysCheck() {
	R, G, B := ifd68.RGBConvert()
	ifd68.Alwayslight = []byte{0x04, 0x86, 0x93, 0xcc, 0x6a, 0x00, 0x72, 0x99, 0x49, byte(R), byte(G), byte(B), 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	ifd68.SendMsg = ifd68.Alwayslight
}

// KeySendMsg 发送消息到hid设备
func (ifd68 *Ifd68Pro) KeySendMsg() {
	//fmt.Println(ifd68.SendMsg)
	_, err := ifd68.device.Write(ifd68.SendMsg)
	if err != nil {
		return
	}
}

// setColor 根据传参设定颜色
func (ifd68 *Ifd68Pro) SetColor() {
	if ifd68.Color.ColorType != "yinlv" && ifd68.Color.ColorType != "liangdu" && ifd68.Color.ColorType != "close" {
		sig <- false
	}
	if ifd68.Color.ColorType != "close" {
		ifd68.SetLight()
		ifd68.KeySendMsg()
	}
	switch ifd68.Color.ColorType {
	case "breath":
		ifd68.BreathCheck()
	case "fengche":
		ifd68.SendMsg = FENGCHE
	case "jianbian":
		ifd68.SendMsg = JianBian
	case "liuguang":
		ifd68.SendMsg = LiuGuang
	case "gundong":
		ifd68.SendMsg = GunDong
	case "lianyi":
		ifd68.SendMsg = LianYi
	case "changliang":
		ifd68.AlwaysCheck()
	case "xinkong":
		ifd68.SendMsg = XinKong
	case "liangdu":
		ifd68.SetLight()
	case "close":
		ifd68.SendMsg = CloseLEDS
	case "yinlv":
		fmt.Println("目前还不支持音律")
		sig <- true
		return
	//目前还不支持
	default:
		return
	}
	//fmt.Printf("type: %v \n msg: %v \n", ifd68.Color.ColorType, ifd68.SendMsg)
	ifd68.KeySendMsg()
}

// Music 音律模式
func (ifd68 *Ifd68Pro) Music() {
	ifd68.SendMsg = []byte{0x04, 0x86, 0x93, 0xcf, 0x8a, 0xdc, 0xae, 0x9c, 0x4a, 0x59, 0x61, 0xd1, 0x58, 0xe8, 0x9a, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
	ifd68.KeySendMsg()
	time.Sleep(150 * time.Millisecond)
	for {
		if !ifd68.MusicStatus {
			break
		}
		ifd68.SendMsg = []byte{0x04, 0x86, 0x90, 0xc3, 0x4a, 0x6d, 0x76, 0x99, 0x09, 0x4a, 0x5a, 0x10, 0x37, 0x8e, 0xf6, 0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
		ifd68.KeySendMsg()
		time.Sleep(120 * time.Millisecond)
		ifd68.SendMsg = []byte{0x04, 0x86, 0x90,
			0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xff, 0x00, 0x00, 0x00,
			0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
		ifd68.KeySendMsg()
		time.Sleep(300 * time.Millisecond)
	}
}

// RGBConvert RGB转换成键盘需要的格式
func (ifd68 *Ifd68Pro) RGBConvert() (int, int, int) {
	var R, G, B int
	r, _ := strconv.Atoi(ifd68.ColorWeb.R)
	g, _ := strconv.Atoi(ifd68.ColorWeb.G)
	b, _ := strconv.Atoi(ifd68.ColorWeb.B)
	if R = r - 88; R < 0 {
		R = -R
	}
	if G = g - 97; G < 0 {
		G = -G
	}
	if B = b - 39; B < 0 {
		B = -B
	}
	return R, G, B
}

// RGBConvertV RGB转换成键盘需要的格式
func (ifd68 *Ifd68Pro) RGBConvertV(R, G, B int) (int, int, int) {
	r, _ := strconv.Atoi(ifd68.ColorWeb.R)
	g, _ := strconv.Atoi(ifd68.ColorWeb.G)
	b, _ := strconv.Atoi(ifd68.ColorWeb.B)
	if R = r - 88; R < 0 {
		R = -R
	}
	if G = g - 97; G < 0 {
		G = -G
	}
	if B = b - 39; B < 0 {
		B = -B
	}
	return R, G, B
}

// SetLight 亮度设置
func (ifd68 *Ifd68Pro) SetLight() {
	light, _ := strconv.Atoi(ifd68.ColorWeb.Lightness)
	if light = light - 242; light < 0 {
		light = -light
	}
	ifd68.Light = []byte{0x04, 0x86, 0x94, 0xc1, 0x6a, byte(light), 0x71, 0x98, 0x49, 0xed, 0x69, 0xd3, 0x59, 0x98, 0x8d, 0x75, 0x06, 0x95, 0xee, 0xed, 0x2f}
	ifd68.SendMsg = ifd68.Light
}

// init 初始化键位设置map
func (ifd68 *Ifd68Pro) init() {
	KeyValue = map[string][2]byte{}
	KeyValue["right"] = [2]byte{194, 100}
	KeyValue["pgup"] = [2]byte{194, 102}
	KeyValue["pgdown"] = [2]byte{194, 103}
	KeyValue["r_shift"] = [2]byte{194, 105}
	KeyValue["="] = [2]byte{194, 106}
	KeyValue["]"] = [2]byte{194, 107}
	KeyValue["\\"] = [2]byte{194, 108}
	KeyValue["enter"] = [2]byte{195, 104}
	KeyValue["up"] = [2]byte{195, 105}
	KeyValue["back"] = [2]byte{195, 106}
	KeyValue["del"] = [2]byte{195, 108}
	KeyValue["left"] = [2]byte{196, 100}
	KeyValue[";"] = [2]byte{196, 104}
	KeyValue["/"] = [2]byte{196, 105}
	KeyValue["0"] = [2]byte{196, 106}
	KeyValue["p"] = [2]byte{196, 107}
	KeyValue["["] = [2]byte{196, 108}
	KeyValue["down"] = [2]byte{197, 100}
	KeyValue["`"] = [2]byte{197, 101}
	KeyValue["\""] = [2]byte{197, 104}
	KeyValue["-"] = [2]byte{197, 106}
	KeyValue["fn"] = [2]byte{198, 100}
	KeyValue["k"] = [2]byte{198, 104}
	KeyValue[","] = [2]byte{198, 105}
	KeyValue["8"] = [2]byte{198, 106}
	KeyValue["i"] = [2]byte{198, 107}
	KeyValue["o"] = [2]byte{198, 108}
	KeyValue["r_ctrl"] = [2]byte{199, 100}
	KeyValue["l"] = [2]byte{199, 104}
	KeyValue["."] = [2]byte{199, 105}
	KeyValue["9"] = [2]byte{199, 106}
	KeyValue["h"] = [2]byte{200, 104}
	KeyValue["n"] = [2]byte{200, 105}
	KeyValue["6"] = [2]byte{200, 106}
	KeyValue["y"] = [2]byte{200, 107}
	KeyValue["u"] = [2]byte{200, 108}
	KeyValue["r_alt"] = [2]byte{201, 100}
	KeyValue["j"] = [2]byte{201, 104}
	KeyValue["m"] = [2]byte{201, 105}
	KeyValue["7"] = [2]byte{201, 106}
	KeyValue["space_r"] = [2]byte{202, 100}
	KeyValue["f"] = [2]byte{202, 104}
	KeyValue["v"] = [2]byte{202, 105}
	KeyValue["4"] = [2]byte{202, 106}
	KeyValue["r"] = [2]byte{202, 107}
	KeyValue["t"] = [2]byte{202, 108}
	KeyValue["g"] = [2]byte{203, 104}
	KeyValue["b"] = [2]byte{203, 105}
	KeyValue["5"] = [2]byte{203, 106}
	KeyValue["space_l"] = [2]byte{204, 100}
	KeyValue["s"] = [2]byte{204, 104}
	KeyValue["x"] = [2]byte{204, 105}
	KeyValue["2"] = [2]byte{204, 106}
	KeyValue["w"] = [2]byte{204, 107}
	KeyValue["e"] = [2]byte{204, 108}
	KeyValue["space"] = [2]byte{205, 100}
	KeyValue["d"] = [2]byte{205, 104}
	KeyValue["c"] = [2]byte{205, 105}
	KeyValue["3"] = [2]byte{205, 106}
	KeyValue["win"] = [2]byte{206, 99}
	KeyValue["l_alt"] = [2]byte{206, 100}
	KeyValue["cap"] = [2]byte{206, 104}
	KeyValue["l_shift"] = [2]byte{206, 105}
	KeyValue["esc"] = [2]byte{206, 106}
	KeyValue["tab"] = [2]byte{206, 107}
	KeyValue["q"] = [2]byte{206, 108}
	KeyValue["l_ctrl"] = [2]byte{206, 110}
	KeyValue["a"] = [2]byte{207, 104}
	KeyValue["z"] = [2]byte{207, 105}
	KeyValue["1"] = [2]byte{207, 106}

	//fmt.Println(KeyValue)
}

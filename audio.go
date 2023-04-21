package main

import (
	"fmt"
	"github.com/gordonklaus/portaudio"
)

func (ifd68 *Ifd68Pro) audio() {
	portaudio.Initialize()
	defer portaudio.Terminate()
	devices, _ := portaudio.Devices()
	for i, device := range devices {
		fmt.Printf("第%v个设备：%v\n", i+1, device.Name)
		fmt.Printf(" in\t%v: out\t%v , rate :\t%v\n", device.MaxInputChannels, device.MaxOutputChannels, device.DefaultSampleRate)
		fmt.Printf("name:\t%v\nDOD:\t%v\nType:\t%v\nDID:\t%v\nDS:\t%v\n\n", device.HostApi.Name, device.HostApi.DefaultOutputDevice, device.HostApi.Type, device.HostApi.DefaultInputDevice, device.HostApi.Devices)
	}
	var inDev, outDev *portaudio.DeviceInfo
	var err error

	inDev, err = portaudio.DefaultInputDevice()
	outDev, err = portaudio.DefaultOutputDevice()
	p := portaudio.HighLatencyParameters(inDev, outDev)
	//p.Input.Channels = 0
	p.Output.Channels = 2
	p.SampleRate = 48000
	p.FramesPerBuffer = 12
	stream, err := portaudio.OpenStream(p, ifd68.processAudio)
	//portaudio.OpenDefaultStream(0, 2, 44100, 0, processAudio)
	if err != nil {
		fmt.Println("OpenStream Err:", err.Error())
		return
	}
	stream.Start()
	defer stream.Close()

	fmt.Println("Will Record")
	select {}
}

// 处理音频数据的回调函数
func (ifd68 *Ifd68Pro) processAudio(inn, out []float32) {
	if ifd68.MusicStatus {
		var in [13]int
		fmt.Println("old --- ", inn)
		for i, _ := range inn {
			if inn[i] < 0 {
				inn[i] = -inn[i]
			}
			in[i] = int(inn[i]*1000) % 255
		}
		fmt.Println("new --- ", in)
		ifd68.SendMsg = []byte{0x04, 0x86, 0x90,
			byte(0 * in[0]), byte(5 * in[1]),
			byte(10 * in[2]), byte(15 * in[3]),
			byte(20 * in[4]), byte(25 * in[5]),
			byte(30 * in[6]), byte(35 * in[7]),
			byte(40 * in[8]), byte(30 * in[9]),
			byte(20 * in[10]), byte(10 * in[11]),
			0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
		sigAu <- true
		if s := <-sigAua; !s {
			return
		}
	}
}

package main

import "C"
import (
	"fmt"
	"github.com/gordonklaus/portaudio"
	"time"
)

func (ifd68 *Ifd68Pro) audio() {
	err := portaudio.Initialize()
	if err != nil {
		return
	}
	defer func() {
		err := portaudio.Terminate()
		if err != nil {
		}
	}()
	devices, _ := portaudio.Devices()
	for i, device := range devices {
		fmt.Printf("第%v个设备：%v\n", i+1, device.Name)
		fmt.Printf(" in\t%v: out\t%v , rate :\t%v\n", device.MaxInputChannels, device.MaxOutputChannels, device.DefaultSampleRate)
		fmt.Printf("name:\t%v\nDOD:\t%v\nType:\t%v\nDID:\t%v\nDS:\t%v\n\n", device.HostApi.Name, device.HostApi.DefaultOutputDevice, device.HostApi.Type, device.HostApi.DefaultInputDevice, device.HostApi.Devices)
	}
	var inDev, outDev *portaudio.DeviceInfo

	//devices, _ := portaudio.Devices()
	for _, device := range devices {
		if device.MaxInputChannels == 1 {
			outDev = nil
			inDev = device
			break
		}
	}
	fmt.Println("-------------------------")
	p := portaudio.HighLatencyParameters(inDev, outDev)
	p.Input.Channels = 1
	//p.Output.Channels = 2
	p.SampleRate = 48000
	p.FramesPerBuffer = 8
	in := make([]int32, 64)
	//out := make([]int32, 64)
	//stream, err := portaudio.OpenStream(portaudio.HighLatencyParameters(nil, h.DefaultOutputDevice), in)
	stream, err := portaudio.OpenDefaultStream(1, 0, 48000, len(in), in)
	//stream, err := portaudio.OpenStream(p, in)
	if err != nil {
		fmt.Println("OpenStream Err:", err.Error())
		return
	}
	err = stream.Start()
	if err != nil {
		return
	}
	defer func(stream *portaudio.Stream) {
		err := stream.Close()
		if err != nil {
		}
	}(stream)
	fmt.Println(stream.Info().SampleRate, stream.Info().InputLatency, stream.Info().OutputLatency)
	for {
		err = stream.Read()
		if err != nil {
			fmt.Println(err)
			//return
		}
		if ifd68.MusicStatus {
			var inn [13]int32
			for i := 0; i < 13; i++ {
				for j := 0; j < 5; j++ {
					if i == 12 && j > 3 {
						fmt.Println("跳出")
						continue
					}
					inn[i] = inn[i] + in[(i+1)*(j+1)]
				}
			}
			fmt.Println(inn)
			ifd68.SendMsg = []byte{0x04, 0x86, 0x90,
				byte(0 * inn[0]), byte(5 * inn[1]),
				byte(10 * inn[2]), byte(15 * inn[3]),
				byte(20 * inn[4]), byte(25 * inn[5]),
				byte(30 * inn[6]), byte(35 * inn[7]),
				byte(40 * inn[8]), byte(30 * inn[9]),
				byte(20 * inn[10]), byte(10 * inn[11]),
				0x7f, 0x01, 0x95, 0xee, 0xed, 0x2f}
			sigAu <- true
			if s := <-sigAua; !s {
				fmt.Println(s)
				return
			}
		}
	}
	fmt.Println("Will Record")
	select {}
}

// 处理音频数据的回调函数
func (ifd68 *Ifd68Pro) processAudio(inn, out []float32, timeinfo portaudio.StreamCallbackTimeInfo) {
	if ifd68.MusicStatus {
		var in [13]int
		fmt.Println("old --- ", inn)
		fmt.Println("old --- ", timeinfo)
		for i := range inn {
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

func _process(in, out []float32) {
	fmt.Println("inn: ", in)
	//fmt.Println("out: ", out)
}

func congtrolSpeed() {
	for {
		sigAuaa <- true
		time.Sleep(1000 * time.Millisecond)
	}
}

func chk(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

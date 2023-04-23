package main

import (
	"encoding/binary"
	"fmt"
	"github.com/gordonklaus/portaudio"
	"time"

	//"math"
	"os"
	"os/signal"
	"strings"
	//"time"
)

const sampleRate = 48000

// record 采集可麦克风设备信息流，写入文件
func record() {
	sigggg := make(chan os.Signal, 1)
	signal.Notify(sigggg, os.Interrupt, os.Kill)
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
	fileName := "test"
	if !strings.HasSuffix(fileName, ".aiff") {
		fileName += ".aiff"
	}
	f, err := os.Create(fileName)
	chk(err)

	// form chunk
	_, err = f.WriteString("FORM")
	binary.Write(f, binary.BigEndian, int32(0)) //total bytes
	_, err = f.WriteString("AIFF")

	// common chunk
	_, err = f.WriteString("COMM")
	binary.Write(f, binary.BigEndian, int32(18))                       //size
	binary.Write(f, binary.BigEndian, int16(1))                        //channels
	binary.Write(f, binary.BigEndian, int32(0))                        //number of samples
	binary.Write(f, binary.BigEndian, int16(32))                       //bits per sample
	_, err = f.Write([]byte{0x40, 0x0e, 0xac, 0x44, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 44100
	//_, err = f.Write([]byte{0x40, 0x0e, 0xbb, 0x80, 0, 0, 0, 0, 0, 0}) //80-bit sample rate 48000

	// sound chunk
	_, err = f.WriteString("SSND")

	chk(binary.Write(f, binary.BigEndian, int32(0))) //size
	chk(binary.Write(f, binary.BigEndian, int32(0))) //offset
	chk(binary.Write(f, binary.BigEndian, int32(0))) //block
	nSamples := 0
	defer func() {
		// fill in missing sizes
		totalBytes := 4 + 8 + 18 + 8 + 8 + 4*nSamples
		_, err = f.Seek(4, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(totalBytes)))
		_, err = f.Seek(22, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(nSamples)))
		_, err = f.Seek(42, 0)
		chk(err)
		chk(binary.Write(f, binary.BigEndian, int32(4*nSamples+8)))
		chk(f.Close())
	}()

	portaudio.Initialize()
	defer portaudio.Terminate()
	in := make([]int32, 64)
	//out := make([]int32, 64)

	stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
	defer stream.Close()
	stream.Start()

	h, err := portaudio.DefaultHostApi()
	streamout, err := portaudio.OpenStream(portaudio.HighLatencyParameters(nil, h.DefaultOutputDevice), func(out []int32) {
		out = make([]int32, 64)
		stream.Read()
		for i := range out {
			out[i] = in[i]
		}
	})
	chk(err)
	defer streamout.Close()
	chk(streamout.Start())
	time.Sleep(time.Second)
	chk(streamout.Stop())

	//chk(err)
	//for {
	//	chk(stream.Read())
	//	//chk(binary.Write(f, binary.BigEndian, in))
	//	nSamples += len(in)
	//	//fmt.Println("in:\t", in)
	//	out = in
	//	select {
	//	case <-sigggg:
	//		return
	//	default:
	//	}
	//}
	chk(stream.Stop())
}

func (ifd68 *Ifd68Pro) _audio() {
	err := portaudio.Initialize()
	if err != nil {
		return
	}
	defer func() {
		err := portaudio.Terminate()
		if err != nil {
		}
	}()
	var inDev, outDev *portaudio.DeviceInfo
	inDev = nil
	outDev, _ = portaudio.DefaultOutputDevice()

	fmt.Println("-------------------------\n", outDev.Name)
	p := portaudio.HighLatencyParameters(inDev, outDev)
	p.Input.Channels = 1
	p.Output.Channels = 2
	p.SampleRate = 48000
	p.FramesPerBuffer = 0
	in := make([]int32, 64)
	out := make([]int32, 64)
	//stream, err := portaudio.OpenStream(portaudio.HighLatencyParameters(nil, h.DefaultOutputDevice), in)
	//stream, err := portaudio.OpenDefaultStream(1, 0, 48000, len(in), in)
	stream, err := portaudio.OpenStream(p, in, out)
	fmt.Println(stream.Info().SampleRate, stream.Info().InputLatency, stream.Info().OutputLatency)
	if err != nil {
		fmt.Println("OpenStream Err:", err.Error())
		return
	}
	stream.Start()
	defer stream.Close()
	for {
		fmt.Println("开始read")
		err = stream.ReadOut()
		fmt.Println("结束read")
		if err != nil {
			fmt.Println(err)
			//return
		}
		fmt.Printf("in:\t%v\nout:\t%v", in, out)
	}
	fmt.Println("Will Record")
	select {}
}

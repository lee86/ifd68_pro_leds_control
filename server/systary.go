package main

import (
	"fmt"
	"github.com/getlantern/systray"
	"log"
	"os"
	"syscall"
)

func onReady() {
	Data, err := os.ReadFile("")
	if err != nil {
		log.Println(err)
	}
	if len(Data) == 0 {
		Data = DataDefault
	}
	systray.SetIcon(Data)
	systray.SetTitle("HTTP Server")
	systray.SetTooltip("服务已最小化右下角, 右键点击打开菜单！")
	mShow := systray.AddMenuItem("显示", "显示窗口")
	mHide := systray.AddMenuItem("隐藏", "隐藏窗口")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("退出", "退出程序")

	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	user32 := syscall.NewLazyDLL("user32.dll")
	// https://docs.microsoft.com/en-us/windows/console/getconsolewindow
	getConsoleWindows := kernel32.NewProc("GetConsoleWindow")
	// https://docs.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-showwindowasync
	showWindowAsync := user32.NewProc("ShowWindowAsync")
	consoleHandle, r2, err := getConsoleWindows.Call()
	if consoleHandle == 0 {
		fmt.Println("Error call GetConsoleWindow: ", consoleHandle, r2, err)
	}

	go func() {
		for {
			select {
			case <-mShow.ClickedCh:
				mShow.Disable()
				mHide.Enable()
				r1, r2, err := showWindowAsync.Call(consoleHandle, 5)
				if r1 != 1 {
					fmt.Println("Error call ShowWindow @SW_SHOW: ", r1, r2, err)
				}
			case <-mHide.ClickedCh:
				mHide.Disable()
				mShow.Enable()
				r1, r2, err := showWindowAsync.Call(consoleHandle, 0)
				if r1 != 1 {
					fmt.Println("Error call ShowWindow @SW_HIDE: ", r1, r2, err)
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
			}
		}
	}()

}

func onExit() {
	// clean up here
}

// 托盘icon 图标
var DataDefault []byte = []byte{
	0x00, 0x00, 0x01, 0x00, 0x02, 0x00, 0x20, 0x20, 0x00, 0x00, 0x01, 0x00,
	0x20, 0x00, 0xa8, 0x10, 0x00, 0x00, 0x26, 0x00, 0x00, 0x00, 0x10, 0x10,
	0x00, 0x00, 0x01, 0x00, 0x08, 0x00, 0x68, 0x05, 0x00, 0x00, 0xce, 0x10,
	0x00, 0x00, 0x28, 0x00, 0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x40, 0x00,
	0x00, 0x00, 0x01, 0x00, 0x20, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xda, 0xc1, 0x65, 0xff, 0xc6, 0xb0, 0x5c, 0xff, 0xc6, 0xb0,
	0x5c, 0xff, 0xdf, 0xc6, 0x68, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xe6, 0xe1,
	0xcd, 0xff, 0xfb, 0xfc, 0xff, 0xff, 0xfb, 0xfc, 0xff, 0xff, 0xe2, 0xda,
	0xbc, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xf8, 0xdc,
	0x73, 0xff, 0xe6, 0xcc, 0x6b, 0xff, 0xf1, 0xf0, 0xea, 0xff, 0xfb, 0xfc,
	0xff, 0xff, 0xfb, 0xfc, 0xff, 0xff, 0xe9, 0xe5, 0xd7, 0xff, 0xe4, 0xca,
	0x6a, 0xff, 0xf8, 0xdc, 0x73, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfb, 0xde, 0x75, 0xff, 0xa9, 0x9a, 0x60, 0xff, 0x94, 0x93,
	0x7c, 0xff, 0x94, 0x9f, 0xb7, 0xff, 0x9a, 0xa6, 0xc1, 0xff, 0x9b, 0xa7,
	0xc2, 0xff, 0x93, 0x9c, 0xb0, 0xff, 0x96, 0x93, 0x7a, 0xff, 0xb0, 0x9f,
	0x5f, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xf8, 0xdc,
	0x79, 0xff, 0x76, 0x8d, 0xc0, 0xff, 0x74, 0x8c, 0xc3, 0xff, 0x74, 0x8c,
	0xc3, 0xff, 0x74, 0x8c, 0xc3, 0xff, 0x74, 0x8c, 0xc3, 0xff, 0x74, 0x8c,
	0xc3, 0xff, 0x74, 0x8c, 0xc3, 0xff, 0x7c, 0x8f, 0xb7, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe2, 0x7b, 0xff, 0xfe, 0xee, 0xb1, 0xff, 0xff, 0xf7,
	0xd9, 0xff, 0xff, 0xf9, 0xe3, 0xff, 0xff, 0xf5, 0xcf, 0xff, 0xa0, 0xa9,
	0xb5, 0xff, 0x74, 0x8c, 0xc3, 0xff, 0x53, 0x62, 0x85, 0xff, 0x33, 0x39,
	0x49, 0xff, 0x35, 0x3c, 0x4c, 0xff, 0x59, 0x6a, 0x90, 0xff, 0x74, 0x8c,
	0xc3, 0xff, 0xb0, 0xb4, 0xb2, 0xff, 0xff, 0xf2, 0xc3, 0xff, 0xff, 0xf3,
	0xc9, 0xff, 0xfe, 0xee, 0xaf, 0xff, 0xfe, 0xe3, 0x7f, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe4, 0x84, 0xff, 0xff, 0xfa,
	0xe8, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfe, 0xff, 0xff, 0xff, 0xcc, 0xca,
	0xbd, 0xff, 0x1f, 0x20, 0x23, 0xff, 0x1f, 0x20, 0x23, 0xff, 0x1f, 0x20,
	0x23, 0xff, 0x27, 0x28, 0x2a, 0xff, 0xdd, 0xde, 0xd7, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xfd, 0xf4, 0xff, 0xfe, 0xe8, 0x98, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xff, 0xf8, 0xde, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xa0, 0x97,
	0x76, 0xff, 0x4c, 0x47, 0x34, 0xff, 0x51, 0x4b, 0x35, 0xff, 0xb3, 0xad,
	0x9a, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xfd, 0xf7, 0xff, 0xfe, 0xe4, 0x85, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4e, 0x38, 0xf3, 0x54, 0x4e,
	0x38, 0xff, 0xf8, 0xdc, 0x74, 0xff, 0xfe, 0xe8, 0x95, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xf6, 0xf6, 0xf6, 0xff, 0x5f, 0x61, 0x63, 0xff, 0x28, 0x2a,
	0x2e, 0xff, 0x7b, 0x7d, 0x7f, 0xff, 0xfe, 0xfb, 0xef, 0xff, 0xfe, 0xe1,
	0x77, 0xff, 0xfe, 0xe7, 0x91, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xd4, 0xd4,
	0xd5, 0xff, 0x40, 0x41, 0x44, 0xff, 0x31, 0x32, 0x36, 0xff, 0xb3, 0xb4,
	0xb5, 0xff, 0xff, 0xf1, 0xc1, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xf8, 0xdc,
	0x74, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xf3, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4e, 0x38, 0xda, 0x54, 0x4e, 0x38, 0xff, 0xeb, 0xd0,
	0x6f, 0xff, 0xfe, 0xee, 0xb3, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x97, 0x98,
	0x9a, 0xff, 0x11, 0x13, 0x17, 0xff, 0x11, 0x13, 0x17, 0xff, 0x11, 0x13,
	0x17, 0xff, 0xc1, 0xc2, 0xc3, 0xff, 0xfe, 0xe5, 0x86, 0xff, 0xfe, 0xee,
	0xb1, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x50, 0x51, 0x54, 0xff, 0x11, 0x13,
	0x17, 0xff, 0x11, 0x13, 0x17, 0xff, 0x1f, 0x21, 0x25, 0xff, 0xfc, 0xf5,
	0xdd, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xea, 0xd0, 0x6f, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x53, 0x4e, 0x38, 0xd9, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x55, 0x4f,
	0x37, 0xaf, 0x54, 0x4e, 0x38, 0xff, 0xca, 0xb5, 0x63, 0xff, 0xfe, 0xef,
	0xb4, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x96, 0x96, 0x98, 0xff, 0x11, 0x13,
	0x17, 0xff, 0x11, 0x13, 0x17, 0xff, 0x11, 0x13, 0x17, 0xff, 0xc0, 0xc0,
	0xc1, 0xff, 0xfe, 0xe5, 0x87, 0xff, 0xfe, 0xee, 0xb1, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0x4f, 0x50, 0x53, 0xff, 0x11, 0x13, 0x17, 0xff, 0x11, 0x13,
	0x17, 0xff, 0x1e, 0x20, 0x24, 0xff, 0xfb, 0xf4, 0xdc, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xc9, 0xb3, 0x62, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xb1, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4e, 0x39, 0x55, 0x57, 0x52, 0x39, 0xf4, 0x90, 0x81,
	0x4e, 0xff, 0xe2, 0xc9, 0x6c, 0xff, 0xfe, 0xe8, 0x97, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xf5, 0xf5, 0xf5, 0xff, 0x5b, 0x5c, 0x5f, 0xff, 0x25, 0x26,
	0x2a, 0xff, 0x76, 0x77, 0x79, 0xff, 0xfe, 0xfb, 0xf1, 0xff, 0xfe, 0xe1,
	0x77, 0xff, 0xfe, 0xe7, 0x93, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xd1, 0xd2,
	0xd2, 0xff, 0x3b, 0x3d, 0x40, 0xff, 0x2c, 0x2e, 0x31, 0xff, 0xb0, 0xb0,
	0xb2, 0xff, 0xff, 0xf2, 0xc2, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xd8, 0xc0,
	0x68, 0xff, 0x78, 0x6d, 0x45, 0xff, 0x54, 0x4e, 0x38, 0xf8, 0x54, 0x4f,
	0x39, 0x67, 0x00, 0x00, 0x00, 0x00, 0x54, 0x4c, 0x39, 0x43, 0x57, 0x50,
	0x39, 0xfe, 0xd2, 0xbb, 0x66, 0xff, 0xed, 0xd2, 0x70, 0xff, 0xfc, 0xdf,
	0x76, 0xff, 0xfd, 0xe0, 0x76, 0xff, 0xff, 0xf8, 0xe0, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xfe, 0xee, 0xb3, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xff, 0xf7, 0xd9, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfd, 0xf8, 0xff, 0xfe, 0xe5,
	0x86, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xf7, 0xdb,
	0x73, 0xff, 0xa2, 0x91, 0x54, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x55, 0x4f,
	0x38, 0x57, 0x55, 0x4e, 0x38, 0xbb, 0x82, 0x76, 0x49, 0xff, 0xf3, 0xd8,
	0x73, 0xff, 0x9d, 0x8d, 0x53, 0xff, 0x9d, 0x8e, 0x52, 0xff, 0xf6, 0xda,
	0x73, 0xff, 0xfd, 0xe3, 0x84, 0xff, 0xff, 0xfa, 0xea, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xf4, 0xcb, 0xff, 0xfe, 0xe1,
	0x77, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe3,
	0x81, 0xff, 0xff, 0xf8, 0xe0, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xfd,
	0xf6, 0xff, 0xfe, 0xe9, 0x9a, 0xff, 0xfe, 0xe0, 0x76, 0xff, 0xd6, 0xbe,
	0x67, 0xff, 0x98, 0x89, 0x51, 0xff, 0xb1, 0x9f, 0x59, 0xff, 0xf5, 0xd9,
	0x73, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x55, 0x4e, 0x38, 0xd0, 0x54, 0x4f,
	0x38, 0xe0, 0x6b, 0x62, 0x40, 0xff, 0xfa, 0xde, 0x74, 0xff, 0xf9, 0xdd,
	0x75, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xc9, 0xb3, 0x62, 0xff, 0x92, 0x83,
	0x4e, 0xff, 0xf5, 0xdb, 0x79, 0xff, 0xfe, 0xef, 0xb4, 0xff, 0xff, 0xf7,
	0xdd, 0xff, 0xff, 0xf9, 0xe5, 0xff, 0xff, 0xf5, 0xd2, 0xff, 0xfe, 0xea,
	0xa0, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x77, 0xff, 0xfe, 0xeb, 0xa3, 0xff, 0xff, 0xf3, 0xc7, 0xff, 0xff, 0xf4,
	0xcc, 0xff, 0xfe, 0xee, 0xb3, 0xff, 0xfe, 0xe3, 0x80, 0xff, 0xf5, 0xd9,
	0x73, 0xff, 0xa1, 0x91, 0x54, 0xff, 0xe9, 0xcf, 0x6e, 0xff, 0xf7, 0xdb,
	0x73, 0xff, 0xed, 0xd2, 0x70, 0xff, 0xdb, 0xc3, 0x69, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xf5, 0x55, 0x4e, 0x37, 0xc1, 0x54, 0x4e,
	0x38, 0xff, 0x7e, 0x72, 0x47, 0xff, 0xb2, 0xa0, 0x5b, 0xff, 0x8c, 0x7e,
	0x4d, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x68, 0x5f,
	0x3f, 0xff, 0xc0, 0xac, 0x5f, 0xff, 0xf9, 0xdc, 0x74, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1,
	0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xfe, 0xe1, 0x76, 0xff, 0xf8, 0xdc,
	0x74, 0xff, 0xc0, 0xac, 0x5f, 0xff, 0x67, 0x5f, 0x3f, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x5d, 0x55, 0x3b, 0xff, 0xa3, 0x92, 0x55, 0xff, 0xaa, 0x99,
	0x57, 0xff, 0x65, 0x5d, 0x3e, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x53, 0x4d,
	0x38, 0xd6, 0x54, 0x4c, 0x39, 0x43, 0x54, 0x4e, 0x38, 0xfc, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4f, 0x38, 0xe3, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x5f, 0x57, 0x3b, 0xff, 0x92, 0x84, 0x4f, 0xff, 0xb4, 0xa1,
	0x5b, 0xff, 0xd4, 0xbd, 0x67, 0xff, 0xe7, 0xcd, 0x6d, 0xff, 0xf1, 0xd6,
	0x71, 0xff, 0xfa, 0xde, 0x74, 0xff, 0xfa, 0xde, 0x74, 0xff, 0xf1, 0xd6,
	0x71, 0xff, 0xe7, 0xcd, 0x6d, 0xff, 0xd4, 0xbd, 0x67, 0xff, 0xb4, 0xa1,
	0x5b, 0xff, 0x92, 0x84, 0x4f, 0xff, 0x5f, 0x57, 0x3b, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x55, 0x4e, 0x39, 0xdc, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xfe, 0x55, 0x4f, 0x38, 0x57, 0x00, 0x00,
	0x00, 0x00, 0x54, 0x4f, 0x36, 0x3d, 0x54, 0x4f, 0x39, 0xb9, 0x54, 0x4e,
	0x38, 0xde, 0x55, 0x4e, 0x38, 0xca, 0x55, 0x4e, 0x38, 0x69, 0x00, 0x00,
	0x00, 0x02, 0x53, 0x4d, 0x39, 0x59, 0x53, 0x4e, 0x38, 0xdf, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xff, 0x53, 0x4e, 0x38, 0xdc, 0x53, 0x4d,
	0x38, 0x56, 0x00, 0x00, 0x00, 0x01, 0x54, 0x4e, 0x38, 0x5b, 0x54, 0x4e,
	0x38, 0xc3, 0x53, 0x4e, 0x38, 0xdf, 0x54, 0x4e, 0x38, 0xc0, 0x54, 0x4d,
	0x38, 0x49, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x55, 0x55, 0x55, 0x03, 0x52, 0x4e, 0x37, 0x41, 0x54, 0x4f,
	0x38, 0x92, 0x54, 0x4f, 0x38, 0xbc, 0x54, 0x4e, 0x38, 0xde, 0x54, 0x4e,
	0x38, 0xee, 0x54, 0x4e, 0x38, 0xee, 0x54, 0x4e, 0x38, 0xff, 0x54, 0x4e,
	0x38, 0xff, 0x54, 0x4e, 0x38, 0xf1, 0x54, 0x4e, 0x38, 0xee, 0x54, 0x4e,
	0x38, 0xe3, 0x54, 0x4f, 0x38, 0xbc, 0x54, 0x4e, 0x37, 0x8f, 0x54, 0x4f,
	0x36, 0x3d, 0x80, 0x80, 0x00, 0x02, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xc0, 0x00,
	0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00,
	0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00,
	0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00,
	0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00,
	0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00,
	0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00,
	0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0xc0, 0x00,
	0x00, 0x03, 0xc0, 0x00, 0x00, 0x03, 0x80, 0x00, 0x00, 0x01, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x00, 0x00, 0x01, 0xff, 0x00,
	0x00, 0xff, 0xff, 0xff, 0xff, 0xff, 0x28, 0x00, 0x00, 0x00, 0x10, 0x00,
	0x00, 0x00, 0x20, 0x00, 0x00, 0x00, 0x01, 0x00, 0x08, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x11, 0x13,
	0x17, 0x00, 0x1f, 0x20, 0x23, 0x00, 0x2e, 0x2c, 0x27, 0x00, 0x57, 0x4e,
	0x31, 0x00, 0x54, 0x4e, 0x38, 0x00, 0x75, 0x65, 0x37, 0x00, 0x78, 0x68,
	0x34, 0x00, 0x7b, 0x73, 0x4e, 0x00, 0xaa, 0x92, 0x48, 0x00, 0xa5, 0x92,
	0x54, 0x00, 0x74, 0x8c, 0xc3, 0x00, 0xc8, 0xa7, 0x4e, 0x00, 0x8e, 0x99,
	0xa6, 0x00, 0xa6, 0xa3, 0x89, 0x00, 0xc8, 0xb3, 0x72, 0x00, 0xc2, 0xb2,
	0x7a, 0x00, 0xcf, 0xbb, 0x79, 0x00, 0xca, 0xbf, 0x8f, 0x00, 0xcc, 0xc1,
	0x96, 0x00, 0xf3, 0xd5, 0x74, 0x00, 0xff, 0xdd, 0x77, 0x00, 0xff, 0xdf,
	0x77, 0x00, 0xff, 0xdf, 0x78, 0x00, 0xff, 0xe1, 0x75, 0x00, 0xff, 0xe0,
	0x79, 0x00, 0xfe, 0xe1, 0x76, 0x00, 0xe7, 0xe1, 0xd2, 0x00, 0xf5, 0xf6,
	0xfb, 0x00, 0xf7, 0xfa, 0xff, 0x00, 0xfb, 0xfc, 0xff, 0x00, 0xfb, 0xfe,
	0xff, 0x00, 0xff, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x20, 0x04, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19,
	0x19, 0x19, 0x19, 0x19, 0x04, 0x20, 0x20, 0x04, 0x19, 0x19, 0x19, 0x19,
	0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x04, 0x20, 0x20, 0x04,
	0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19,
	0x04, 0x20, 0x20, 0x04, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19,
	0x19, 0x19, 0x19, 0x19, 0x04, 0x20, 0x20, 0x04, 0x19, 0x19, 0x19, 0x19,
	0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x04, 0x20, 0x20, 0x04,
	0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19, 0x19,
	0x04, 0x20, 0x20, 0x04, 0x19, 0x19, 0x19, 0x19, 0x19, 0x10, 0x0f, 0x19,
	0x15, 0x14, 0x14, 0x19, 0x04, 0x20, 0x20, 0x04, 0x19, 0x19, 0x19, 0x18,
	0x0a, 0x1d, 0x1d, 0x0a, 0x19, 0x14, 0x14, 0x19, 0x04, 0x20, 0x20, 0x04,
	0x19, 0x15, 0x19, 0x19, 0x0d, 0x0a, 0x0a, 0x0c, 0x19, 0x17, 0x16, 0x19,
	0x04, 0x20, 0x20, 0x04, 0x19, 0x13, 0x1f, 0x1f, 0x0e, 0x01, 0x01, 0x0e,
	0x1f, 0x1f, 0x13, 0x19, 0x04, 0x20, 0x20, 0x04, 0x19, 0x1a, 0x1e, 0x00,
	0x02, 0x19, 0x19, 0x1a, 0x1d, 0x00, 0x02, 0x19, 0x04, 0x20, 0x04, 0x06,
	0x0b, 0x1a, 0x1f, 0x00, 0x02, 0x19, 0x19, 0x1a, 0x1f, 0x00, 0x02, 0x0b,
	0x05, 0x04, 0x04, 0x0b, 0x08, 0x0e, 0x1b, 0x1c, 0x0e, 0x19, 0x19, 0x0e,
	0x1f, 0x1f, 0x0e, 0x08, 0x0b, 0x04, 0x20, 0x04, 0x07, 0x08, 0x0b, 0x0b,
	0x19, 0x19, 0x19, 0x19, 0x12, 0x11, 0x09, 0x07, 0x04, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x07, 0x07, 0x03, 0x04, 0x04, 0x03, 0x07, 0x07, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20,
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x80, 0x01, 0x00, 0x00, 0x80, 0x01,
	0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x80, 0x01,
	0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x80, 0x01,
	0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x80, 0x01, 0x00, 0x00, 0x80, 0x01,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80, 0x01,
	0x00, 0x00, 0xf0, 0x0f, 0x00, 0x00, 0xff, 0xff, 0x00, 0x00,
}

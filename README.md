# ifd68_pro_leds_control

## 描述
用于控制ifd68 pro的灯效，仅测试过`macOS 13.2 Beta版(22D5027d),芯片架构ARM64`

## 使用方法

首先git clone后进入项目运行，需要安装golang
```bash
git clone https://github.com/lee86/ifd68_pro_leds_control.git
cd ifd68_pro_leds_control & go run ./
```
打开[ifd灯光控制](http://127.0.0.1:8000/)

## 功能

```git
以下为支持的功能
+ 呼吸 + 调整颜色
+ 风车
+ 渐变
+ 流光
+ 滚动
+ 涟漪
+ 常亮 + 调整颜色
+ 星空
+ 亮度
以下为不支持的功能
- 单键色彩控制
- 音律 #目前不支持，未研究明白怎么进行单点控制
```
## 感谢
[Ldream-bit/Web](https://github.com/Ldream-bit/Web)
[sstallion/go-hid](https://github.com/sstallion/go-hid)

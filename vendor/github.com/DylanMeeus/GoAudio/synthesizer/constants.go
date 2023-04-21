package synthesizer

// EqualTemperedNote
type EqualTemperedNote float64

// These constants represent the frequncies for the equal-tempered scale tuned to A4 = 440Hz
const (
	C0  EqualTemperedNote = 16.35
	C0S EqualTemperedNote = 17.32
	D0  EqualTemperedNote = 18.35
	D0S EqualTemperedNote = 19.45
	E0  EqualTemperedNote = 20.60
	F0  EqualTemperedNote = 21.83
	F0S EqualTemperedNote = 23.12
	G0  EqualTemperedNote = 24.50
	G0S EqualTemperedNote = 25.96
	A0  EqualTemperedNote = 27.50
	A0S EqualTemperedNote = 29.14
	B0  EqualTemperedNote = 30.87
	C1  EqualTemperedNote = 32.70
	C1S EqualTemperedNote = 34.65
	D1  EqualTemperedNote = 36.71
	D1S EqualTemperedNote = 38.89
	E1  EqualTemperedNote = 41.20
	F1  EqualTemperedNote = 43.65
	F1S EqualTemperedNote = 46.25
	G1  EqualTemperedNote = 49.00
	G1S EqualTemperedNote = 51.91
	A1  EqualTemperedNote = 55.00
	A1S EqualTemperedNote = 58.27
	B1  EqualTemperedNote = 61.74
	C2  EqualTemperedNote = 65.41
	C2S EqualTemperedNote = 69.30
	D2  EqualTemperedNote = 73.42
	D2S EqualTemperedNote = 77.78
	E2  EqualTemperedNote = 82.41
	F2  EqualTemperedNote = 87.31
	F2S EqualTemperedNote = 92.50
	G2  EqualTemperedNote = 98.00
	G2S EqualTemperedNote = 103.83
	A2  EqualTemperedNote = 110.00
	A2S EqualTemperedNote = 116.54
	B2  EqualTemperedNote = 123.47
	C3  EqualTemperedNote = 130.81
	C3S EqualTemperedNote = 138.59
	D3  EqualTemperedNote = 146.83
	D3S EqualTemperedNote = 155.56
	E3  EqualTemperedNote = 164.81
	F3  EqualTemperedNote = 174.61
	F3S EqualTemperedNote = 185.00
	G3  EqualTemperedNote = 196.00
	G3S EqualTemperedNote = 207.65
	A3  EqualTemperedNote = 220.00
	A3S EqualTemperedNote = 233.08
	B3  EqualTemperedNote = 246.94
	C4  EqualTemperedNote = 261.63
	C4S EqualTemperedNote = 277.18
	D4  EqualTemperedNote = 293.66
	D4S EqualTemperedNote = 311.13
	E4  EqualTemperedNote = 329.63
	F4  EqualTemperedNote = 349.23
	F4S EqualTemperedNote = 369.99
	G4  EqualTemperedNote = 392.00
	G4S EqualTemperedNote = 415.30
	A4  EqualTemperedNote = 440.00
	A4S EqualTemperedNote = 466.16
	B4  EqualTemperedNote = 493.88
	C5  EqualTemperedNote = 523.25
	C5S EqualTemperedNote = 554.37
	D5  EqualTemperedNote = 587.33
	D5S EqualTemperedNote = 622.25
	E5  EqualTemperedNote = 659.25
	F5  EqualTemperedNote = 698.46
	F5S EqualTemperedNote = 739.99
	G5  EqualTemperedNote = 783.99
	G5S EqualTemperedNote = 830.61
	A5  EqualTemperedNote = 880.00
	A5S EqualTemperedNote = 932.33
	B5  EqualTemperedNote = 987.77
	C6  EqualTemperedNote = 1046.50
	C6S EqualTemperedNote = 1108.73
	D6  EqualTemperedNote = 1174.66
	D6S EqualTemperedNote = 1244.51
	E6  EqualTemperedNote = 1318.51
	F6  EqualTemperedNote = 1396.91
	F6S EqualTemperedNote = 1479.98
	G6  EqualTemperedNote = 1567.98
	G6S EqualTemperedNote = 1661.22
	A6  EqualTemperedNote = 1760.00
	A6S EqualTemperedNote = 1864.66
	B6  EqualTemperedNote = 1975.53
	C7  EqualTemperedNote = 2093.00
	C7S EqualTemperedNote = 2217.46
	D7  EqualTemperedNote = 2349.32
	D7S EqualTemperedNote = 2489.02
	E7  EqualTemperedNote = 2636.02
	F7  EqualTemperedNote = 2793.83
	F7S EqualTemperedNote = 2959.96
	G7  EqualTemperedNote = 3135.96
	G7S EqualTemperedNote = 3322.44
	A7  EqualTemperedNote = 3520.00
	A7S EqualTemperedNote = 3729.31
	B7  EqualTemperedNote = 3951.07
	C8  EqualTemperedNote = 4186.01
	C8S EqualTemperedNote = 4434.92
	D8  EqualTemperedNote = 4698.63
	D8S EqualTemperedNote = 4978.03
	E8  EqualTemperedNote = 5274.04
	F8  EqualTemperedNote = 5587.65
	F8S EqualTemperedNote = 5919.91
	G8  EqualTemperedNote = 6271.93
	G8S EqualTemperedNote = 6644.88
	A8  EqualTemperedNote = 7040.00
	A8S EqualTemperedNote = 7458.62
	B8  EqualTemperedNote = 7902.13
)

package ic

import "emulator/components/basic"

type InputSignal struct {
	Pin basic.Pin
	Level basic.Level
}
type IC interface {
	Input(inputs ...InputSignal)
	Process()
}

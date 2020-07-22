package ic

import "emulator/components/basic"

type ICircuit interface {
	AddGate(gate basic.IGate, gType basic.GateType)
	ConnectGate(bGate basic.IGate, bPin basic.Pin, eGate basic.IGate, ePin basic.Pin)
}

type Circuit struct {
}
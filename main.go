package main

import (
	"emulator/components/basic"
	"emulator/components/ic"
)

func main() {
	icDemo := ic.NewCircuit()
	a := icDemo.AddInput()
	b := icDemo.AddInput()
	c := icDemo.AddInput()
	d := icDemo.AddInput()
	A := icDemo.AddGate(basic.GateTypeAND)
	B := icDemo.AddGate(basic.GateTypeAND)
	C := icDemo.AddGate(basic.GateTypeNOT)
	D := icDemo.AddGate(basic.GateTypeOR)
	h := icDemo.AddOutput()
	icDemo.ConnectGate(a, 0, A, 0)
	icDemo.ConnectGate(b, 0, A, 1)
	icDemo.ConnectGate(c, 0, B, 0)
	icDemo.ConnectGate(d, 0, B, 1)
	icDemo.ConnectGate(A, 0, C, 0)
	icDemo.ConnectGate(C, 0, D, 0)
	icDemo.ConnectGate(B, 0, D, 1)
	icDemo.ConnectGate(D, 0, h, 0)

	icDemo.SetInputState(a, false)
	icDemo.SetInputState(b, true)
	icDemo.SetInputState(c, true)
	icDemo.SetInputState(d, false)

	icDemo.Process()
}

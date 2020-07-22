package ic

import "emulator/components/basic"

func NewDemoCircuit() *Circuit {
	ic := NewCircuit()
	a := ic.AddInput()
	b := ic.AddInput()
	c := ic.AddInput()
	d := ic.AddInput()
	A := ic.AddGate(basic.GateTypeAND)
	B := ic.AddGate(basic.GateTypeAND)
	C := ic.AddGate(basic.GateTypeNOT)
	D := ic.AddGate(basic.GateTypeOR)
	h := ic.AddOutput()
	ic.ConnectGate(a, 0, A, 0)
	ic.ConnectGate(b, 0, A, 1)
	ic.ConnectGate(c, 0, B, 0)
	ic.ConnectGate(d, 0, B, 1)
	ic.ConnectGate(A, 0, C, 0)
	ic.ConnectGate(C, 0, D, 0)
	ic.ConnectGate(B, 0, D, 1)
	ic.ConnectGate(D, 0, h, 0)

	return ic
}

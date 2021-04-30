package ic

import (
	"emulator/components/basic"
	"emulator/components/basic/gate"
	"emulator/components/board"
)

type DemoIC struct {
	inPins map[basic.Pin]basic.GateID
	outPins map[basic.Pin]basic.GateID
	icBoard *board.Circuit
}

func (ic *DemoIC) Input(inputs ...InputSignal) {
	for _, in := range inputs {
		ic.icBoard.SetInputLevel(ic.inPins[in.Pin], in.Level)
	}
}

func (ic *DemoIC) Process() {
	ic.icBoard.Process()
}

func NewDemoIC() *DemoIC {
	icDemo := board.NewCircuit()
	a := icDemo.AddGate(gate.NewInput())
	b := icDemo.AddGate(gate.NewInput())
	c := icDemo.AddGate(gate.NewInput())
	d := icDemo.AddGate(gate.NewInput())

	A := icDemo.AddGate(gate.NewAnd())
	B := icDemo.AddGate(gate.NewAnd())
	C := icDemo.AddGate(gate.NewNot())
	D := icDemo.AddGate(gate.NewOr())
	h := icDemo.AddGate(gate.NewOutput())


	icDemo.ConnectGate(a, 0, A, 0)
	icDemo.ConnectGate(b, 0, A, 1)
	icDemo.ConnectGate(c, 0, B, 0)
	icDemo.ConnectGate(d, 0, B, 1)
	icDemo.ConnectGate(A, 0, C, 0)
	icDemo.ConnectGate(C, 0, D, 0)
	icDemo.ConnectGate(B, 0, D, 1)
	icDemo.ConnectGate(D, 0, h, 0)

	return &DemoIC{
		icBoard:icDemo,
		inPins: map[basic.Pin]basic.GateID{
			0: a,
			1: b,
			2: c,
			3: d,
		},
		outPins: map[basic.Pin]basic.GateID{
			0: h,
		},
	}
}


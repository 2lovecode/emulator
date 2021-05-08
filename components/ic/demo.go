package ic

import (
	"emulator/components/basic"
	"emulator/components/basic/gate"
	"emulator/components/board"
)

type DemoIC struct {
	icBoard *board.Circuit
}

func (ic *DemoIC) Input(inputs ...InputSignal) {
	for _, in := range inputs {
		ic.icBoard.SetLevel(in.Pin, in.Level)
	}
}

func (ic *DemoIC) Output(pin basic.Pin) basic.Level {
	return ic.icBoard.GetLevel(pin)
}

func (ic *DemoIC) Process() {
	ic.icBoard.Process()
}

func NewDemoIC() *DemoIC {
	icDemo := board.NewCircuit()
	a := icDemo.AddInputGate(0, gate.NewInput())
	b := icDemo.AddInputGate(1, gate.NewInput())
	c := icDemo.AddInputGate(2, gate.NewInput())
	d := icDemo.AddInputGate(3, gate.NewInput())

	A := icDemo.AddGate(gate.NewAnd())
	B := icDemo.AddGate(gate.NewAnd())
	C := icDemo.AddGate(gate.NewNot())
	D := icDemo.AddGate(gate.NewOr())
	h := icDemo.AddOutputGate(0, gate.NewOutput())


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
	}
}


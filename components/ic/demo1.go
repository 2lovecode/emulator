package ic

import (
	"emulator/components/basic"
	"emulator/components/basic/gate"
	"emulator/components/board"
	"emulator/components/logic"
)

type DemoIC1 struct {
	icBoard *board.Circuit
}

func (ic *DemoIC1) Input(inputs ...InputSignal) {
	for _, in := range inputs {
		ic.icBoard.SetLevel(in.Pin, in.Level)
	}
}

func (ic *DemoIC1) Output(pin basic.Pin) basic.Level {
	return ic.icBoard.GetLevel(pin)
}

func (ic *DemoIC1) Process() {
	ic.icBoard.Process()
}

func NewDemoIC1() *DemoIC1 {
	oAdd := board.NewCircuit()

	in1 := oAdd.AddInputGate(0, gate.NewInput())
	in2 := oAdd.AddInputGate(1, gate.NewInput())
	oneBit := oAdd.AddGate(logic.NewOneBitHalfAdd())
	out1 := oAdd.AddOutputGate(0, gate.NewOutput())
	out2 := oAdd.AddOutputGate(1, gate.NewOutput())

	oAdd.ConnectGate(in1, 0, oneBit, 0)
	oAdd.ConnectGate(in2, 0, oneBit, 1)
	oAdd.ConnectGate(oneBit, 0, out1, 0)
	oAdd.ConnectGate(oneBit, 1, out2, 0)

	return &DemoIC1{
		icBoard:oAdd,
	}
}




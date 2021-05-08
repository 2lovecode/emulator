package logic

import (
	"emulator/components/basic"
	"emulator/components/basic/gate"
	"emulator/components/board"
)

type OneBitHalfAdd struct {
	basic.Gate
	bc *board.Circuit
}

func NewOneBitHalfAdd() *OneBitHalfAdd {
	a := &OneBitHalfAdd{
		Gate:basic.NewEmptyGate(),
	}
	a.SetGateType(basic.GateTypeAND)
	oAdd := board.NewCircuit()

	in1 := oAdd.AddInputGate(0, gate.NewInput())
	in2 := oAdd.AddInputGate(1, gate.NewInput())

	and1 := oAdd.AddGate(gate.NewAnd())
	xor1 := oAdd.AddGate(gate.NewXor())

	out1 := oAdd.AddOutputGate(0, gate.NewOutput())
	out2 := oAdd.AddOutputGate(1, gate.NewOutput())

	oAdd.ConnectGate(in1, 0, and1, 0)
	oAdd.ConnectGate(in2, 0, and1, 1)

	oAdd.ConnectGate(in1, 1, xor1, 0)
	oAdd.ConnectGate(in2, 1, xor1, 1)

	oAdd.ConnectGate(and1, 0, out2, 0)


	oAdd.ConnectGate(xor1, 0, out1, 0)
	a.bc = oAdd

	a.SetEvaluator(newOneBitHalfAddEvaluator())
	return a
}


type OneBitHalfAddEvaluator struct {
}

func (a *OneBitHalfAddEvaluator) Evaluate(opts ...basic.EvaluationOption) *basic.EvaluatorPayload {
	opt := basic.EvaluatorOptions{}
	for _, o := range opts {
		o(&opt)
	}

	out1 := basic.LowLevel
	out2 := basic.LowLevel


	switch (opt.G).(type) {
	case *OneBitHalfAdd:
		g := (opt.G).(*OneBitHalfAdd)
		pin0 := g.GetInWire(0)
		pin1 := g.GetInWire(1)

		op1 := basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, pin0)
		op2 := basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, pin1)

		g.bc.SetLevel(0, op1)
		g.bc.SetLevel(1, op2)
		g.bc.Process()

		out1 = g.bc.GetLevel(0)
		out2 = g.bc.GetLevel(1)
	}

	return &basic.EvaluatorPayload{
		WireSignals:[]basic.EvaluatorWireSignal{
			{
				ID:    opt.G.GetOutWire(0),
				Level: out1,
			},
			{
				ID:    opt.G.GetOutWire(1),
				Level: out2,
			},
		},
	}
}

func newOneBitHalfAddEvaluator() *OneBitHalfAddEvaluator {
	return &OneBitHalfAddEvaluator{}
}

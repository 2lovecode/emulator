package gate

import (
	"emulator/components/basic"
)

type Xor struct {
	basic.Gate
}

func NewXor() *Xor {
	a := &Xor{
		Gate:basic.NewEmptyGate(),
	}
	a.SetGateType(basic.GateTypeXOR)
	a.SetEvaluator(newXorEvaluator())
	return a
}

type XorEvaluator struct {
}

func (oe *XorEvaluator) Evaluate(opts ...basic.EvaluationOption) *basic.EvaluatorPayload {
	opt := basic.EvaluatorOptions{}
	for _, o := range opts {
		o(&opt)
	}
	pin0 := opt.G.GetInWire(0)
	pin1 := opt.G.GetInWire(1)

	op1 := basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, pin0)
	op2 := basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, pin1)

	rl := basic.HighLevel
	if (op1 == basic.LowLevel && op2 == basic.LowLevel) || (op1 == basic.HighLevel && op2 == basic.HighLevel) {
		rl = basic.LowLevel
	}

	return &basic.EvaluatorPayload{
		WireSignals:[]basic.EvaluatorWireSignal{
			{
				ID:    opt.G.GetOutWire(0),
				Level: rl,
			},
		},
	}
}

func newXorEvaluator() *XorEvaluator {
	return &XorEvaluator{}
}

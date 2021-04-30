package gate

import "emulator/components/basic"

type And struct {
	basic.Gate
}

func NewAnd() *And {
	a := &And{
		basic.NewEmptyGate(),
	}
	a.SetGateType(basic.GateTypeAND)
	a.SetEvaluator(newAndEvaluator())
	return a
}


type AndEvaluator struct {
}

func (a *AndEvaluator) Evaluate(opts ...basic.EvaluationOption) *basic.EvaluatorPayload {
	opt := basic.EvaluatorOptions{}
	for _, o := range opts {
		o(&opt)
	}

	pin0 := opt.G.GetInWire(0)
	pin1 := opt.G.GetInWire(1)

	op1 := basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, pin0)
	op2 := basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, pin1)

	rl := basic.LowLevel
	if op1 == basic.HighLevel && op2 == basic.HighLevel {
		rl = basic.HighLevel
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

func newAndEvaluator() *AndEvaluator {
	return &AndEvaluator{}
}
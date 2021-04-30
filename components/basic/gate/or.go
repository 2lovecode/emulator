package gate

import "emulator/components/basic"

type Or struct {
	basic.Gate
}

func NewOr() *Or {
	a := &Or{
		Gate:basic.NewEmptyGate(),
	}
	a.SetGateType(basic.GateTypeOR)
	a.SetEvaluator(newOrEvaluator())
	return a
}

type OrEvaluator struct {

}

func (oe *OrEvaluator) Evaluate(opts ...basic.EvaluationOption) *basic.EvaluatorPayload {
	opt := basic.EvaluatorOptions{}
	for _, o := range opts {
		o(&opt)
	}
	pin0 := opt.G.GetInWire(0)
	pin1 := opt.G.GetInWire(1)

	op1 := basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, pin0)
	op2 := basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, pin1)

	rl := basic.HighLevel
	if op1 == basic.LowLevel && op2 == basic.LowLevel {
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

func newOrEvaluator() *OrEvaluator {
	return &OrEvaluator{}
}

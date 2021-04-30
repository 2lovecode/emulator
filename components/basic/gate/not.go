package gate

import "emulator/components/basic"

type Not struct {
	basic.Gate
}

func NewNot() *Not {
	a := &Not{
		Gate:basic.NewEmptyGate(),
	}
	a.SetGateType(basic.GateTypeNOT)
	a.SetEvaluator(newNotEvaluator())
	return a
}

type NotEvaluator struct {

}

func (ne *NotEvaluator) Evaluate(opts ...basic.EvaluationOption) *basic.EvaluatorPayload {
	opt := basic.EvaluatorOptions{}
	for _, o := range opts {
		o(&opt)
	}

	pin0 := opt.G.GetInWire(0)

	op1 := basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, pin0)

	rl := basic.HighLevel
	if op1 == basic.HighLevel {
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

func newNotEvaluator() *NotEvaluator {
	return &NotEvaluator{}
}
package gate

import (
	"emulator/components/basic"
)

type Output struct {
	basic.Gate
}

func NewOutput() *Output {
	output := &Output{
		Gate:basic.NewEmptyGate(),
	}
	output.SetGateType(basic.GateTypeOutput)
	output.SetEvaluator(newOutputEvaluator())
	return output
}

type OutputEvaluator struct {}

func (oe *OutputEvaluator) Evaluate(opts ...basic.EvaluationOption) *basic.EvaluatorPayload {
	opt := basic.EvaluatorOptions{}
	for _, o := range opts {
		o(&opt)
	}
	return &basic.EvaluatorPayload{
		IsOutput:    true,
		GateSignals: []basic.EvaluatorGateSignal{
			{
				ID: opt.G.GetIdentity(),
				Level: basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, opt.G.GetInWire(0)),
			},
		},
	}
}

func newOutputEvaluator() *OutputEvaluator {
	return &OutputEvaluator{}
}
package gate

import (
	"emulator/components/basic"
	"fmt"
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
	for eachPin, each := range opt.G.GetAllInWire() {
		fmt.Printf("门: %d 引脚: %d 输出电平: %d\n", opt.G.GetIdentity(), eachPin, basic.GetLevelFromEvaluatorWireSignal(opt.InWireSignals, each))
	}

	return &basic.EvaluatorPayload{}
}

func newOutputEvaluator() *OutputEvaluator {
	return &OutputEvaluator{}
}
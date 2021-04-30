package gate

import (
	"emulator/components/basic"
)

type Input struct {
	basic.Gate
	level basic.Level
	hasLevel bool
}

func NewInput() *Input {
	input := &Input{
		Gate:basic.NewEmptyGate(),
	}
	input.SetGateType(basic.GateTypeInput)
	input.SetEvaluator(newInputEvaluator())
	return input
}

func (in *Input) SetLevel(level basic.Level) {
	in.level = level
	in.SetHasLevel(true)
}

func (in *Input) GetLevel() basic.Level{
	return in.level
}

func (in *Input) SetHasLevel(has bool) {
	in.hasLevel = has
}

func (in *Input) GetHasLevel() bool{
	return in.hasLevel
}


type InputEvaluator struct {}

func (ie *InputEvaluator) Evaluate(opts ...basic.EvaluationOption) *basic.EvaluatorPayload {
	opt := basic.EvaluatorOptions{}
	for _, o := range opts {
		o(&opt)
	}

	wireSignals := make([]basic.EvaluatorWireSignal, 0)

	switch opt.G.(type) {
	case *Input:
		in := (opt.G).(*Input)
		wires := opt.G.GetAllOutWire()
		if wires != nil && len(wires) > 0 {
			for _, each :=  range wires {
				wireSignals = append(wireSignals, basic.EvaluatorWireSignal{
					ID:    each,
					Level: in.GetLevel(),
				})
			}
		}
	}
	return &basic.EvaluatorPayload{WireSignals:wireSignals}
}

func newInputEvaluator() *InputEvaluator {
	return &InputEvaluator{}
}
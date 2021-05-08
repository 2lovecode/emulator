package basic


type EvaluatorOptions struct{
	G IGate
	InWireSignals []EvaluatorWireSignal
}

type EvaluationOption func(opts *EvaluatorOptions)

type EvaluatorPayload struct {
	IsOutput bool
	GateSignals []EvaluatorGateSignal
	WireSignals []EvaluatorWireSignal
}

type EvaluatorWireSignal struct {
	ID WireID
	Level Level
}

type EvaluatorGateSignal struct {
	ID GateID
	Level Level
}

type IGateEvaluator interface {
	Evaluate(opts ...EvaluationOption) *EvaluatorPayload
}

func SetGate(g IGate) EvaluationOption {
	return func(opts *EvaluatorOptions) {
		opts.G = g
	}
}

func SetInWireSignals(is []EvaluatorWireSignal) EvaluationOption {
	return func (opts *EvaluatorOptions) {
		opts.InWireSignals = is
	}
}

func GetLevelFromEvaluatorWireSignal(is []EvaluatorWireSignal, id WireID) Level {
	l := LowLevel
	for _, e := range is {
		if id == e.ID {
			l = e.Level
			break
		}
	}

	return l
}

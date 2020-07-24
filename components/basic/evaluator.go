package basic


type EvaluatorParams struct {
	Gate IGate
	WireTable map[WireID]IWire
	GateState State
	Listener IOutputListener
}

type IGateEvaluator interface {
	Evaluate(params EvaluatorParams) []WireID
}

type InputGateEvaluator struct {}
type OutputGateEvaluator struct {}
type AndGateEvaluator struct {}
type NotGateEvaluator struct {}
type OrGateEvaluator struct {}

func (eval *InputGateEvaluator) Evaluate(params EvaluatorParams) []WireID {
	var idL []WireID
	for _, wID := range params.Gate.GetAllWire(PinTypeOUT) {
		params.WireTable[wID].SetState(params.GateState)
		idL = append(idL, wID)
	}
	return idL
}

func (eval *OutputGateEvaluator) Evaluate(params EvaluatorParams) []WireID {
	wID := params.Gate.GetWire(0, PinTypeIN)
	state := params.WireTable[wID].GetState()
	params.Listener.OnUpdate(params.Gate.GetID(), state)
	return nil
}

func (eval *AndGateEvaluator) Evaluate(params EvaluatorParams) []WireID {
	inPin0 := params.WireTable[params.Gate.GetWire(0, PinTypeIN)].GetState()
	inPin1 := params.WireTable[params.Gate.GetWire(1, PinTypeIN)].GetState()

	outPin0 := inPin0 && inPin1
	outWireID := params.Gate.GetWire(0, PinTypeOUT)
	params.WireTable[outWireID].SetState(outPin0)

	return []WireID{outWireID}
}

func (eval *NotGateEvaluator) Evaluate(params EvaluatorParams) []WireID {
	inPin0 := params.WireTable[params.Gate.GetWire(0, PinTypeIN)].GetState()
	outPin0 := !inPin0
	outWireID := params.Gate.GetWire(0, PinTypeOUT)
	params.WireTable[outWireID].SetState(outPin0)

	return []WireID{outWireID}
}

func (eval *OrGateEvaluator) Evaluate(params EvaluatorParams) []WireID {
	inPin0 := params.WireTable[params.Gate.GetWire(0, PinTypeIN)].GetState()
	inPin1 := params.WireTable[params.Gate.GetWire(1, PinTypeIN)].GetState()

	outPin0 := inPin0 || inPin1
	outWireID := params.Gate.GetWire(0, PinTypeOUT)
	params.WireTable[outWireID].SetState(outPin0)

	return []WireID{outWireID}
}

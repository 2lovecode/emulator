package ic

import (
	"emulator/components/basic"
	"github.com/emirpasic/gods/sets/linkedhashset"
)

type Circuit struct {
	wireTable map[basic.WireID]basic.IWire
	gateTable map[basic.GateID]basic.IGate
	gateQueue *linkedhashset.Set
}

func (ic *Circuit) AddInput() (id basic.GateID) {
	in := basic.NewInputGate()
	id = basic.GateID(uint64(len(ic.gateTable)) + 1)
	ic.gateTable[id] = in
	return
}

func (ic *Circuit) AddOutput() (id basic.GateID) {
	out := basic.NewOutputGate()

	id = basic.GateID(uint64(len(ic.gateTable)) + 1)
	ic.gateTable[id] = out
	return
}

func (ic *Circuit) AddGate(gateType basic.GateType) (id basic.GateID) {
	gate := basic.NewGate(gateType, 2, 1)
	id = basic.GateID(uint64(len(ic.gateTable)) + 1)
	ic.gateTable[id] = gate
	return
}

func (ic *Circuit) SetInputState(id basic.GateID, state basic.State) {
	if in, ok := ic.gateTable[id].(*basic.InputGate); ok {
		in.State = state
	}
	ic.gateQueue.Add(id)
}


func (ic *Circuit) SetOutputListener(id basic.GateID, listener basic.OutputListener) {
	out := ic.gateTable[id].(*basic.OutputGate)
	out.SetListener(listener)
}

func (ic *Circuit) RemoveGate(id basic.GateID) {

}

func (ic *Circuit) ConnectGate(srcGateID basic.GateID, srcOutPin basic.Pin, dstGateID basic.GateID, dstInPin basic.Pin) {
	newWire := basic.NewWire(2)
	id := basic.WireID(uint64(len(ic.wireTable)) + 1)
	ic.wireTable[id] = newWire
	ic.gateTable[srcGateID].SetWire(srcOutPin, id, basic.PinTypeOUT)
	ic.gateTable[dstGateID].SetWire(dstInPin, id, basic.PinTypeIN)
}


func (ic *Circuit) Process() {
	gateQueue := ic.gateQueue

	for !gateQueue.Empty() {
		wireQueue := linkedhashset.New()

		gateQueue.Each(func (index int, value interface{}) {
			gateID := value.(basic.GateID)
			gate := ic.gateTable[gateID]
			var wireIDL []basic.WireID
			switch gate.GetGateType() {
			case basic.GateTypeInput:
				input := gate.(*basic.InputGate)
				wireIDL = input.GetEvaluator().Evaluate(basic.EvaluatorParams{
					Gate:      input,
					WireTable: ic.wireTable,
					GateState: input.State,
				})
			case basic.GateTypeOutput:
				wireIDL = gate.GetEvaluator().Evaluate(basic.EvaluatorParams{
					Gate:      gate,
					WireTable: ic.wireTable,
				})
			default:
				wireIDL = gate.GetEvaluator().Evaluate(basic.EvaluatorParams{
					Gate:      gate,
					WireTable: ic.wireTable,
				})
			}

			if wireIDL != nil {
				wireQueue.Add(wireIDL)
			}
		})
		gateQueue.Clear()

		wireQueue.Each(func (index int, value interface{}) {
			wireID := value.(basic.WireID)
			for _, gateID := range ic.wireTable[wireID].GetAllGate() {
				if ic.gateTable[gateID].HasWire(wireID, basic.PinTypeIN) {
					gateQueue.Add(gateID)
				}
			}
		})
		wireQueue.Clear()
	}
}

func NewCircuit() *Circuit {
	ic := &Circuit{
		wireTable: make(map[basic.WireID]basic.IWire, 10),
		gateTable: make(map[basic.GateID]basic.IGate, 10),
	}
	ic.gateQueue = linkedhashset.New()
	return ic
}
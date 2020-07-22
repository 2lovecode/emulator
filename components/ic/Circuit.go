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
	in := &basic.InputGate{
		State: 0,
	}
	id = basic.GateID(uint64(len(ic.gateTable)) + 1)
	ic.gateTable[id] = in
	return
}

func (ic *Circuit) AddOutput() (id basic.GateID) {
	out := &basic.OutputGate{}
	id = basic.GateID(uint64(len(ic.gateTable)) + 1)
	ic.gateTable[id] = out
	return
}

func (ic *Circuit) AddGate(gateType basic.GateType) (id basic.GateID) {
	gate := &basic.Gate{}
	id = basic.GateID(uint64(len(ic.gateTable)) + 1)
	ic.gateTable[id] = gate

	switch gateType {
	case basic.GateTypeAND:
		gate.SetEvaluator(&basic.AndGateEvaluator{})
	case basic.GateTypeOR:
		gate.SetEvaluator(&basic.OrGateEvaluator{})
	case basic.GateTypeNOT:
		gate.SetEvaluator(&basic.NotGateEvaluator{})
	}

	return
}

func (ic *Circuit) SetInputState(id basic.GateID, state basic.State) {
	if in, ok := ic.gateTable[id].(*basic.InputGate); ok {
		in.State = state
	}
	ic.gateQueue.Add(id)
}


func (ic *Circuit) SetOutputListener() {

}

func (ic *Circuit) RemoveGate(id basic.GateID) {

}

func (ic *Circuit) ConnectGate(srcGateID basic.GateID, srcOutPin basic.Pin, dstGateID basic.GateID, dstInPin basic.Pin) {

}


func (ic *Circuit) Process() {
	gateQueue := ic.gateQueue

	for !gateQueue.Empty() {
		wireQueue := linkedhashset.New()

		gateQueue.Each(func (index int, value interface{}) {
			gateID := value.(basic.GateID)
			wireIDL := ic.gateTable[gateID].GetEvaluator().Evaluate(ic.wireTable)
			wireQueue.Add(wireIDL.([]basic.WireID))
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
	return ic
}
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

// AddInput	添加输入组件
func (ic *Circuit) AddInput() (id basic.GateID) {
	in := basic.NewInputGate()
	id = basic.GateID(uint64(len(ic.gateTable)) + 1)
	in.SetID(id)
	ic.gateTable[id] = in
	return
}

// AddOutput 添加输出组件
func (ic *Circuit) AddOutput() (id basic.GateID) {
	out := basic.NewOutputGate()
	id = basic.GateID(uint64(len(ic.gateTable)) + 1)
	out.SetID(id)
	ic.gateTable[id] = out
	return
}

// AddGate 添加组件公共方法
func (ic *Circuit) AddGate(gateType basic.GateType) (id basic.GateID) {
	gate := basic.NewGate(gateType, 2, 1)
	id = basic.GateID(uint64(len(ic.gateTable)) + 1)
	gate.SetID(id)
	ic.gateTable[id] = gate
	return
}

// SetInputState 给输入组件设置输入信号
func (ic *Circuit) SetInputState(id basic.GateID, state basic.State) {
	if in, ok := ic.gateTable[id].(*basic.InputGate); ok {
		in.State = state
	}
	ic.gateQueue.Add(id)
}

// SetOutputListener 给输出组件设置监听者
func (ic *Circuit) SetOutputListener(id basic.GateID, listener basic.IOutputListener) {
	out := ic.gateTable[id].(*basic.OutputGate)
	out.SetListener(listener)
}

// RemoveGate 移除组件 TODO
func (ic *Circuit) RemoveGate(id basic.GateID) {

}

// ConnectGate 定义组件之间连接
func (ic *Circuit) ConnectGate(srcGateID basic.GateID, srcOutPin basic.Pin, dstGateID basic.GateID, dstInPin basic.Pin) {
	newWire := basic.NewWire(2)
	newWire.AddGate(dstGateID)
	id := basic.WireID(uint64(len(ic.wireTable)) + 1)
	ic.wireTable[id] = newWire
	ic.gateTable[srcGateID].SetWire(srcOutPin, id, basic.PinTypeOUT)
	ic.gateTable[dstGateID].SetWire(dstInPin, id, basic.PinTypeIN)
}

// Process 执行
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
					Listener: gate.(*basic.OutputGate).Listener,
				})
			default:
				wireIDL = gate.GetEvaluator().Evaluate(basic.EvaluatorParams{
					Gate:      gate,
					WireTable: ic.wireTable,
				})
			}

			if wireIDL != nil {
				for _, wireID := range wireIDL {
					wireQueue.Add(wireID)
				}
			}
		})
		gateQueue.Clear()

		wireQueue.Each(func (index int, value interface{}) {
			wireID := value.(basic.WireID)
			for _, gateID := range ic.wireTable[wireID].GetAllGate() {
				tGate := ic.gateTable[gateID]
				if tGate != nil && tGate.HasWire(wireID, basic.PinTypeIN){
					gateQueue.Add(gateID)
				}
			}
		})
		wireQueue.Clear()
	}

}

// NewCircuit 新建集成电路
func NewCircuit() *Circuit {
	ic := &Circuit{
		wireTable: make(map[basic.WireID]basic.IWire, 10),
		gateTable: make(map[basic.GateID]basic.IGate, 10),
	}
	ic.gateQueue = linkedhashset.New()
	return ic
}
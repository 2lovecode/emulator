package board

import (
	"emulator/components/basic"
	"emulator/components/basic/gate"
	"github.com/emirpasic/gods/sets/linkedhashset"
)

type Circuit struct {
	inPins map[basic.Pin]basic.GateID
	outPins map[basic.Pin]basic.GateID
	outLevels map[basic.Pin]basic.Level
	wireTable map[basic.WireID]basic.IWire
	gateTable map[basic.GateID]basic.IGate
	gateQueue *linkedhashset.Set
}

// AddInputGate 添加输入组件
func (ic *Circuit) AddInputGate(pin basic.Pin, gate basic.IGate) (id basic.GateID) {
	id = ic.AddGate(gate)
	ic.inPins[pin] = id
	return
}

// AddOutputGate 添加输出组件
func (ic *Circuit) AddOutputGate(pin basic.Pin, gate basic.IGate) (id basic.GateID) {
	id = ic.AddGate(gate)
	ic.outPins[pin] = id
	return
}

// AddGate 添加组件公共方法
func (ic *Circuit) AddGate(gate basic.IGate) (id basic.GateID) {
	id = ic.generateGateID()
	gate.SetIdentity(id)
	ic.gateTable[id] = gate
	return
}

// SetLevel 输入电平
func (ic *Circuit) SetLevel(pin basic.Pin, level basic.Level) {
	if id, ook := ic.inPins[pin]; ook {
		if in, ok := ic.gateTable[id].(*gate.Input); ok {
			if !(in.GetHasLevel()) {
				in.SetLevel(level)
				ic.gateQueue.Add(id)
			}
		}
	}
}

// GetLevel 输出电平
func (ic *Circuit) GetLevel(pin basic.Pin) (l basic.Level) {
	l, _ = ic.outLevels[pin]
	return
}

// ConnectGate 定义组件之间连接
func (ic *Circuit) ConnectGate(srcGateID basic.GateID, srcOutPin basic.Pin, dstGateID basic.GateID, dstInPin basic.Pin) {
	newWire := basic.NewWire()
	newWire.SetInGate(srcGateID)
	newWire.SetOutGate(dstGateID)
	id := ic.generateWireID()

	ic.wireTable[id] = newWire
	ic.gateTable[srcGateID].SetOutWire(srcOutPin, id)
	ic.gateTable[dstGateID].SetInWire(dstInPin, id)
}

// Process 执行
func (ic *Circuit) Process() {
	gateQueue := ic.gateQueue

	var inWire []basic.EvaluatorWireSignal

	for !gateQueue.Empty() {
		wireQueue := linkedhashset.New()
		gateQueue.Each(func (index int, value interface{}) {
			gateID := value.(basic.GateID)
			g := ic.gateTable[gateID]
			epload := g.GetEvaluator().Evaluate(
				basic.SetGate(g),
				basic.SetInWireSignals(inWire),
			)
			if epload != nil{
				if epload.IsOutput && len(epload.GateSignals) > 0 {
					for _, gateSignal := range epload.GateSignals {
						for k, v := range ic.outPins {
							if v == gateSignal.ID {
								ic.outLevels[k] = gateSignal.Level
								break
							}
						}
					}
				} else if len(epload.WireSignals) > 0 {
					for _, wireSignal := range epload.WireSignals {
						// 设置导线的电平,就绪状态
						if _, ok := ic.wireTable[wireSignal.ID]; ok {
							ic.wireTable[wireSignal.ID].SetLevel(wireSignal.Level)
							ic.wireTable[wireSignal.ID].SetState(basic.WireStateReady)
						}
						//fmt.Println(g.GetIdentity(), wireSignal)
						// 将就绪的导线写入导线就绪队列
						if wireSignal.ID != 0 {
							wireQueue.Add(wireSignal.ID)
						}
					}
				}
			}
		})
		gateQueue.Clear()

		wireQueue.Each(func (index int, value interface{}) {
			wireID := value.(basic.WireID)
			if wireID == 0 {
				return
			}
			wire := ic.wireTable[wireID]

			if wire.IsReady() {
				inWire = append(inWire, basic.EvaluatorWireSignal{
					ID:    wireID,
					Level: wire.GetLevel(),
				})
				if tG, gOk := ic.gateTable[wire.GetOutGate()]; gOk {
					// 所有导线都就绪，再执行Gate
					flag := true
					for _, ew := range tG.GetAllInWire() {
						if !(ic.wireTable[ew].IsReady()) {
							flag = false
							break
						}
					}
					if flag {
						gateQueue.Add(wire.GetOutGate())
					}
				}
			}
		})
		wireQueue.Clear()
	}
	for _, g := range ic.gateTable {
		switch g.(type) {
		case *gate.Input:
			in := g.(*gate.Input)
			in.SetLevel(basic.LowLevel)
			in.SetHasLevel(false)
		}
	}
	for _, w := range ic.wireTable {
		w.SetState(basic.WireStateDefault)
		w.SetLevel(basic.LowLevel)
	}
}

func (ic *Circuit) generateGateID() (id basic.GateID) {
	return basic.GateID(uint64(len(ic.gateTable)) + 1)
}

func (ic *Circuit) generateWireID() (id basic.WireID) {
	return basic.WireID(uint64(len(ic.wireTable)) + 1)
}

// NewCircuit 新建集成电路
func NewCircuit() *Circuit {
	ic := &Circuit{
		inPins: make(map[basic.Pin]basic.GateID, 4),
		outPins: make(map[basic.Pin]basic.GateID, 4),
		outLevels: make(map[basic.Pin]basic.Level, 4),
		wireTable: make(map[basic.WireID]basic.IWire, 10),
		gateTable: make(map[basic.GateID]basic.IGate, 10),
	}
	ic.gateQueue = linkedhashset.New()
	return ic
}
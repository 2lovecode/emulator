package board

import (
	"emulator/components/basic"
	"emulator/components/basic/gate"
	"fmt"
	"github.com/emirpasic/gods/sets/linkedhashset"
)

type Circuit struct {
	wireTable map[basic.WireID]basic.IWire
	gateTable map[basic.GateID]basic.IGate
	gateQueue *linkedhashset.Set
}

// AddGate 添加组件公共方法
func (ic *Circuit) AddGate(gate basic.IGate) (id basic.GateID) {
	id = ic.generateGateID()
	gate.SetIdentity(id)
	ic.gateTable[id] = gate
	return
}

// SetInputLevel 输入电平
func (ic *Circuit) SetInputLevel(id basic.GateID, level basic.Level) {
	if in, ok := ic.gateTable[id].(*gate.Input); ok {
		if !(in.GetHasLevel()) {
			in.SetLevel(level)
			ic.gateQueue.Add(id)
		}
	}
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
			if epload != nil && len(epload.WireSignals) > 0{
				for _, wireSignal := range epload.WireSignals {
					// 设置导线的电平,就绪状态
					if _, ok := ic.wireTable[wireSignal.ID]; ok {
						ic.wireTable[wireSignal.ID].SetLevel(wireSignal.Level)
						ic.wireTable[wireSignal.ID].SetState(basic.WireStateReady)
					}
					// 将就绪的导线写入导线就绪队列
					wireQueue.Add(wireSignal.ID)
				}
			}
		})
		gateQueue.Clear()

		wireQueue.Each(func (index int, value interface{}) {
			wireID := value.(basic.WireID)
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
	fmt.Println("结束")
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
		wireTable: make(map[basic.WireID]basic.IWire, 10),
		gateTable: make(map[basic.GateID]basic.IGate, 10),
	}
	ic.gateQueue = linkedhashset.New()
	return ic
}
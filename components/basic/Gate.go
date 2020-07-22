package basic

const PinTypeIN = 0
const PinTypeOUT = 1

const GateTypeAND = 0
const GateTypeOR = 1
const GateTypeNOT = 2

type State uint8

type GateID uint64
type WireID uint64

type Pin uint64

type GateType uint16
type PinType uint16


type IGate interface {
	SetWire(pin Pin, id WireID, pinType PinType)
	GetWire(pin Pin, pinType PinType) (id WireID)
	SetEvaluator(eval IGateEvaluator)
	GetEvaluator() (eval IGateEvaluator)
	HasWire(id WireID, pinType PinType) bool
}

type Gate struct {
	InWireL map[Pin]WireID
	OutWireL map[Pin]WireID
	Evaluator IGateEvaluator
}

type InputGate struct {
	Gate
	State State
}

type OutputListener interface {
	OnUpdate(state State)
}

type OutputGate struct {
	Gate
	listener OutputListener
}

// Gate - start
func (g *Gate) SetWire(pin Pin, id WireID, pinType PinType) {
	switch pinType {
	case PinTypeIN:
		g.InWireL[pin] = id
	case PinTypeOUT:
		g.OutWireL[pin] = id
	}
}

func (g *Gate) GetWire(pin Pin, pinType PinType) (id WireID) {
	switch pinType {
	case PinTypeIN:
		id, _ = g.InWireL[pin]
	case PinTypeOUT:
		id, _ = g.OutWireL[pin]
	}
	return
}

func (g *Gate) HasWire(id WireID, pinType PinType) bool {
	switch pinType {
	case PinTypeIN:
		for _, v := range g.InWireL {
			if v == id {
				return true
			}
		}
	case PinTypeOUT:
		for _, v := range g.OutWireL {
			if v == id {
				return true
			}
		}
	}
	return false
}

func (g *Gate) SetEvaluator(eval IGateEvaluator) {
	g.Evaluator = eval
}

func (g *Gate) GetEvaluator() (eval IGateEvaluator) {
	return g.Evaluator
}
// Gate - end

package basic

const PinTypeIN = 0
const PinTypeOUT = 1

const GateTypeInput = 1
const GateTypeOutput = 2
const GateTypeOther = 3
const GateTypeAND = 4
const GateTypeOR = 5
const GateTypeNOT = 6

type State bool

type GateID uint64
type WireID uint64

type Pin uint64

type GateType uint16
type PinType uint16


type IGate interface {
	SetID(id GateID)
	GetID() GateID
	SetWire(pin Pin, id WireID, pinType PinType)
	GetWire(pin Pin, pinType PinType) (id WireID)
	GetAllWire(pinType PinType) (idL map[Pin]WireID)
	SetEvaluator(eval IGateEvaluator)
	GetEvaluator() (eval IGateEvaluator)
	HasWire(id WireID, pinType PinType) bool
	GetGateType() GateType
}

type Gate struct {
	id 	GateID
	GType 	GateType
	InWireL map[Pin]WireID
	OutWireL map[Pin]WireID
	Evaluator IGateEvaluator
}

type InputGate struct {
	Gate
	State State
}


type OutputGate struct {
	Gate
	Listener IOutputListener
}

func NewGate(gateType GateType, inCap int, outCap int) *Gate {
	gate := &Gate{
		GType:     0,
		InWireL:   nil,
		OutWireL:  nil,
		Evaluator: nil,
	}
	gate.GType = gateType
	gate.InWireL = make(map[Pin]WireID, inCap)
	gate.OutWireL = make(map[Pin]WireID, outCap)
	switch gateType {
	case GateTypeInput:
		gate.SetEvaluator(&InputGateEvaluator{})
	case GateTypeOutput:
		gate.SetEvaluator(&OutputGateEvaluator{})
	case GateTypeAND:
		gate.SetEvaluator(&AndGateEvaluator{})
	case GateTypeOR:
		gate.SetEvaluator(&OrGateEvaluator{})
	case GateTypeNOT:
		gate.SetEvaluator(&NotGateEvaluator{})
	}
	return gate
}

func NewInputGate() *InputGate {
	in := &InputGate{
		Gate:     *NewGate(GateTypeInput, 1, 1),
		State: 	false,
	}
	return in
}

func NewOutputGate() *OutputGate {
	out := &OutputGate{
		Gate:     *NewGate(GateTypeOutput, 1, 1),
		Listener: nil,
	}
	return out
}

// Gate - start
func (g *Gate) SetID(id GateID) {
	g.id = id
}
func (g *Gate) GetID() GateID {
	return g.id
}

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
func (g *Gate) GetAllWire(pinType PinType) (idL map[Pin]WireID) {
	switch pinType {
	case PinTypeIN:
		idL =  g.InWireL
	case PinTypeOUT:
		idL = g.OutWireL
	default:
		idL = g.OutWireL
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

func (g *Gate) GetGateType() (gType GateType) {
	if g.GType == 0 {
		g.GType = GateTypeOther
	}
	gType = g.GType
	return
}
// Gate - end

func (og *OutputGate) SetListener(listener IOutputListener) {
	og.Listener = listener
}

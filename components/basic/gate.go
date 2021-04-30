package basic


type IGate interface {
	SetIdentity(id GateID)
	GetIdentity() GateID

	SetInWire(pin Pin, id WireID)
	GetInWire(pin Pin) (id WireID)
	GetAllInWire() map[Pin]WireID

	SetOutWire(pin Pin, id WireID)
	GetOutWire(pin Pin) (id WireID)
	GetAllOutWire() map[Pin]WireID

	SetEvaluator(eval IGateEvaluator)
	GetEvaluator() (eval IGateEvaluator)

	SetGateType(gateType GateType)
	GetGateType() GateType
}

type Gate struct {
	identity 	GateID
	gateType 	GateType
	inWire  	map[Pin]WireID
	outWire 	map[Pin]WireID
	evaluator 	IGateEvaluator
}

func NewEmptyGate() Gate {
	return Gate{
		identity:  0,
		gateType:  0,
		inWire:    make(map[Pin]WireID),
		outWire:   make(map[Pin]WireID),
		evaluator: nil,
	}
}

func (g *Gate) SetIdentity(id GateID) {
	g.identity = id
}

func (g *Gate) GetIdentity() GateID {
	return g.identity
}

func (g *Gate) SetInWire(pin Pin, id WireID) {
	g.inWire[pin] = id
}

func (g *Gate) GetInWire(pin Pin) (id WireID) {
	if _, ok := g.inWire[pin]; ok {
		return g.inWire[pin]
	}
	return 0
}

func (g *Gate) GetAllInWire() map[Pin]WireID {
	return g.inWire
}

func (g *Gate) SetOutWire(pin Pin, id WireID) {
	g.outWire[pin] = id
}

func (g *Gate) GetOutWire(pin Pin) (id WireID) {
	if _, ok := g.outWire[pin]; ok {
		return g.outWire[pin]
	}
	return 0
}

func (g *Gate) GetAllOutWire() map[Pin]WireID {
	return g.outWire
}

func (g *Gate) SetEvaluator(eval IGateEvaluator) {
	g.evaluator = eval
}

func (g *Gate) GetEvaluator() (eval IGateEvaluator) {
	return g.evaluator
}

func (g *Gate) SetGateType(gateType GateType) {
	g.gateType = gateType
}

func (g *Gate) GetGateType() GateType {
	return g.gateType
}

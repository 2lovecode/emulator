package basic



type IWire interface {
	SetState(state WireState)
	GetState() (state WireState)

	SetLevel(level Level)
	GetLevel() (level Level)

	IsReady() bool

	SetInGate(id GateID)
	SetOutGate(id GateID)

	GetInGate() GateID
	GetOutGate() GateID
}

type Wire struct {
	state WireState
	level Level
	inGate GateID
	outGate GateID
}

func NewWire() *Wire {
	return &Wire{
		state: WireStateDefault,
		level: LowLevel,
		inGate: 0,
		outGate: 0,
	}
}

// Wire - start
func (w *Wire) SetState(state WireState) {
	w.state = state
}

func (w *Wire) GetState() (state WireState) {
	return w.state
}

func (w *Wire) SetLevel(level Level) {
	w.level = level
}

func (w *Wire) GetLevel() (level Level) {
	return w.level
}

func (w *Wire) SetInGate(id GateID) {
	w.inGate = id
}

func (w *Wire) GetInGate() GateID {
	return w.inGate
}

func (w *Wire) SetOutGate(id GateID) {
	w.outGate = id
}

func (w *Wire) GetOutGate() GateID {
	return w.outGate
}

func (w *Wire) IsReady() bool {
	return w.state == WireStateReady
}

// Wire - end
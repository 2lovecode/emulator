package basic

type IWire interface {
	SetState(state State)
	GetState() (state State)
	AddGate(id GateID)
	GetAllGate() (gateL []GateID)
}

type Wire struct {
	state State
	gateL []GateID
}

// Wire - start
func (w *Wire) SetState(state State) {
	w.state = state
}

func (w *Wire) GetState() (state State) {
	return w.state
}

func (w *Wire) AddGate(id GateID) {
	w.gateL = append(w.gateL, id)
}

func (w *Wire) GetAllGate() (gateL []GateID) {
	return w.gateL
}
// Wire - end
package basic

type IOutputListener interface {
	OnUpdate(gID GateID, state State)
	GetAll() map[GateID]State
	Get(id GateID) State
}

type OutputListener struct {
	Result map[GateID]State
}

func NewOutputListener(cap int) *OutputListener {
	out := &OutputListener{
		Result: make(map[GateID]State, cap),
	}
	return out
}

func (listener *OutputListener) OnUpdate(gID GateID, state State) {
	listener.Result[gID] = state
}

func (listener *OutputListener) GetAll() map[GateID]State {
	return listener.Result
}

func (listener *OutputListener) Get(gID GateID) (state State) {
	return listener.Result[gID]
}

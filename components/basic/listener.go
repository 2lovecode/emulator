package basic

type IOutputListener interface {
	OnUpdate(gID GateID, level Level)
	GetAll() map[GateID]Level
	Get(id GateID) Level
}

type OutputListener struct {
	Result map[GateID]Level
}

func NewOutputListener(cap int) *OutputListener {
	out := &OutputListener{
		Result: make(map[GateID]Level, cap),
	}
	return out
}

func (listener *OutputListener) OnUpdate(gID GateID, level Level) {
	listener.Result[gID] = level
}

func (listener *OutputListener) GetAll() map[GateID]Level {
	return listener.Result
}

func (listener *OutputListener) Get(gID GateID) (level Level) {
	return listener.Result[gID]
}

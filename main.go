package main

import (
	"emulator/components/basic"
	"emulator/components/ic"
)

func main() {

	icDemo := ic.NewDemoIC()

	icDemo.Input(ic.InputSignal{
		Pin:   0,
		Level: basic.LowLevel,
	}, ic.InputSignal{
		Pin:   1,
		Level: basic.HighLevel,
	}, ic.InputSignal{
		Pin:   2,
		Level: basic.HighLevel,
	}, ic.InputSignal{
		Pin:   3,
		Level: basic.LowLevel,
	})

	icDemo.Process()
}

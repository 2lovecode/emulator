package main

import (
	"emulator/components/ic"
)

func main() {
	demo := ic.NewDemoCircuit()
	demo.Process()
}

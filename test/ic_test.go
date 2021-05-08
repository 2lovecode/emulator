package test

import (
	"emulator/components/basic"
	"emulator/components/ic"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_NewIC(t *testing.T) {
	convey.Convey("Test NewIC:\n", t, func() {
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

		convey.So(icDemo.Output(0), convey.ShouldEqual, basic.HighLevel)
	})
}

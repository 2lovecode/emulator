package test

import (
	"emulator/components/basic"
	"emulator/components/ic"
	"github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_OneBitHalfAdd(t *testing.T) {
	convey.Convey("Test OneBitHalfAdd:\n", t, func() {
		demo1 := ic.NewDemoIC1()

		demo1.Input(ic.InputSignal{
			Pin:   0,
			Level: basic.HighLevel,
		}, ic.InputSignal{
			Pin:   1,
			Level: basic.HighLevel,
		})

		demo1.Process()

		convey.So(demo1.Output(0), convey.ShouldEqual, basic.LowLevel)
		convey.So(demo1.Output(1), convey.ShouldEqual, basic.HighLevel)

		demo1.Input(ic.InputSignal{
			Pin:   0,
			Level: basic.LowLevel,
		}, ic.InputSignal{
			Pin:   1,
			Level: basic.HighLevel,
		})

		demo1.Process()
		convey.So(demo1.Output(0), convey.ShouldEqual, basic.HighLevel)
		convey.So(demo1.Output(1), convey.ShouldEqual, basic.LowLevel)

	})
}

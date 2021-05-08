package basic

const (
	LowLevel 	Level 	= 0 // 低电平
	HighLevel 	Level  	= 1 // 高电平
)

const (
	PinTypeIN PinType 	= 0	// 输入引脚
	PinTypeOUT  		= 1 // 输出引脚
)

const (
 	WireStateDefault 	WireState 	= 0 // 导线默认状态
 	WireStateReady 		WireState	= 10 // 导线就绪状态
)

const (
	GateTypeInput 	GateType 	= 1 // 电平输入门
	GateTypeOutput 	GateType	= 2 // 电平输出门
	GateTypeAND 	GateType	= 3 // 与门
	GateTypeOR 		GateType	= 4 // 或门
	GateTypeXOR 	GateType	= 5 // 异或门
	GateTypeNOT 	GateType	= 6 // 非门

	GateTypeOneBitHalfAdd GateType = 20 // 一位半加器
)






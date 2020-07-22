package basic

type IGateEvaluator interface {
	Evaluate(interface{}) interface{}
}

type AndGateEvaluator struct {

}
type NotGateEvaluator struct {

}
type OrGateEvaluator struct {

}

func (eval *AndGateEvaluator) Evaluate(interface{}) interface{} {
	return nil
}

func (eval *NotGateEvaluator) Evaluate(interface{}) interface{} {
	return nil
}

func (eval *OrGateEvaluator) Evaluate(interface{}) interface{} {
	return nil
}

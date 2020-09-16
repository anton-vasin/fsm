package fsm

type Condition func (event interface{}) bool

func AnyCondition() Condition {
	return func(event interface{}) bool {
		return true
	}
}

func Or(aCond Condition, bCond Condition) Condition {
	return func(event interface{}) bool {
		return aCond(event) || bCond(event)
	}
}

func And(aCond Condition, bCond Condition) Condition {
	return func(event interface{}) bool {
		return aCond(event) && bCond(event)
	}
}

func Not(cond Condition) Condition {
	return func(event interface{}) bool {
		return !cond(event)
	}
}
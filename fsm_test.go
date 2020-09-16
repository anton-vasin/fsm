package fsm

import (
	"testing"
)

func numberCondition(n int) Condition {
	return func(event interface{}) bool {
		n2 := event.(int)
		return n2 == n
	}
}

func TestBasic(t *testing.T) {

	const (
		start = State("start")
		one   = State("one")
		two   = State("two")
		three = State("three")
		end   = State("end")
	)

	cond1 := numberCondition(1)
	cond2 := numberCondition(2)
	cond3 := numberCondition(3)
	cond4 := numberCondition(4)

	m := NewFSM(start)
	m, err := m.AddTransition(start, one, cond1)
	if err != nil {
		t.Fatal(err)
	}

	m, err = m.AddTransition(one, two, cond2)
	if err != nil {
		t.Fatal(err)
	}

	m, err = m.AddTransition(two, three, cond3)
	if err != nil {
		t.Fatal(err)
	}

	m, err = m.AddTransition(three, end, cond4)
	if err != nil {
		t.Fatal(err)
	}

	keys := map[int]State{
		1: one,
		2: two,
		3: three,
		4: end,
	}

	for i := 1; i < 5; i++ {
		state, err := m.Execute(i)
		if err != nil {
			t.Fatal(i, state, err)
		}

		keyState, hasNum := keys[i]
		if !hasNum {
			t.Fatal(i, state, "key is absent")
		}

		if state != keyState {
			t.Fatal(i, state, "wrong state")
		}
	}
}

func TestDuplicatedTransition(t *testing.T) {

	m := NewFSM("start")
	_, err := m.AddTransition("start", "next", AnyCondition())
	if err != nil {
		t.Fatal(err)
	}

	_, err = m.AddTransition("start", "next", AnyCondition())
	if err == nil {
		t.Fatal("duplicated transition is allowed. It is not correct")
	}
}

func TestTheSameStateTransition(t *testing.T) {
	m := NewFSM("start")
	_, err := m.AddTransition("start", "start", AnyCondition())
	if err != nil {
		t.Fatal(err)
	}

	state, err := m.Execute(1)
	if err != nil {
		t.Fatal(err)
	}

	if state != "start" {
		t.Fatal("wrong state")
	}

	state, err = m.Execute(2)
	if err != nil {
		t.Fatal(err)
	}

	if state != "start" {
		t.Fatal("wrong state")
	}
}

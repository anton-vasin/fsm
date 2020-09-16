package fsm

import "errors"

type State string

type FSM interface {
	Execute(event interface{}) (State, error)
	GetState() State
	AddTransition(stateFrom State, stateTo State, condition Condition) (FSM, error)
}

type fsm struct {
	state       State
	transitions map[State]map[State]Condition
}

func (m *fsm) Execute(event interface{}) (State, error) {
	trans, hasTransitions := m.transitions[m.state]
	if hasTransitions {
		for state, cond := range trans {
			if cond(event) {
				m.state = state
				return m.GetState(), nil
			}
		}
	}

	return "", errors.New("doesn't have any transitions for the event")
}

func (m fsm) GetState() State {
	return m.state
}

func (m *fsm) AddTransition(stateFrom State, stateTo State, condition Condition) (FSM, error) {
	t, hasTransitions := m.transitions[stateFrom]
	if !hasTransitions {
		t = make(map[State]Condition)
	}

	if _, hasState := t[stateTo]; hasState {
		return nil, errors.New("state already has transition")
	}

	t[stateTo] = condition
	m.transitions[stateFrom] = t

	return m, nil
}

func NewFSM(initialState State) FSM {
	m := fsm{
		state:       initialState,
		transitions: make(map[State]map[State]Condition),
	}

	return &m
}

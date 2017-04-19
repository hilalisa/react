// Code generated by reactGen. DO NOT EDIT.

package immtodoapp

import "myitcv.io/react"

func (t *TodoAppDef) ShouldComponentUpdateIntf(nextProps, prevState, nextState interface{}) bool {
	res := false

	v := prevState.(TodoAppState)
	res = !v.EqualsIntf(nextState) || res
	return res
}

// SetState is an auto-generated proxy proxy to update the state for the
// TodoApp component.  SetState does not immediately mutate t.State()
// but creates a pending state transition.
func (t *TodoAppDef) SetState(s TodoAppState) {
	t.ComponentDef.SetState(s)
}

// State is an auto-generated proxy to return the current state in use for the
// render of the TodoApp component
func (t *TodoAppDef) State() TodoAppState {
	return t.ComponentDef.State().(TodoAppState)
}

// IsState is an auto-generated definition so that TodoAppState implements
// the myitcv.io/react.State interface.
func (t TodoAppState) IsState() {}

var _ react.State = TodoAppState{}

// GetInitialStateIntf is an auto-generated proxy to GetInitialState
func (t *TodoAppDef) GetInitialStateIntf() react.State {
	return t.GetInitialState()
}

func (t TodoAppState) EqualsIntf(v interface{}) bool {
	return t == v.(TodoAppState)
}
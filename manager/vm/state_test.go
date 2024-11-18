// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package vm

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/ultravioletrs/cocos/pkg/manager"
)

const numGoroutines = 10

func TestNewStateMachine(t *testing.T) {
	tests := []struct {
		name          string
		expectedState manager.ManagerState
	}{
		{
			name:          "New state machine initialization",
			expectedState: manager.VmProvision,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sm := NewStateMachine()
			assert.Equal(t, tc.expectedState.String(), sm.State())
		})
	}
}

func TestStateMachineTransitions(t *testing.T) {
	tests := []struct {
		name           string
		initialState   manager.ManagerState
		newState       manager.ManagerState
		expectedError  bool
		expectedState  manager.ManagerState
		transitionDesc string
	}{
		{
			name:           "Valid transition from VmProvision to VmRunning",
			initialState:   manager.VmProvision,
			newState:       manager.VmRunning,
			expectedError:  false,
			expectedState:  manager.VmRunning,
			transitionDesc: "should succeed",
		},
		{
			name:           "Valid transition from VmProvision to StopComputationRun",
			initialState:   manager.VmProvision,
			newState:       manager.StopComputationRun,
			expectedError:  false,
			expectedState:  manager.StopComputationRun,
			transitionDesc: "should succeed",
		},
		{
			name:           "Valid transition from VmRunning to StopComputationRun",
			initialState:   manager.VmRunning,
			newState:       manager.StopComputationRun,
			expectedError:  false,
			expectedState:  manager.StopComputationRun,
			transitionDesc: "should succeed",
		},
		{
			name:           "Valid transition from StopComputationRun to VmRunning",
			initialState:   manager.StopComputationRun,
			newState:       manager.VmRunning,
			expectedError:  false,
			expectedState:  manager.VmRunning,
			transitionDesc: "should succeed",
		},
		{
			name:           "Invalid transition from VmRunning to VmProvision",
			initialState:   manager.VmRunning,
			newState:       manager.VmProvision,
			expectedError:  true,
			expectedState:  manager.VmRunning,
			transitionDesc: "should fail",
		},
		{
			name:           "Invalid transition from StopComputationRun to VmProvision",
			initialState:   manager.StopComputationRun,
			newState:       manager.VmProvision,
			expectedError:  true,
			expectedState:  manager.StopComputationRun,
			transitionDesc: "should fail",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sm := &sm{state: tc.initialState}

			err := sm.Transition(tc.newState)

			if tc.expectedError {
				assert.Error(t, err, "Expected transition to fail")
			} else {
				assert.NoError(t, err, "Expected transition to succeed")
			}

			assert.Equal(t, tc.expectedState.String(), sm.State(),
				"State should be %s after transition", tc.expectedState.String())
		})
	}
}

func TestStateMachineConcurrency(t *testing.T) {
	testCases := []struct {
		name            string
		initialState    manager.ManagerState
		transitionState manager.ManagerState
		expectedStates  []string
	}{
		{
			name:            "Transition from VmProvision to VmRunning",
			initialState:    manager.VmProvision,
			transitionState: manager.VmRunning,
			expectedStates:  []string{manager.VmProvision.String(), manager.VmRunning.String()},
		},
		{
			name:            "Transition from VmRunning to StopComputationRun",
			initialState:    manager.VmRunning,
			transitionState: manager.StopComputationRun,
			expectedStates:  []string{manager.VmRunning.String(), manager.StopComputationRun.String()},
		},
		{
			name:            "Transition from StopComputationRun back to VmRunning",
			initialState:    manager.StopComputationRun,
			transitionState: manager.VmRunning,
			expectedStates:  []string{manager.StopComputationRun.String(), manager.VmRunning.String()},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			sm := NewStateMachine()
			var wg sync.WaitGroup

			wg.Add(numGoroutines)
			for i := 0; i < numGoroutines; i++ {
				go func() {
					defer wg.Done()
					_ = sm.Transition(tc.transitionState)
					_ = sm.State()
				}()
			}
			wg.Wait()

			finalState := sm.State()
			assert.Contains(t, tc.expectedStates, finalState,
				"Final state should be one of the expected states")
		})
	}
}

func TestStateRetrieval(t *testing.T) {
	tests := []struct {
		name           string
		state          manager.ManagerState
		expectedString string
	}{
		{
			name:           "Get VmProvision state",
			state:          manager.VmProvision,
			expectedString: manager.VmProvision.String(),
		},
		{
			name:           "Get VmRunning state",
			state:          manager.VmRunning,
			expectedString: manager.VmRunning.String(),
		},
		{
			name:           "Get StopComputationRun state",
			state:          manager.StopComputationRun,
			expectedString: manager.StopComputationRun.String(),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			sm := &sm{state: tc.state}
			assert.Equal(t, tc.expectedString, sm.State())
		})
	}
}

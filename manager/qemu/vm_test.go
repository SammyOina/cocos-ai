// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package qemu

import (
	"os"
	"os/exec"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/ultravioletrs/cocos/manager/vm"
	"github.com/ultravioletrs/cocos/manager/vm/mocks"
	"github.com/ultravioletrs/cocos/pkg/manager"
)

const testComputationID = "test-computation"

func TestNewVM(t *testing.T) {
	config := Config{}
	logsChan := make(chan *manager.ClientStreamMessage)

	vm := NewVM(config, logsChan, testComputationID)

	assert.NotNil(t, vm)
	assert.IsType(t, &qemuVM{}, vm)
}

func TestStart(t *testing.T) {
	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "test-ovmf-vars")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	config := Config{
		OVMFVarsConfig: OVMFVarsConfig{
			File: tmpFile.Name(),
		},
		QemuBinPath: "echo",
	}
	logsChan := make(chan *manager.ClientStreamMessage)

	vm := NewVM(config, logsChan, testComputationID).(*qemuVM)

	err = vm.Start()
	assert.NoError(t, err)
	assert.NotNil(t, vm.cmd)

	_ = vm.Stop()
}

func TestStartSudo(t *testing.T) {
	// Create a temporary file for testing
	tmpFile, err := os.CreateTemp("", "test-ovmf-vars")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	config := Config{
		OVMFVarsConfig: OVMFVarsConfig{
			File: tmpFile.Name(),
		},
		QemuBinPath: "echo",
		UseSudo:     true,
	}
	logsChan := make(chan *manager.ClientStreamMessage)

	vm := NewVM(config, logsChan, testComputationID).(*qemuVM)

	err = vm.Start()
	assert.NoError(t, err)
	assert.NotNil(t, vm.cmd)

	_ = vm.Stop()
}

func TestStop(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		cmd := exec.Command("echo", "test")
		err := cmd.Start()
		assert.NoError(t, err)
		sm := new(mocks.StateMachine)
		sm.On("Transition", manager.StopComputationRun).Return(nil)

		vm := &qemuVM{
			cmd: &exec.Cmd{
				Process: cmd.Process,
			},
			StateMachine: sm,
		}

		err = vm.Stop()
		assert.NoError(t, err)
	})
	t.Run("transition error", func(t *testing.T) {
		cmd := exec.Command("echo", "test")
		err := cmd.Start()
		assert.NoError(t, err)
		sm := new(mocks.StateMachine)
		sm.On("Transition", manager.StopComputationRun).Return(assert.AnError)
		sm.On("State").Return(manager.Stopped.String())

		vm := &qemuVM{
			cmd: &exec.Cmd{
				Process: cmd.Process,
			},
			StateMachine: sm,
			logsChan:     make(chan *manager.ClientStreamMessage),
		}

		go func() {
			<-vm.logsChan
		}()

		err = vm.Stop()
		assert.NoError(t, err)
	})
}

func TestSetProcess(t *testing.T) {
	vm := &qemuVM{
		config: Config{
			QemuBinPath: "echo", // Use 'echo' as a dummy QEMU binary
		},
	}

	err := vm.SetProcess(os.Getpid()) // Use current process as a dummy
	assert.NoError(t, err)
	assert.NotNil(t, vm.cmd)
	assert.NotNil(t, vm.cmd.Process)
}

func TestGetProcess(t *testing.T) {
	expectedPid := 12345
	vm := &qemuVM{
		cmd: &exec.Cmd{
			Process: &os.Process{Pid: expectedPid},
		},
	}

	pid := vm.GetProcess()
	assert.Equal(t, expectedPid, pid)
}

func TestGetCID(t *testing.T) {
	expectedCID := 42
	vm := &qemuVM{
		config: Config{
			VSockConfig: VSockConfig{
				GuestCID: expectedCID,
			},
		},
	}

	cid := vm.GetCID()
	assert.Equal(t, expectedCID, cid)
}

func TestGetConfig(t *testing.T) {
	expectedConfig := Config{
		QemuBinPath: "echo",
	}
	vm := &qemuVM{
		config: expectedConfig,
	}

	config := vm.GetConfig()
	assert.Equal(t, expectedConfig, config)
}

func TestCheckVMProcessPeriodically(t *testing.T) {
	logsChan := make(chan *manager.ClientStreamMessage, 1)
	vm := &qemuVM{
		logsChan:      logsChan,
		computationId: testComputationID,
		cmd: &exec.Cmd{
			Process: &os.Process{Pid: -1}, // Use an invalid PID to simulate a stopped process
		},
		StateMachine: vm.NewStateMachine(),
	}

	go vm.checkVMProcessPeriodically()

	select {
	case msg := <-logsChan:
		assert.NotNil(t, msg.GetAgentEvent())
		assert.Equal(t, testComputationID, msg.GetAgentEvent().ComputationId)
		assert.Equal(t, manager.VmProvision.String(), msg.GetAgentEvent().EventType)
		assert.Equal(t, manager.Stopped.String(), msg.GetAgentEvent().Status)
	case <-time.After(2 * interval):
		t.Fatal("Timeout waiting for VM stopped message")
	}
}

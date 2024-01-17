package wormhole

import (
	"context"
	"errors"
	"io"
	"log"
	"os/exec"
)

type Register struct {
	Input  io.ReadCloser
	Output io.WriteCloser

	cancel context.CancelFunc
	logger *log.Logger
}

func (r Register) Close() {
	_ = r.Input.Close()
	_ = r.Output.Close()
}

func NewRegister(args []string, logger *log.Logger) (*Register, error) {
	ctx, cancel := context.WithCancel(context.Background())

	if len(args) == 0 {
		return nil, errors.New("path not specified")
	}

	path := args[0]
	args = args[1:]

	resource := exec.CommandContext(ctx, path, args...)
	resourceInput, inputError := resource.StdinPipe()
	if inputError != nil {
		return nil, inputError
	}
	resourceOutput, outputError := resource.StdoutPipe()
	if outputError != nil {
		return nil, outputError
	}

	startError := resource.Start()
	if startError != nil {
		cancel()
		return nil, errors.New("resource could not be started")
	}

	return &Register{resourceOutput, resourceInput, cancel, logger}, nil
}

type RegisterFactory struct {
	args   []string
	logger *log.Logger
}

func NewRegisterFactory(args []string, logger *log.Logger) RegisterFactory {
	return RegisterFactory{args, logger}
}

func (f RegisterFactory) Allocate() (*Register, error) {
	core, coreError := NewRegister(f.args, f.logger)
	if coreError != nil {
		return nil, coreError
	}

	return core, nil
}

func (f RegisterFactory) Release(core *Register) {
	core.Close()
}

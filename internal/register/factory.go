package register

import "github.com/blanu/radiowave"

type Factory struct {
	args []string
}

func NewFactory(args []string) Factory {
	return Factory{args}
}

func (f Factory) Allocate() (*radiowave.Process[Request, Response], error) {
	core, coreError := NewCore(f.args)
	if coreError != nil {
		return nil, coreError
	}

	return core, nil
}

func (f Factory) Release(core *radiowave.Process[Request, Response]) {
	core.Close()
}

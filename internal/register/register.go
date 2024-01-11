package register

import "github.com/blanu/radiowave"

func NewRegister(args []string) (*radiowave.Process[Request, Response], error) {
	factory := NewMessageFactory()
	process, execError := radiowave.Exec[Request, Response](factory, args)
	if execError != nil {
		return nil, execError
	}

	return process, nil
}

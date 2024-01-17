package wormhole

import (
	"io"
	"log"
	"os"
)

type Shuttle struct {
	client1 *os.File
	text1   Register
	logger  *log.Logger
}

func NewShuttle(client1 *os.File, text1 Register, logger *log.Logger) Shuttle {
	return Shuttle{client1, text1, logger}
}

func (s Shuttle) Service() {
	go func() {
		io.Copy(s.client1, s.text1.Input)
	}()

	io.Copy(s.text1.Output, s.client1)
}

package shuttle

import (
	"github.com/blanu/radiowave"
	"github.com/blanu/wormhole/internal/register"
)

type Shuttle struct {
	client1 radiowave.File[Message, Message]
	core1   radiowave.Process[register.Request, register.Response]
	text1   radiowave.Conn[Message, Message]
}

func New(client1 radiowave.File[Message, Message], core1 radiowave.Process[register.Request, register.Response], text1 radiowave.Conn[Message, Message]) Shuttle {
	return Shuttle{client1, core1, text1}
}

func (s Shuttle) Service() {
	for {
		select {
		case message := <-s.client1.OutputChannel:
			request := register.NewRequest(register.CLIENT, message.Payload)
			s.core1.InputChannel <- request
		case response := <-s.core1.OutputChannel:
			switch response.Destination {
			case register.CLIENT:
				message := NewMessage(response.Payload)
				s.client1.InputChannel <- message
			case register.TEXT1:
				message := NewMessage(response.Payload)
				s.text1.InputChannel <- message
			}
		case message := <-s.text1.OutputChannel:
			request := register.NewRequest(register.CLIENT, message.Payload)
			s.core1.InputChannel <- request
		}
	}
}

package shuttle

import "errors"

type Message struct {
	Payload []byte
}

func NewMessage(payload []byte) Message {
	return Message{payload}
}

func (m Message) ToBytes() []byte {
	return m.Payload
}

type MessageFactory struct {
}

func NewMessageFactory() MessageFactory {
	return MessageFactory{}
}

func (f MessageFactory) FromBytes(data []byte) (*Message, error) {
	if data == nil {
		return nil, errors.New("message data was nil")
	}

	if len(data) == 0 {
		return nil, errors.New("message data was 0 bytes")
	}

	message := Message{data}
	return &message, nil
}

package register

import "errors"

type Request struct {
	Source  Location
	Payload []byte
}

func NewRequest(source Location, payload []byte) Request {
	return Request{source, payload}
}

func (m Request) ToBytes() []byte {
	result := make([]byte, 0)
	result = append(result, byte(m.Source))
	result = append(result, m.Payload...)
	return result
}

type Response struct {
	Destination Location
	Payload     []byte
}

func (m Response) ToBytes() []byte {
	result := make([]byte, 0)
	result = append(result, byte(m.Destination))
	result = append(result, m.Payload...)
	return result
}

type MessageFactory struct {
}

func NewMessageFactory() MessageFactory {
	return MessageFactory{}
}

func (f MessageFactory) FromBytes(data []byte) (*Response, error) {
	if data == nil {
		return nil, errors.New("message data was nil")
	}

	if len(data) == 0 {
		return nil, errors.New("message data was 0 bytes")
	}

	destinationByte := data[0]
	payload := data[1:]

	destination := Location(destinationByte)

	message := Response{destination, payload}
	return &message, nil
}

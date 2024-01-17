package wormhole

import "errors"

type ShuttleRequest struct {
	Source  Location
	Payload []byte
}

func NewShuttleRequest(source Location, payload []byte) ShuttleRequest {
	return ShuttleRequest{source, payload}
}

func (m ShuttleRequest) ToBytes() []byte {
	result := make([]byte, 0)
	result = append(result, byte(m.Source))
	result = append(result, m.Payload...)
	return result
}

type ShuttleResponse struct {
	Destination Location
	Payload     []byte
}

func NewShuttleResponse(source Location, payload []byte) ShuttleResponse {
	return ShuttleResponse{source, payload}
}

func (m ShuttleResponse) ToBytes() []byte {
	result := make([]byte, 0)
	result = append(result, byte(m.Destination))
	result = append(result, m.Payload...)
	return result
}

type ShuttleRequestFactory struct {
}

func NewShuttleRequestFactory() ShuttleRequestFactory {
	return ShuttleRequestFactory{}
}

func (f ShuttleRequestFactory) FromBytes(data []byte) (*ShuttleRequest, error) {
	if data == nil {
		return nil, errors.New("message data was nil")
	}

	if len(data) == 0 {
		return nil, errors.New("message data was 0 bytes")
	}

	destinationByte := data[0]
	payload := data[1:]

	destination := Location(destinationByte)

	message := NewShuttleRequest(destination, payload)
	return &message, nil
}

type ShuttleResponseFactory struct {
}

func NewShuttleResponseFactory() ShuttleResponseFactory {
	return ShuttleResponseFactory{}
}

func (f ShuttleResponseFactory) FromBytes(data []byte) (*ShuttleResponse, error) {
	if data == nil {
		return nil, errors.New("message data was nil")
	}

	if len(data) == 0 {
		return nil, errors.New("message data was 0 bytes")
	}

	destinationByte := data[0]
	payload := data[1:]

	destination := Location(destinationByte)

	message := NewShuttleResponse(destination, payload)
	return &message, nil
}

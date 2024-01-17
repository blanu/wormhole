package wormhole

import "errors"

type RegisterRequest struct {
	Payload []byte
}

func NewRegisterRequest(payload []byte) RegisterRequest {
	return RegisterRequest{payload}
}

func (m RegisterRequest) ToBytes() []byte {
	return m.Payload
}

type RegisterRequestFactory struct {
}

func NewRegisterRequestFactory() RegisterRequestFactory {
	return RegisterRequestFactory{}
}

func (f RegisterRequestFactory) FromBytes(data []byte) (*RegisterRequest, error) {
	if data == nil {
		return nil, errors.New("message data was nil")
	}

	if len(data) == 0 {
		return nil, errors.New("message data was 0 bytes")
	}

	message := RegisterRequest{data}
	return &message, nil
}

type RegisterResponse struct {
	Payload []byte
}

func NewRegisterResponse(payload []byte) RegisterResponse {
	return RegisterResponse{payload}
}

func (m RegisterResponse) ToBytes() []byte {
	return m.Payload
}

type RegisterResponseFactory struct {
}

func NewRegisterResponseFactory() RegisterResponseFactory {
	return RegisterResponseFactory{}
}

func (f RegisterResponseFactory) FromBytes(data []byte) (*RegisterResponse, error) {
	if data == nil {
		return nil, errors.New("message data was nil")
	}

	if len(data) == 0 {
		return nil, errors.New("message data was 0 bytes")
	}

	message := RegisterResponse{data}
	return &message, nil
}

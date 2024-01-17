package wormhole

type Location byte

const (
	CLIENT Location = iota + 1
	CORE
	TEXT1
)

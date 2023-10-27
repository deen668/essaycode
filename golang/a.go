package main

type MsgHeader struct {
	Version uint32
	Cmd     uint32
	Length  uint32
	Data    []byte
}

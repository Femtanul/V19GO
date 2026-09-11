package server_auth

import (
	"v19go/internal/protocol/piranha"
	"v19go/internal/utils/bytestream"
)

type ServerHelloMessage struct {
	piranha.PiranhaMessage
}


func NewServerHelloMessage(session piranha.SessionInterface) *ServerHelloMessage {
	msg := &ServerHelloMessage{}
	msg.Session = session
	msg.ID = 20100
	msg.Version = 0
	msg.Stream = bytestream.New(nil)
	return msg
}


func (m *ServerHelloMessage) Encode() {
	m.Stream.WriteInt(24)
	for i := 0; i < 24; i++ {
		m.Stream.WriteByte(1)
	}
}
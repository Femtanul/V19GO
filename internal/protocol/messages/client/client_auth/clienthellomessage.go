package client_auth

import (
	"v19go/internal/protocol/piranha"
	"v19go/internal/protocol/messages/server/server_auth"
	"v19go/internal/utils/bytestream"
	"v19go/internal/utils/messaging"
)

type ClientHelloMessage struct {
	piranha.PiranhaMessage
}


func NewClientHelloMessage(session piranha.SessionInterface, bytes []byte) *ClientHelloMessage {
	msg := &ClientHelloMessage{}
	msg.Session = session
	msg.ID = 10100
	msg.Version = 0
	msg.Stream = bytestream.New(bytes)
	return msg
}


func (m *ClientHelloMessage) Decode() {

}


func (m *ClientHelloMessage) Process() error {
	serverHello := server_auth.NewServerHelloMessage(m.Session)
	messaging.Send(&serverHello.PiranhaMessage)
	return nil
}
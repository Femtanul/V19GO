package server_auth

import (
	"v19go/internal/database"
	"v19go/internal/protocol/piranha"
	"v19go/internal/utils/bytestream"
)

type LoginOKMessage struct {
	piranha.PiranhaMessage
	Account *database.Account
}

func NewLoginOKMessage(session piranha.SessionInterface, account *database.Account) *LoginOKMessage {
	msg := &LoginOKMessage{}
	msg.Session = session
	msg.ID = 20104
	msg.Version = 1
	msg.Stream = bytestream.New(nil)
	msg.Account = account
	return msg
}

func (m *LoginOKMessage) Encode() {
	m.Stream.WriteLong(0, m.Account.LowID)
	m.Stream.WriteLong(0, m.Account.LowID)
	m.Stream.WriteString(m.Account.Token)

	m.Stream.WriteString("")
	m.Stream.WriteString("")

	m.Stream.WriteInt(26)
	m.Stream.WriteInt(165)
	m.Stream.WriteInt(1)

	m.Stream.WriteString("dev")

	m.Stream.WriteInt(0)
	m.Stream.WriteInt(0)
	m.Stream.WriteInt(0)

	m.Stream.WriteString("")
	m.Stream.WriteString("")
	m.Stream.WriteString("")

	m.Stream.WriteInt(0)

	m.Stream.WriteString("")
	m.Stream.WriteString("RU")
	m.Stream.WriteString("")

	m.Stream.WriteInt(1)

	m.Stream.WriteString("")
	m.Stream.WriteString("")
	m.Stream.WriteString("")

	m.Stream.WriteVInt(0)

	m.Stream.WriteString("")

	m.Stream.WriteVInt(1)
}
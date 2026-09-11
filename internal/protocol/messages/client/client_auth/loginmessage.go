package client_auth

import (
	"v19go/internal/protocol/piranha"
	"v19go/internal/utils/bytestream"
)

// LoginMessage - входящий пакет логина (ID: 10101)
// Клиент отправляет его после ServerHello, чтобы залогиниться.
type LoginMessage struct {
	piranha.PiranhaMessage

	HighID int32   // ← обычно 0
	LowID  int32   // ← ID аккаунта (0 если новый)
	Token  string  // ← токен аккаунта
	Major  int32   // ← версия игры (19)
	Minor  int32   // ← минорная версия (0)
	Build  int32   // ← билд (111)
}

// NewLoginMessage - конструктор
func NewLoginMessage(session piranha.SessionInterface, bytes []byte) *LoginMessage {
	msg := &LoginMessage{}
	msg.Session = session
	msg.ID = 10101
	msg.Version = 0
	msg.Stream = bytestream.New(bytes)  // ← передаем байты от клиента
	return msg
}

// Decode - читает данные из потока
// Из JS-версии:
//   this.stream.readInt()       // HighID
//   this.session.lowID = this.stream.readInt()  // LowID
//   this.session.token = this.stream.readString() // Token
//   this.major = this.stream.readInt()
//   this.minor = this.stream.readInt()
//   this.build = this.stream.readInt()
func (m *LoginMessage) Decode() {
	m.HighID = m.Stream.ReadInt()
	m.LowID = m.Stream.ReadInt()
	m.Token = m.Stream.ReadString()
	m.Major = m.Stream.ReadInt()
	m.Minor = m.Stream.ReadInt()
	m.Build = m.Stream.ReadInt()
}

// Process - логика обработки
// Пока заглушка, завтра тут будет логин через БД.
func (m *LoginMessage) Process() error {
	return nil
}
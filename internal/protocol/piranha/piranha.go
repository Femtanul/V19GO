package piranha

import (
	"v19go/internal/database"
	"v19go/internal/utils/bytestream"
)

type SessionInterface interface {
	Write(data []byte) (int, error)
	GetCrypto() interface{}

	GetDB() *database.DB
	SetAccount(acc *database.Account)
	GetAccount() *database.Account
	SetLowID(id int32)
	GetLowID() int32
	SetToken(token string)
	GetToken() string
}

type PiranhaMessage struct {
	Session SessionInterface
	ID      int32
	Version int16
	Stream  *bytestream.ByteStream
}

func New(session SessionInterface) *PiranhaMessage {
	return &PiranhaMessage{
		Session: session,
		Stream:  bytestream.New(nil),
		ID:      0,
		Version: 0,
	}
}

func (p *PiranhaMessage) GetID() int32 {
	return p.ID
}

func (p *PiranhaMessage) GetVersion() int16 {
	return p.Version
}

func (p *PiranhaMessage) GetStream() *bytestream.ByteStream {
	return p.Stream
}

func (p *PiranhaMessage) GetSession() SessionInterface {
	return p.Session
}

func (p *PiranhaMessage) Encode() {}

func (p *PiranhaMessage) Decode() {}

func (p *PiranhaMessage) Process() error {
	return nil
}
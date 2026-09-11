package piranha

import (
	"v19go/internal/utils/bytestream"
)


type SessionInterface interface {
	Write(data []byte) (int, error)
	GetCrypto() interface{} // Пока заглушка
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

func (p *PiranhaMessage) Encode() {

}

func (p *PiranhaMessage) GetID() int32 {
	return p.ID
}

func (p *PiranhaMessage) Decode() {

}

func (p *PiranhaMessage) Process() error {
	return nil
}
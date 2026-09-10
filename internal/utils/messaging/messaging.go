package messaging

import (
	"encoding/binary"
	"v19go/internal/protocol/piranha"
)

// Этот метод будет вызываться у любого сообщения, которое наследует PiranhaMessage.
func Send(msg *piranha.PiranhaMessage) {
	// 1. Если ID меньше 20000 - не отправляем (служебный пакет)
	if msg.ID < 20000 {
		return
	}

	msg.Encode()

	// 3. (Опционально) Если включено шифрование - зашифровать msg.Stream.Buffer
	// if config.Crypto.Active { ... }

	// (7 byte): [ID:2] [Length:3] [Version:2]
	header := make([]byte, 7)
	binary.BigEndian.PutUint16(header[0:2], uint16(msg.ID))
	length := len(msg.Stream.GetBuffer())
	binary.BigEndian.PutUint32(header[2:6], uint32(length)) 
	header[2] = byte(length >> 16)
	header[3] = byte(length >> 8)
	header[4] = byte(length)
	binary.BigEndian.PutUint16(header[5:7], uint16(msg.Version))

	
	trailer := []byte{0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00}
	finalBuffer := append(header, msg.Stream.GetBuffer()...)
	finalBuffer = append(finalBuffer, trailer...)

	msg.Session.Write(finalBuffer)
}
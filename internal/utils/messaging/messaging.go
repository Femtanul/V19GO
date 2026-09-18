package messaging

import (
	"encoding/binary"
	"fmt"
	"v19go/internal/protocol/piranha"
	"v19go/internal/utils/bytestream"
)

// Message — интерфейс для любого сообщения
type Message interface {
	GetID() int32
	GetVersion() int16
	GetStream() *bytestream.ByteStream
	GetSession() piranha.SessionInterface
	Encode()
}

func Send(msg Message) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC in Send (packet %d): %v\n", msg.GetID(), r)
		}
	}()
	
	if msg.GetID() < 20000 {
		return
	}
	msg.Encode()

	length := len(msg.GetStream().GetBuffer())
	fmt.Printf("Sending packet %d (size=%d)\n", msg.GetID(), length)

	header := make([]byte, 7)
	binary.BigEndian.PutUint16(header[0:2], uint16(msg.GetID()))
	header[2] = byte(length >> 16)
	header[3] = byte(length >> 8)
	header[4] = byte(length)
	binary.BigEndian.PutUint16(header[5:7], uint16(msg.GetVersion()))

	trailer := []byte{0xFF, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00}
	finalBuffer := append(header, msg.GetStream().GetBuffer()...)
	finalBuffer = append(finalBuffer, trailer...)

	msg.GetSession().Write(finalBuffer)
}
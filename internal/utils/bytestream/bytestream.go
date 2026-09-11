package bytestream

import (
	"encoding/binary"
)

type ByteStream struct {
	Buffer []byte
	Offset int
}

func New(data []byte) *ByteStream {
	if data == nil {
		data = make([]byte, 0)
	}
	return &ByteStream{
		Buffer: data,
		Offset: 0,
	}
}

// === Базовые типы ===

func (bs *ByteStream) ReadInt() int32 {
	if len(bs.Buffer) < bs.Offset+4 {
		return 0
	}
	val := int32(binary.BigEndian.Uint32(bs.Buffer[bs.Offset:bs.Offset+4]))
	bs.Offset += 4
	return val
}

func (bs *ByteStream) WriteInt(value int32) {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, uint32(value))
	bs.Buffer = append(bs.Buffer, bytes...)
	bs.Offset += 4
}

func (bs *ByteStream) ReadShort() int16 {
	if len(bs.Buffer) < bs.Offset+2 {
		return 0
	}
	val := int16(binary.BigEndian.Uint16(bs.Buffer[bs.Offset:bs.Offset+2]))
	bs.Offset += 2
	return val
}

func (bs *ByteStream) WriteShort(value int16) {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, uint16(value))
	bs.Buffer = append(bs.Buffer, bytes...)
	bs.Offset += 2
}

func (bs *ByteStream) ReadByte() byte {
	if bs.Offset >= len(bs.Buffer) {
		return 0
	}
	val := bs.Buffer[bs.Offset]
	bs.Offset++
	return val
}

func (bs *ByteStream) WriteByte(value byte) {
	bs.Buffer = append(bs.Buffer, value)
	bs.Offset++
}

// === Строки ===

func (bs *ByteStream) ReadString() string {
	length := bs.ReadInt()
	if length <= 0 || length > 90000 {
		return ""
	}
	if bs.Offset+int(length) > len(bs.Buffer) {
		return ""
	}
	str := string(bs.Buffer[bs.Offset : bs.Offset+int(length)])
	bs.Offset += int(length)
	return str
}

func (bs *ByteStream) WriteString(value string) {
	if value == "" || len(value) > 90000 {
		bs.WriteInt(-1)
		return
	}
	bytes := []byte(value)
	bs.WriteInt(int32(len(bytes)))
	bs.Buffer = append(bs.Buffer, bytes...)
	bs.Offset += len(bytes)
}

// === VInt ===

func (bs *ByteStream) WriteVInt(value int32) {
	zigzag := (value << 1) ^ (value >> 31)
	for {
		b := byte(zigzag & 0x7F)
		zigzag >>= 7
		if zigzag != 0 {
			b |= 0x80
		}
		bs.WriteByte(b)
		if zigzag == 0 {
			break
		}
	}
}

func (bs *ByteStream) ReadVInt() int32 {
	var result int32
	var shift uint
	for {
		b := bs.ReadByte()
		result |= int32(b&0x7F) << shift
		shift += 7
		if (b & 0x80) == 0 {
			break
		}
	}
	return (result >> 1) ^ (-(result & 1))
}

func (bs *ByteStream) WriteLong(value1, value2 int32) {
	bs.WriteInt(value1)
	bs.WriteInt(value2)
}

// ReadLong - читает Long как два int32
func (bs *ByteStream) ReadLong() (int32, int32) {
	return bs.ReadInt(), bs.ReadInt()
}

// === Boolean ===

func (bs *ByteStream) WriteBoolean(value bool) {
	if value {
		bs.WriteVInt(1)
	} else {
		bs.WriteVInt(0)
	}
}

func (bs *ByteStream) ReadBoolean() bool {
	return bs.ReadVInt() != 0
}

// === DataReference ===

func (bs *ByteStream) WriteDataReference(csvID, refID int32) {
	if csvID < 1 {
		bs.WriteVInt(0)
	} else {
		bs.WriteVInt(csvID)
		bs.WriteVInt(refID)
	}
}

func (bs *ByteStream) ReadDataReference() (int32, int32) {
	csvID := bs.ReadVInt()
	if csvID == 0 {
		return 0, 0
	}
	refID := bs.ReadVInt()
	return csvID, refID
}

// === Утилиты ===

func (bs *ByteStream) WriteBytes(data []byte) {
	if data == nil {
		bs.WriteInt(-1)
		return
	}
	bs.WriteInt(int32(len(data)))
	bs.Buffer = append(bs.Buffer, data...)
	bs.Offset += len(data)
}

func (bs *ByteStream) GetBuffer() []byte {
	return bs.Buffer
}

func (bs *ByteStream) Reset() {
	bs.Offset = 0
}

func (bs *ByteStream) Skip(n int) {
	bs.Offset += n
}

func (bs *ByteStream) Remaining() int {
	return len(bs.Buffer) - bs.Offset
}
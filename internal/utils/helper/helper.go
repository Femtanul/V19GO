package helper

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomString - генерирует случайную строку
func GenerateRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		// Fallback
		return base64.URLEncoding.EncodeToString(bytes)[:length]
	}
	return base64.URLEncoding.EncodeToString(bytes)[:length]
}

// GenerateToken - генерирует токен для аккаунта (40 символов)
func GenerateToken() string {
	return GenerateRandomString(40)
}

// IntToBytes - конвертирует int32 в 4 байта (Big Endian)
func IntToBytes(value int32) []byte {
	return []byte{
		byte(value >> 24),
		byte(value >> 16),
		byte(value >> 8),
		byte(value),
	}
}

// BytesToInt - конвертирует байты в int32
func BytesToInt(bytes []byte) int32 {
	if len(bytes) == 0 {
		return 0
	}
	if len(bytes) == 1 {
		return int32(bytes[0])
	}
	if len(bytes) == 2 {
		return int32(bytes[0])<<8 | int32(bytes[1])
	}
	if len(bytes) == 3 {
		return int32(bytes[0])<<16 | int32(bytes[1])<<8 | int32(bytes[2])
	}
	// 4+ байт
	return int32(bytes[0])<<24 | int32(bytes[1])<<16 | int32(bytes[2])<<8 | int32(bytes[3])
}
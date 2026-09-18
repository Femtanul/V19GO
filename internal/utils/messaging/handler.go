package messaging

import (
	"fmt"
	"v19go/internal/protocol/factory"
	"v19go/internal/protocol/piranha"
)

// Handler - обработчик входящих сообщений
type Handler struct {
	factory *factory.Factory
}

// NewHandler создает новый обработчик
func NewHandler(f *factory.Factory) *Handler {
	return &Handler{factory: f}
}

// Handle принимает ID пакета, байты и сессию, обрабатывает пакет
func (h *Handler) Handle(id int32, bytes []byte, session piranha.SessionInterface) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC in Handle (packet %d): %v\n", id, r)
		}
	}()

	if !h.factory.Exists(id) {
		fmt.Printf("Unknown packet: %d\n", id)
		return
	}

	constructor := h.factory.Get(id)
	msg := constructor(session, bytes)
	if msg == nil {
		return
	}

	msg.Decode()
	if err := msg.Process(); err != nil {
		fmt.Printf("Error processing packet %d: %v\n", id, err)
	}
}
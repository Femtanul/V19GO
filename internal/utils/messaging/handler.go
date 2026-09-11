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
	// 1. Проверяем, есть ли пакет в фабрике
	if !h.factory.Exists(id) {
		fmt.Printf("Unknown packet: %d\n", id)
		return
	}

	// 2. Получаем конструктор
	constructor := h.factory.Get(id)
	if constructor == nil {
		fmt.Printf("No constructor for packet: %d\n", id)
		return
	}

	// 3. Создаем сообщение
	// session имеет тип piranha.SessionInterface,
	// а MessageConstructor ожидает interface{} — это совместимо!
	msg := constructor(session, bytes)
	if msg == nil {
		fmt.Printf("Constructor returned nil for packet: %d\n", id)
		return
	}

	// 4. Декодируем
	msg.Decode()

	// 5. Обрабатываем
	if err := msg.Process(); err != nil {
		fmt.Printf("Error processing packet %d: %v\n", id, err)
	}
}
package factory

import (
	"fmt"
	"sync"
)

// MessageConstructor - функция-конструктор для сообщения
// Принимает сессию и байты, возвращает экземпляр сообщения
type MessageConstructor func(session interface{}, bytes []byte) Message

// Message - интерфейс, которому должны соответствовать все сообщения
// (в дальнейшем расширим)
type Message interface {
	GetID() int32
	Decode()
	Process() error
}

// Factory - фабрика сообщений
type Factory struct {
	mu       sync.RWMutex
	messages map[int32]MessageConstructor
}

// NewFactory создает пустую фабрику
func NewFactory() *Factory {
	return &Factory{
		messages: make(map[int32]MessageConstructor),
	}
}

// Register - регистрирует конструктор сообщения в фабрике
// Вызывается при инициализации каждого сообщения
func (f *Factory) Register(id int32, constructor MessageConstructor) {
	f.mu.Lock()
	defer f.mu.Unlock()
	f.messages[id] = constructor
}

// Get - возвращает конструктор сообщения по ID
// Если сообщения нет — возвращает nil
func (f *Factory) Get(id int32) MessageConstructor {
	f.mu.RLock()
	defer f.mu.RUnlock()
	return f.messages[id]
}

// Exists - проверяет, зарегистрировано ли сообщение с таким ID
func (f *Factory) Exists(id int32) bool {
	f.mu.RLock()
	defer f.mu.RUnlock()
	_, ok := f.messages[id]
	return ok
}

// ListIDs - возвращает список всех зарегистрированных ID
func (f *Factory) ListIDs() []int32 {
	f.mu.RLock()
	defer f.mu.RUnlock()
	ids := make([]int32, 0, len(f.messages))
	for id := range f.messages {
		ids = append(ids, id)
	}
	return ids
}

// Report - выводит отчет: какие пакеты есть, какие отсутствуют (по переданному списку)
func (f *Factory) Report(allKnownIDs []int32) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	fmt.Println("\n=== MessageFactory Report ===")
	fmt.Printf("Зарегистрировано пакетов: %d\n\n", len(f.messages))

	fmt.Println("✅ Зарегистрированные пакеты:")
	for id := range f.messages {
		fmt.Printf("   - %d\n", id)
	}

	fmt.Println("\n❌ Отсутствующие пакеты (из известных):")
	missingCount := 0
	for _, id := range allKnownIDs {
		if _, ok := f.messages[id]; !ok {
			fmt.Printf("   - %d\n", id)
			missingCount++
		}
	}
	fmt.Printf("\nВсего отсутствует: %d\n", missingCount)
	fmt.Println("================================\n")
}
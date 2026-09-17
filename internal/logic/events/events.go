package events

// Event - одно событие
type Event struct {
	Reward         int32
	SlotID         int32
	ID             int32
	Timer          int32
	Ended          bool
	TokensClaimed  []int32
}

// Events - менеджер событий
type Events struct {
	events []*Event
}

// NewEvents создает менеджер с дефолтными событиями
func NewEvents() *Events {
	return &Events{
		events: []*Event{
			{Reward: 20, SlotID: 0, ID: 11, Timer: 3600, Ended: false, TokensClaimed: []int32{}},
			{Reward: 20, SlotID: 1, ID: 13, Timer: 3600, Ended: false, TokensClaimed: []int32{}},
			{Reward: 20, SlotID: 2, ID: 5, Timer: 3600, Ended: false, TokensClaimed: []int32{}},
			{Reward: 20, SlotID: 3, ID: 25, Timer: 3600, Ended: false, TokensClaimed: []int32{}},
		},
	}
}

// GetAll возвращает все события
func (e *Events) GetAll() []*Event {
	return e.events
}

// GetIDBySlotID возвращает ID события по слоту
func (e *Events) GetIDBySlotID(slotID int32) int32 {
	for _, ev := range e.events {
		if ev.SlotID == slotID {
			return ev.ID
		}
	}
	return 0
}

// AddToken добавляет игрока в tokensclaimed
func (e *Events) AddToken(slotID int32, lowID int32) *Event {
	for _, ev := range e.events {
		if ev.SlotID == slotID {
			ev.TokensClaimed = append(ev.TokensClaimed, lowID)
			return ev
		}
	}
	return nil
}
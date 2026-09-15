package home

import (
	"v19go/internal/database"
	"v19go/internal/protocol/piranha"
	"v19go/internal/utils/bytestream"
)


type OwnHomeDataMessage struct {
	piranha.PiranhaMessage
	Account *database.Account
}

func NewOwnHomeDataMessage(session piranha.SessionInterface, account *database.Account) *OwnHomeDataMessage {
	msg := &OwnHomeDataMessage{}
	msg.Session = session
	msg.ID = 24101
	msg.Version = 0
	msg.Stream = bytestream.New(nil)
	msg.Account = account
	return msg
}

func (m *OwnHomeDataMessage) Encode() {
	acc := m.Account

	// === БАЗОВЫЕ ДАННЫЕ ===
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)

	m.Stream.WriteVInt(acc.Trophies)
	m.Stream.WriteVInt(acc.Trophies)

	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(acc.TrophyRoadTier)

	m.Stream.WriteVInt(acc.Experience)

	m.Stream.WriteDataReference(28, acc.Thumbnail)
	m.Stream.WriteDataReference(43, acc.Namecolor)

	m.Stream.WriteVInt(0) // PlayedGamemodesArray

	m.Stream.WriteVInt(0) // SelectedSkins
	m.Stream.WriteVInt(0) // UnlockedSkins

	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)

	m.Stream.WriteBoolean(false)
	m.Stream.WriteVInt(1)
	m.Stream.WriteBoolean(true)

	m.Stream.WriteVInt(acc.TokensDoubler)
	m.Stream.WriteVInt(0)

	m.Stream.WriteByte(8)

	m.Stream.WriteBoolean(false)
	m.Stream.WriteBoolean(false)
	m.Stream.WriteBoolean(false)

	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)

	// === МАГАЗИН ===
	m.Stream.WriteVInt(0) // пока пустой магазин

	// === ТОКЕНЫ ===
	m.Stream.WriteVInt(0) // AdsArray
	m.Stream.WriteVInt(acc.BattleTokens)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(acc.Tickets)
	m.Stream.WriteVInt(0)

	m.Stream.WriteDataReference(16, acc.BrawlerID)

	m.Stream.WriteString("RU")
	m.Stream.WriteString(acc.AuthorCode)

	// === ANIMATION ===
	m.Stream.WriteVInt(3)
	m.Stream.WriteInt(3)
	m.Stream.WriteInt(0) // tokensAnim
	m.Stream.WriteInt(4)
	m.Stream.WriteInt(0) // trophiesAnim
	m.Stream.WriteInt(5)
	m.Stream.WriteInt(0) // bigTokensAnim

	m.Stream.WriteVInt(2019049)
	m.Stream.WriteVInt(100)
	m.Stream.WriteVInt(10)

	m.Stream.WriteVInt(30)
	m.Stream.WriteVInt(3)
	m.Stream.WriteVInt(80)
	m.Stream.WriteVInt(10)
	m.Stream.WriteVInt(50)
	m.Stream.WriteVInt(1000)
	m.Stream.WriteVInt(500)
	m.Stream.WriteVInt(50)
	m.Stream.WriteVInt(999900)

	m.Stream.WriteVInt(0) // Array

	// === СОБЫТИЯ (пока пусто) ===
	m.Stream.WriteVInt(0) // EventSlots
	m.Stream.WriteVInt(0) // Events
	m.Stream.WriteVInt(0) // Coming events

	// === ПРОКАЧКА ===
	m.Stream.WriteVInt(8)
	for _, i := range []int32{20, 35, 75, 140, 290, 480, 800, 1250} {
		m.Stream.WriteVInt(i)
	}

	m.Stream.WriteVInt(8)
	for _, i := range []int32{1, 2, 3, 4, 5, 10, 15, 20} {
		m.Stream.WriteVInt(i)
	}

	m.Stream.WriteVInt(3)
	for _, i := range []int32{10, 30, 80} {
		m.Stream.WriteVInt(i)
	}

	m.Stream.WriteVInt(3)
	for _, i := range []int32{6, 20, 60} {
		m.Stream.WriteVInt(i)
	}

	m.Stream.WriteVInt(4)
	for _, i := range []int32{20, 50, 140, 280} {
		m.Stream.WriteVInt(i)
	}

	m.Stream.WriteVInt(4)
	for _, i := range []int32{150, 400, 1200, 2600} {
		m.Stream.WriteVInt(i)
	}

	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(200)
	m.Stream.WriteVInt(20)

	m.Stream.WriteVInt(8640)
	m.Stream.WriteVInt(10)
	m.Stream.WriteVInt(5)

	m.Stream.WriteVInt(6)
	m.Stream.WriteVInt(50)
	m.Stream.WriteVInt(604800)

	m.Stream.WriteBoolean(true)

	m.Stream.WriteVInt(0) // array

	m.Stream.WriteVInt(1)
	m.Stream.WriteInt(1)
	m.Stream.WriteInt(4100001)

	m.Stream.WriteInt(0)
	m.Stream.WriteInt(1)

	m.Stream.WriteVInt(0) // NotificationFactory
	m.Stream.WriteBoolean(false)

	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)

	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(acc.LowID)

	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)

	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)

	// === ИМЯ ИГРОКА ===
	m.Stream.WriteString(acc.Name)
	m.Stream.WriteVInt(0)

	m.Stream.WriteString("")

	// === БОЙЦЫ ===
	m.Stream.WriteVInt(8)
	m.Stream.WriteVInt(int32(len(acc.Brawlers)) + 4)

	for _, b := range acc.Brawlers {
		m.Stream.WriteDataReference(23, b.CardID)
		if b.Unlocked {
			m.Stream.WriteVInt(1)
		} else {
			m.Stream.WriteVInt(0)
		}
	}

	// === РЕСУРСЫ ===
	m.Stream.WriteDataReference(5, 1)
	m.Stream.WriteVInt(acc.Box)

	m.Stream.WriteDataReference(5, 8)
	m.Stream.WriteVInt(acc.Gold)

	m.Stream.WriteDataReference(5, 9)
	m.Stream.WriteVInt(acc.BigBox)

	m.Stream.WriteDataReference(5, 10)
	m.Stream.WriteVInt(acc.Starpoints)

	// Трофеи бойцов
	m.Stream.WriteVInt(int32(len(acc.Brawlers)))
	for _, b := range acc.Brawlers {
		m.Stream.WriteDataReference(16, b.BrawlerID)
		m.Stream.WriteVInt(b.Trophies)
	}

	// Ещё раз трофеи (да, в протоколе дублируется)
	m.Stream.WriteVInt(int32(len(acc.Brawlers)))
	for _, b := range acc.Brawlers {
		m.Stream.WriteDataReference(16, b.BrawlerID)
		m.Stream.WriteVInt(b.Trophies)
	}

	m.Stream.WriteVInt(0) // UnknownArray

	// Очки бойцов
	m.Stream.WriteVInt(int32(len(acc.Brawlers)))
	for _, b := range acc.Brawlers {
		m.Stream.WriteDataReference(16, b.BrawlerID)
		m.Stream.WriteVInt(b.Points)
	}

	// Уровни бойцов
	m.Stream.WriteVInt(int32(len(acc.Brawlers)))
	for _, b := range acc.Brawlers {
		m.Stream.WriteDataReference(16, b.BrawlerID)
		m.Stream.WriteVInt(b.Level)
	}

	// Навыки бойцов
	m.Stream.WriteVInt(int32(len(acc.Brawlers)))
	for _, b := range acc.Brawlers {
		m.Stream.WriteDataReference(23, b.SkillID)
		if b.SkillID != 0 {
			m.Stream.WriteVInt(1)
		} else {
			m.Stream.WriteVInt(0)
		}
	}

	m.Stream.WriteVInt(0)

	// === ГЕМЫ И ПРОЧЕЕ ===
	m.Stream.WriteVInt(int32(acc.Gems))
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(1)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(0)
	m.Stream.WriteVInt(2)
	m.Stream.WriteVInt(1585502369)

	// === ОБНОВЛЯЕМ СЕССИЮ ===
	m.Session.SetLowID(acc.LowID)
}

func (m *OwnHomeDataMessage) Decode() {}

func (m *OwnHomeDataMessage) Process() error { return nil }
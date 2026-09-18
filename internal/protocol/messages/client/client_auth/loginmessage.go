package client_auth

import ( 
	"v19go/internal/protocol/messages/server/server_auth"
	"v19go/internal/protocol/piranha"
	"v19go/internal/utils/bytestream"
	"v19go/internal/utils/helper"
	"v19go/internal/utils/messaging"
	"v19go/internal/protocol/messages/server/home"
	"fmt"
)

// LoginMessage - входящий пакет логина (ID: 10101)
type LoginMessage struct {
	piranha.PiranhaMessage

	HighID int32
	LowID  int32
	Token  string
	Major  int32
	Minor  int32
	Build  int32
}

func NewLoginMessage(session piranha.SessionInterface, bytes []byte) *LoginMessage {
	msg := &LoginMessage{}
	msg.Session = session
	msg.ID = 10101
	msg.Version = 0
	msg.Stream = bytestream.New(bytes)
	return msg
}

func (m *LoginMessage) Decode() {
	m.HighID = m.Stream.ReadInt()
	m.LowID = m.Stream.ReadInt()
	m.Token = m.Stream.ReadString()
	m.Major = m.Stream.ReadInt()
	m.Minor = m.Stream.ReadInt()
	m.Build = m.Stream.ReadInt()
}

// Process - логика логина
func (m *LoginMessage) Process() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("PANIC in LoginMessage.Process: %v\n", r)
			err = fmt.Errorf("panic: %v", r)
		}
	}()

	db := m.Session.GetDB()

	if db == nil {
		return nil
	}

	// 2. Если токен пустой — генерируем новый
	token := m.Token
	if token == "" {
		token = helper.GenerateToken()
	}

	// 3. Ищем аккаунт в БД
	account, err := db.GetAccountByToken(token)
	if err != nil {
		return err
	}

	// 4. Если не найден — создаем
	if account == nil {
		account, err = db.CreateAccount(token, "GOV19")
		if err != nil {
			return err
		}
	}

	// 5. Сохраняем данные в сессию
	m.Session.SetAccount(account)

	// 6. Отправляем LoginOKMessage
	loginOK := server_auth.NewLoginOKMessage(m.Session, account)
	messaging.Send(loginOK)


	ohd := home.NewOwnHomeDataMessage(m.Session, account)
	messaging.Send(ohd)
	return nil
}


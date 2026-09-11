package main

import (
	"fmt"
	"net"
	"v19go/config"
	"v19go/internal/database"
	"v19go/internal/protocol/factory"
	"v19go/internal/protocol/messages/client/client_auth"
	"v19go/internal/protocol/piranha"
	"v19go/internal/utils/messaging"
)

// ==================== SESSION ====================

type Session struct {
	conn    net.Conn
	db      *database.DB
	LowID   int32
	Token   string
	account *database.Account
}

func (s *Session) Write(data []byte) (int, error)   { return s.conn.Write(data) }
func (s *Session) GetCrypto() interface{}           { return nil }
func (s *Session) GetDB() *database.DB              { return s.db }
func (s *Session) SetLowID(id int32)                { s.LowID = id }
func (s *Session) GetLowID() int32                  { return s.LowID }
func (s *Session) SetToken(token string)            { s.Token = token }
func (s *Session) GetToken() string                 { return s.Token }
func (s *Session) SetAccount(acc *database.Account) { s.account = acc }
func (s *Session) GetAccount() *database.Account    { return s.account }

// ==================== MAIN ====================

func main() {
	cfg := config.Default

	// Подключение к БД
	db, err := database.NewDB(database.Config{
		Host:     cfg.Database.Host,
		Port:     cfg.Database.Port,
		User:     cfg.Database.User,
		Password: cfg.Database.Password,
		DBName:   cfg.Database.DBName,
	})
	if err != nil {
		fmt.Printf("Failed to connect to DB: %v\n", err)
		return
	}
	defer db.Close()
	fmt.Println("Connected to PostgreSQL!")

	// Фабрика и обработчик
	f := factory.NewFactory()
	handler := messaging.NewHandler(f)

	// Регистрация пакетов
	f.Register(10100, func(session interface{}, bytes []byte) factory.Message {
		sess, ok := session.(piranha.SessionInterface)
		if !ok {
			return nil
		}
		return client_auth.NewClientHelloMessage(sess, bytes)
	})

	f.Register(10101, func(session interface{}, bytes []byte) factory.Message {
		sess, ok := session.(piranha.SessionInterface)
		if !ok {
			return nil
		}
		return client_auth.NewLoginMessage(sess, bytes)
	})

	f.Report(factory.KnownClientIDs)

	// TCP listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	defer listener.Close()

	fmt.Printf("%s started on port %d\n", cfg.ServerName, cfg.Port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept error: %v\n", err)
			continue
		}
		fmt.Printf("New connection from %s\n", conn.RemoteAddr())
		go handleConnection(conn, handler, db)  // ← передаем db
	}
}

// ==================== HANDLE CONNECTION ====================

func handleConnection(conn net.Conn, handler *messaging.Handler, db *database.DB) {
	defer conn.Close()

	// ← ВОТ ЗДЕСЬ создаем session с db
	session := &Session{conn: conn, db: db}

	for {
		header := make([]byte, 7)
		_, err := conn.Read(header)
		if err != nil {
			fmt.Printf("Connection closed: %v\n", err)
			return
		}

		id := int32(header[0])<<8 | int32(header[1])
		length := int32(header[2])<<16 | int32(header[3])<<8 | int32(header[4])

		data := make([]byte, length)
		if length > 0 {
			_, err = conn.Read(data)
			if err != nil {
				fmt.Printf("Read data error: %v\n", err)
				return
			}
		}

		fmt.Printf("Packet %d (len=%d) received\n", id, length)

		handler.Handle(id, data, session)
	}
}
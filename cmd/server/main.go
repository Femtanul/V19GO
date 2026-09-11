package main

import (
	"fmt"
	"net"
	"v19go/config"
	"v19go/internal/protocol/factory"
	"v19go/internal/protocol/messages/client/client_auth"
	"v19go/internal/protocol/piranha"
	"v19go/internal/utils/messaging"
)

// Session - обертка над TCP-соединением
type Session struct {
	conn net.Conn
}

func (s *Session) Write(data []byte) (int, error) {
	return s.conn.Write(data)
}

func (s *Session) GetCrypto() interface{} {
	return nil
}

func main() {
	cfg := config.Default

	f := factory.NewFactory()
	handler := messaging.NewHandler(f)

	// Регистрируем пакеты
	f.Register(10100, func(session interface{}, bytes []byte) factory.Message {
		sess, ok := session.(piranha.SessionInterface)
		if !ok {
			return nil
		}
		return client_auth.NewClientHelloMessage(sess, bytes)
	})

	f.Report(factory.KnownClientIDs)

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
		go handleConnection(conn, handler)
	}
}

func handleConnection(conn net.Conn, handler *messaging.Handler) {
	defer conn.Close()

	session := &Session{conn: conn}

	for {
		// Читаем 7 байт заголовка
		header := make([]byte, 7)
		_, err := conn.Read(header)
		if err != nil {
			fmt.Printf("Connection closed: %v\n", err)
			return
		}

		// Парсим заголовок
		id := int32(header[0])<<8 | int32(header[1])
		length := int32(header[2])<<16 | int32(header[3])<<8 | int32(header[4])

		// Читаем данные
		data := make([]byte, length)
		if length > 0 {
			_, err = conn.Read(data)
			if err != nil {
				fmt.Printf("Read data error: %v\n", err)
				return
			}
		}

		fmt.Printf("Packet %d (len=%d) received\n", id, length)

		// Handler сам вызовет Decode() и Process()
		// А Process() отправит ServerHello
		handler.Handle(id, data, session)
	}
}
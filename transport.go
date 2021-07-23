package transport

import (
	"../actions"
	"bufio"
	"log"
	"net"
)

func HandleServerConnection(conn net.Conn, i int) {
	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("Closing client %v connection returned error: %v\n", i, err)
		}
	}()

	// scan message
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		//получили сообщение
		msg := scanner.Text()
		var msgNew string
		switch {
		case string(msg) == "Help":
			msgNew = actions.Help(msg)
		case string(msg) == "Hi":
			msgNew = actions.Hi(msg)
		case string(msg) == "Bye":
			msgNew = actions.Bye(msg)
		default :
			msgNew = actions.Unknown()
		}
		if _, err := conn.Write([]byte(msgNew + "\n")); err != nil {
			log.Printf("Client %v has closed writer: %v\n", i, err)
			break
		}
		log.Printf("Client %v\n received: %v\n returned: %v\n", i, msg, msgNew)
	}
	log.Printf("Client %v disconnected...\n", i)
}

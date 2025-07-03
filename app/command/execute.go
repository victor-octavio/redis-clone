package command

import (
	"fmt"
	"net"
	"strings"
)

// CLI Commands
const (
	ECHO = "ECHO"
	PING = "PING"
)

func Execute(conn net.Conn, args []string) {
	if len(args) == 0 {
		fmt.Println("No arguments")
		return
	}

	switch strings.ToUpper(args[0]) {
	case PING:
		conn.Write([]byte("+PONG\r\n"))
	case ECHO:
		if len(args) > 1 {
			msg := args[1]
			conn.Write([]byte(fmt.Sprintf("$%d\r\n%s\r\n", len(msg), msg)))
		} else {
			conn.Write([]byte("$0\r\n"))
		}
	default:
		conn.Write([]byte(fmt.Sprintf("-ERR unknown command '%s'\r\n", args[0])))
	}
}

package command

import (
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/database"
	"net"
	"strings"
)

// CLI Commands
const (
	ECHO = "ECHO"
	PING = "PING"
	SET  = "SET"
	GET  = "GET"
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
	case SET:
		key := args[1]
		val := args[2]
		database.SyncMap.Store(key, val)
		conn.Write([]byte("+OK\r\n"))
	case GET:
		key := args[1]
		value, ok := database.SyncMap.Load(key)
		if ok {
			res := fmt.Sprintf("$%d\r\n%s\r\n", len(value.(string)), value.(string))
			conn.Write([]byte(res))
		} else {
			conn.Write([]byte("$-1\r\n"))
		}
	default:
		conn.Write([]byte(fmt.Sprintf("-ERR unknown command '%s'\r\n", args[0])))
	}
}

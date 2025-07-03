package connection

import (
	"bufio"
	"fmt"
	"github.com/codecrafters-io/redis-starter-go/app/command"
	"net"
)

func Handle(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		cmd, err := command.ParseRESPArray(reader)
		if err != nil {
			fmt.Println(err)
			return
		}
		command.Execute(conn, cmd)
	}
}

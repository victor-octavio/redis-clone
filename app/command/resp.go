package command

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// RESP kind
const (
	Integer = ":"
	String  = "+"
	Bulk    = "$"
	Array   = "*"
	Error   = "-"
)

// TODO : *2\r\n$5\r\nhello\r\n$5\r\nworld\r\n --> RESP input example

func ParseRESPArray(reader *bufio.Reader) ([]string, error) {
	line, err := reader.ReadString('\n')
	if err != nil || !strings.HasPrefix(line, Array) {
		return nil, fmt.Errorf("Invalid array line: %s", line)
	}

	numArgs, err := strconv.Atoi(strings.TrimSpace(line[1:]))
	if err != nil {
		return nil, fmt.Errorf("Invalid array number: %s", line)
	}

	args := make([]string, 0, numArgs)

	for i := 0; i < numArgs; i++ {
		lenLine, _ := reader.ReadString('\n')

		strLen, err := strconv.Atoi(strings.TrimSpace(lenLine[1:]))
		if err != nil {
			return nil, fmt.Errorf("Invalid array length: %s", lenLine)
		}

		data := make([]byte, strLen+2)

		_, err = reader.Read(data)
		if err != nil {
			return nil, fmt.Errorf("Invalid array argument: %s", lenLine)
		}

		argStr := string(data[:strLen])

		args = append(args, argStr)
	}
	return args, nil
}

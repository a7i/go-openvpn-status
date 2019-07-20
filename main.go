package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ParseFile parses OpenVPN Status file ad returns a Status struct
func ParseFile(file string) (*Status, error) {
	conn, err := os.Open(file)
	if err != nil {
		return &Status{IsUp: false}, err
	}
	defer conn.Close()

	reader := bufio.NewReader(conn)
	return Parse(reader)
}

func ParseStr(content string) (*Status, error) {
	reader := bufio.NewReader(strings.NewReader(content))
	return Parse(reader)
}

func Parse(reader *bufio.Reader) (*Status, error) {
	if reader == nil {
		return nil, fmt.Errorf("reader cannot be nil")
	}
	if reader.Size() == 0 {
		return nil, fmt.Errorf("there is nothing to process. buffer size is zero")
	}

	return process(reader)
}

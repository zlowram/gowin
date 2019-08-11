package gowin

import "io"

func NewCString(reader io.Reader) string {
	readString := make([]byte, 0)
	b := make([]byte, 1)
	for {
		_, err := reader.Read(b)
		if err != nil || b[0] == 0 {
			break
		}
		readString = append(readString, b[0])
	}
	return string(readString)
}

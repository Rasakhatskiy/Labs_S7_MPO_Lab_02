package utils

import (
	"bufio"
	"io"
)

const (
	defaultBufSize = 4096
)

// MyReader implements bufio.Reader but returns value without delim on ReadString
type MyReader struct {
	reader *bufio.Reader
}

// NewReader overrides bufio.NewReader
func NewReader(rd io.Reader) *MyReader {
	return &MyReader{reader: bufio.NewReaderSize(rd, defaultBufSize)}
}

// ReadString reads until the first occurrence of delim in the input,
// returning a string containing the data up to but without the delimiter.
// If ReadString encounters an error before finding a delimiter,
// it returns the data read before the error and the error itself (often io.EOF).
// ReadString returns err != nil if and only if the returned data does not end in
// delim.
func (b *MyReader) ReadString(delim byte) (string, error) {
	result, err := b.reader.ReadString(delim)
	if err != nil {
		return result, err
	}
	return result[:len(result)-1], nil
}

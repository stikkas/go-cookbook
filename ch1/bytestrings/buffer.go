package bytestrings

import (
	"bytes"
	"io"
)

func Buffer(rawString string) *bytes.Buffer {
	rawBytes := []byte(rawString)
	// one way
	var b = new(bytes.Buffer)
	b.Write(rawBytes)
	// another way
	b = bytes.NewBuffer(rawBytes)
	// third way
	b = bytes.NewBufferString(rawString)
	return b
}

func toString(r io.Reader) (string, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

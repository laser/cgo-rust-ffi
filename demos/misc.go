package demos

import (
	"crypto/rand"
	"io"
)

func randomBytes(len int) []byte {
	slice := make([]byte, len)

	_, err := io.ReadFull(rand.Reader, slice)
	if err != nil {
		panic(err)
	}

	return slice
}

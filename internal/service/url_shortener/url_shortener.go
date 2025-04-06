package urlshortener

import (
	"crypto/rand"
	"strings"
)

const (
	length = 10
	alf    = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_"
)

func GenerateRandomAlias() string {
	b := make([]byte, length)
	rand.Read(b)

	builder := strings.Builder{}

	for i := 0; i < length; i++ {
		indx := int(b[i]) % len(alf)
		builder.WriteByte(alf[indx])
	}

	return builder.String()
}

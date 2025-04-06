package urlshortener

import (
	"testing"
)

func TestUniqueResults(t *testing.T) {
	res1 := GenerateRandomAlias()

	count := 0

	for i := 0; i < 10; i++ {
		if GenerateRandomAlias() == res1 {
			count++
		}
	}

	if count == 10 {
		t.Error("too many collisions")
	}
}

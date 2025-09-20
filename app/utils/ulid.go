package utils

import (
	"math/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

func GenerateULID() string {
	t := time.Now()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

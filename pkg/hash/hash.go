package hash

import (
	"github.com/Najwan160/go-experiment-3/code/domain/base"
	"golang.org/x/crypto/bcrypt"
)

type Hasher struct {
}

func NewHasher() base.Hasher {
	return &Hasher{}
}

func (*Hasher) Hash(b []byte, cost int) ([]byte, error) {
	return bcrypt.GenerateFromPassword(b, cost)
}

func (*Hasher) Verify(plain, hashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

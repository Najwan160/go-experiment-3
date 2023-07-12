package base

type Hasher interface {
	Hash(b []byte, cost int) ([]byte, error)
	Verify(plain, hashed string) bool
}

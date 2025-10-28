package bcrypt

import (
	"github.com/RealBirdMan91/blog/internal/modules/iam/application/ports"
	bc "golang.org/x/crypto/bcrypt"
)

const defaultCost = 12

type Hasher struct{}

func New() *Hasher { return &Hasher{} }

func (h *Hasher) Hash(plaintext string) (string, error) {
	bytes, err := bc.GenerateFromPassword([]byte(plaintext), defaultCost)
	return string(bytes), err
}
func (h *Hasher) Verify(hash, plaintext string) bool {
	return bc.CompareHashAndPassword([]byte(hash), []byte(plaintext)) == nil
}

// Compile-time check (optional, aber nett):
var _ ports.Hasher = (*Hasher)(nil)

package bcrypt

import (
	domain "github.com/RealBirdMan91/blog/internal/domain/user"
	bc "golang.org/x/crypto/bcrypt"
)

const defaultCost = 12 

type Hasher struct{}

func New() *Hasher { return &Hasher{} }

func (h *Hasher) Hash(plaintext string) (domain.PasswordHash, error) {
	bytes, err := bc.GenerateFromPassword([]byte(plaintext), defaultCost)
	if err != nil {
		return domain.PasswordHash{}, err
	}
	return domain.NewPasswordHash(string(bytes))
}

func (h *Hasher) Verify(hash domain.PasswordHash, plaintext string) bool {
	return bc.CompareHashAndPassword([]byte(hash.Hash()), []byte(plaintext)) == nil
}

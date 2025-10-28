package ports

// Hasher: App-Usecases brauchen "Hash" und optional "Verify".
type Hasher interface {
	Hash(plaintext string) (string, error)
	Verify(hash, plaintext string) bool
}

// TokenClaims: rein technisch (keine Domainlogik).
type TokenClaims struct {
	UserID  string
	Email   string
	ExpUnix int64
}

// TokenSigner/Verifier: App will Tokens signieren & prüfen.
type TokenSigner interface {
	Sign(userID, email string) (string, error)
}
type TokenVerifier interface {
	Verify(token string) (*TokenClaims, error)
}

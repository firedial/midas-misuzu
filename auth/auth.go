package auth

import (
    "crypto/rand"
    "math/big"
    "crypto/sha256"
    "encoding/hex"
)

func GetSalt() string {
    salt := ""

    for i := 0; i < 8; i++ {
        n, _ := rand.Int(rand.Reader, big.NewInt(4294967295))
        salt += n.Text(16)
    }
    return salt
}

func GetAuthToken(salt string, pass string) string {
    authToken := salt + pass

    for i := 0; i < 10; i++ {
        hash := sha256.Sum256([]byte(salt + authToken))
        authToken = hex.EncodeToString(hash[:])
    }

    return authToken
}


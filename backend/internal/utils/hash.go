package utils

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"os"
	"strings"
)

const (
	argonTime    uint32 = 3         // iterations
	argonMemory  uint32 = 64 * 1024 // 64 MiB
	argonThreads uint8  = 1
	keyLen       uint32 = 32
	saltLen             = 16
)

// HashPassword returns a PHC-encoded Argon2id hash string.
// pepper is your server-side secret (e.g. []byte(os.Getenv("PASSWORD_PEPPER"))).
func HashPassword(password string) (string, error) {
	pepper := []byte(os.Getenv("PASSWORD_PEPPER"))
	if len(pepper) == 0 {
		return "", errors.New("pepper (secret key) is required")
	}

	// 1) Derive a peppered password via HMAC(pepper, password)
	h := hmac.New(sha256.New, pepper)
	h.Write([]byte(password))
	peppered := h.Sum(nil)

	// 2) Generate a unique random salt to store with the hash
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("salt gen: %w", err)
	}

	// 3) Argon2id KDF
	hash := argon2.IDKey(peppered, salt, argonTime, argonMemory, argonThreads, keyLen)

	// 4) PHC string: $argon2id$v=19$m=...,t=...,p=...$base64(salt)$base64(hash)
	phc := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		argonMemory, argonTime, argonThreads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)
	return phc, nil
}

// VerifyPassword checks a password against a stored PHC string using the same pepper.
func VerifyPassword(storedPHC, password string) (bool, error) {
	pepper := []byte(os.Getenv("PASSWORD_PEPPER"))
	if !strings.HasPrefix(storedPHC, "$argon2id$") {
		return false, errors.New("unsupported hash format")
	}
	parts := strings.Split(storedPHC, "$")
	// ["", "argon2id", "v=19", "m=...,t=...,p=...", "saltB64", "hashB64"]
	if len(parts) != 6 {
		return false, errors.New("malformed PHC string")
	}

	// Parse params
	var m uint32
	var t uint32
	var p uint8
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &m, &t, &p); err != nil {
		return false, fmt.Errorf("parse params: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("salt decode: %w", err)
	}
	want, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("hash decode: %w", err)
	}

	// Recompute with pepper
	h := hmac.New(sha256.New, pepper)
	h.Write([]byte(password))
	peppered := h.Sum(nil)

	got := argon2.IDKey(peppered, salt, t, m, p, uint32(len(want)))

	// Constant-time compare
	if subtle.ConstantTimeCompare(got, want) == 1 {
		return true, nil
	}
	return false, nil
}

func HashMagicToken(token string) (string, error) {
	pepper := []byte(os.Getenv("MAGIC_TOKEN_PEPPER"))
	if len(pepper) == 0 {
		return "", errors.New("pepper (secret key) is required")
	}

	// 1) Derive a peppered password via HMAC(pepper, password)
	h := hmac.New(sha256.New, pepper)
	h.Write([]byte(token))
	peppered := h.Sum(nil)

	// 2) Generate a unique random salt to store with the hash
	salt := make([]byte, saltLen)
	if _, err := rand.Read(salt); err != nil {
		return "", fmt.Errorf("salt gen: %w", err)
	}

	// 3) Argon2id KDF
	hash := argon2.IDKey(peppered, salt, argonTime, argonMemory, argonThreads, keyLen)

	// 4) PHC string: $argon2id$v=19$m=...,t=...,p=...$base64(salt)$base64(hash)
	phc := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		argonMemory, argonTime, argonThreads,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)
	return phc, nil
}

// VerifyPassword checks a password against a stored PHC string using the same pepper.
func VerifyMagicToken(storedPHC, token string) (bool, error) {
	pepper := []byte(os.Getenv("PASSWORD_PEPPER"))
	if !strings.HasPrefix(storedPHC, "$argon2id$") {
		return false, errors.New("unsupported hash format")
	}
	parts := strings.Split(storedPHC, "$")
	// ["", "argon2id", "v=19", "m=...,t=...,p=...", "saltB64", "hashB64"]
	if len(parts) != 6 {
		return false, errors.New("malformed PHC string")
	}

	// Parse params
	var m uint32
	var t uint32
	var p uint8
	if _, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &m, &t, &p); err != nil {
		return false, fmt.Errorf("parse params: %w", err)
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, fmt.Errorf("salt decode: %w", err)
	}
	want, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, fmt.Errorf("hash decode: %w", err)
	}

	// Recompute with pepper
	h := hmac.New(sha256.New, pepper)
	h.Write([]byte(token))
	peppered := h.Sum(nil)

	got := argon2.IDKey(peppered, salt, t, m, p, uint32(len(want)))

	// Constant-time compare
	if subtle.ConstantTimeCompare(got, want) == 1 {
		return true, nil
	}
	return false, nil
}

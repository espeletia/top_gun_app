package hash

import (
	"FenceLive/internal/config"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

type HashUsecase struct {
	Salt string
	Hash hash.Hash
}

func NewHashUsecase(config config.HashConfig) HashUsecase {
	hash := sha512.New()
	return HashUsecase{Salt: config.Salt, Hash: hash}
}

func (hu HashUsecase) HashPassword(password string) string {
	passordBytes := []byte(password)
	saltBytes := []byte(hu.Salt)
	passordBytes = append(passordBytes, saltBytes...)
	hu.Hash.Write(passordBytes)
	hashedPasswordBytes := hu.Hash.Sum(nil)
	hashedPasswordHex := hex.EncodeToString(hashedPasswordBytes)
	return hashedPasswordHex
}

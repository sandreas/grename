package crypt

import (
	"encoding/hex"
	"hash"
	"io"
	"os"
)

func HashFileString(fileToHash string, h hash.Hash) (string, error) {
	f, err := os.Open(fileToHash)
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return hex.EncodeToString(h.Sum(nil)), nil
}


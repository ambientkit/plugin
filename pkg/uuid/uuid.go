// Package uuid generates a UUID.
package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

// Generate a UUID for use as an ID.
func Generate() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

// RandomString generates a random string.
// Source: https://devpy.wordpress.com/2013/10/24/create-random-string-in-golang/
func RandomString(length int) string {
	alphanum := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var bytes = make([]byte, length)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphanum[b%byte(len(alphanum))]
	}

	return string(bytes)
}

// EncodedString returns an encoded key.
func EncodedString(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	// Encode key in bytes to string for saving.
	return hex.EncodeToString(bytes)
}

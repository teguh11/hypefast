package helpers

import (
	"errors"
	"math/rand"
	"net/url"
	"time"
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Source of CIF randomness
var (
	src = rand.NewSource(time.Now().UnixNano())
)

func ShortenURL(URL string, length int) (shortValue string, err error) {
	isValid := isValidURL(URL)
	if !isValid {
		return "", errors.New("not valid URL")
	}
	b := make([]byte, length)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := length-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b), nil
}

func isValidURL(URL string) bool {
	_, err := url.ParseRequestURI(URL)
	if err != nil {
		return false
	}

	u, err := url.Parse(URL)
	if err != nil || u.Host == "" || u.Scheme == "" {
		return false
	}
	return true
}

package auth

import (
	"sync"
)

var (
	secretStr string
	once      sync.Once
)

func LoadSecret() {
	secretStr, _ = GetENV()
}

func GetSecret() string {
	once.Do(LoadSecret)
	return secretStr
}

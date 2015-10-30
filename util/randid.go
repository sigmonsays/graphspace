package util

import (
	"github.com/pborman/uuid"
)

const StringDictionary = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func RandomString() string {
	bytes := make([]byte, 16)
	u := uuid.NewUUID()
	for k, v := range u {
		bytes[k] = StringDictionary[v%byte(len(StringDictionary))]
	}
	return string(bytes)
}

package helpers

import (
	"github.com/lithammer/shortuuid/v4"
)

func GetUUID() string {
	return shortuuid.New()
}

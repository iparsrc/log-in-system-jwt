package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GetID() string {
	uuidWithHyphen := uuid.New()
	uuid := strings.Replace(uuidWithHyphen.String(), "-", "", -1)
	return uuid
}

package util

import (
	"strings"

	"github.com/google/uuid"
)

func GenID() string {
	u := uuid.New()
	return strings.ReplaceAll(u.String(),"-","")
}

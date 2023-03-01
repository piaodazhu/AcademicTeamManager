package common

import (
	"log"

	"github.com/gofrs/uuid"
)

func GenUUID() string {
	u, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	return u.String()
}
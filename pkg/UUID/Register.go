package UUID

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateUserID() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %v", err)
	}

	key := uuidObj.String()
	return key[:16], nil
}

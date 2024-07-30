package UUID

import (
	"fmt"

	"github.com/google/uuid"
)

func GenerateInterMediate() (string, error) {
	uuidObj, err := uuid.NewRandom()
	if err != nil {
		return "", fmt.Errorf("failed to generate UUID: %v", err)
	}

	key := uuidObj.String()
	return key[:21], nil
}

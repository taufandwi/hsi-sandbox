package etc

import "github.com/google/uuid"

func GenerateUUID() string {
	uuid, err := uuid.NewV7()
	if err != nil {
		// Handle the error appropriately, for example, log it or return an empty string
		return ""
	}
	return uuid.String()
}

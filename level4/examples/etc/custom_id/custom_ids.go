package custom_id

import "github.com/google/uuid"

func GetUUID() string {
	uuid, err := uuid.NewV7()
	if err != nil {
		return ""
	}
	return uuid.String()
}

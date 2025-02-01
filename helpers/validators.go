package helpers

import (
	"errors"
	"strings"
)

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return errors.New("invalid email format")
	}
	return nil
}

func ValidateTaskStatus(status string) error {
	validStatuses := []string{"pending", "in-progress", "completed"}
	for _, s := range validStatuses {
		if s == status {
			return nil
		}
	}
	return errors.New("invalid task status")
}

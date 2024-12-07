package user

import (
	"regexp"
	"time"
)

func (u *InputResponse) ValidateUserInput(user *InputResponse) bool {
	if user.Name == "" {
		return false
	}
	emailRegex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(user.Email) {
		return false
	}
	t, err := time.Parse("2006-01-02", user.RegistrationDate)
	if err != nil {
		return false
	}
	if t.Before(time.Now()) {
		return false
	}
	if user.Role == "" {
		return false
	}
	return true
}

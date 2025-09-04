package utils

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/radifan9/minitask-w10/internal/models"
)

var (
	reMin8           = regexp.MustCompile(`^.{8,}$`)
	reMinSmall       = regexp.MustCompile(`[a-z]`)
	reMinLarge       = regexp.MustCompile(`[A-Z]`)
	reMinSpecialChar = regexp.MustCompile(`[!@#$%^&*/()]`)
)

func ValidatePassword(user models.User) error {
	fmt.Println(user.Email)
	fmt.Println(user.Password)

	if !reMin8.MatchString(user.Password) {
		return errors.New("password harus minimal 8 karakter")
	}
	if !reMinSmall.MatchString(user.Password) {
		return errors.New("password minimal harus 1 karakter kecil")
	}
	if !reMinLarge.MatchString(user.Password) {
		return errors.New("password minimal harus 1 karakter besar")
	}
	if !reMinSpecialChar.MatchString(user.Password) {
		return errors.New("password harus ada karakter spesial")
	}
	return nil
}

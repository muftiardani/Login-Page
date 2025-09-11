package validator

import (
	"errors"
	"unicode"
)

// ValidatePassword memeriksa apakah kata sandi memenuhi kriteria keamanan.
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("kata sandi harus memiliki minimal 8 karakter")
	}

	var (
		hasUpper   bool
		hasLower   bool
		hasDigit   bool
		hasSpecial bool
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasUpper {
		return errors.New("kata sandi harus mengandung setidaknya satu huruf besar")
	}
	if !hasLower {
		return errors.New("kata sandi harus mengandung setidaknya satu huruf kecil")
	}
	if !hasDigit {
		return errors.New("kata sandi harus mengandung setidaknya satu angka")
	}
	if !hasSpecial {
		return errors.New("kata sandi harus mengandung setidaknya satu karakter spesial")
	}

	return nil
}
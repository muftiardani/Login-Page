package validator

import "testing"

func TestValidatePassword(t *testing.T) {
	testCases := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{"Password Kuat", "Arda123@", false},
		{"Tanpa Huruf Besar", "arda123@", true},
		{"Kurang dari 8 Karakter", "Arda1@", true},
		{"Tanpa Angka", "Arda@@@", true},
		{"Tanpa Simbol", "Arda123", true},
		{"Hanya 7 Karakter", "Arda12@", true},
		{"Tanpa Huruf Kecil", "ARDA123@", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidatePassword(tc.password)
			if (err != nil) != tc.wantErr {
				t.Errorf("ValidatePassword() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
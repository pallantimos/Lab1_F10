package main

import "testing"

func TestRegistratePass(t *testing.T) {
	testCases := []struct {
		login string
		pass1 string
		pass2 string
	}{
		{login: "Aldar", pass1: "123", pass2: "123"},
		{login: "Aldas", pass1: "987", pass2: "987"},
	}

	expectedError := "Пароль не содержит строчную букву"

	for _, tc := range testCases {
		error, result := checkRegistrate(tc.login, tc.pass1, tc.pass2)
		if error != expectedError {
			t.Errorf("Ожидалась строка '%s', получена строка '%s'", expectedError, error)
		}

		if result != false {
			t.Errorf("Ожидалось, что result будет равно false, но получено: %v", result)
		}

	}
}

func TestRegistrateLogin(t *testing.T) {
	error, result := checkRegistrate("Aldar", "Алдар123!", "Алдар123!")

	expectedError := "Логин уже существует"
	if error != expectedError {
		t.Errorf("Ожидалась строка '%s', получена строка '%s'", expectedError, error)
	}

	if result != false {
		t.Errorf("Ожидалось, что result будет равно false, но получено: %v", result)
	}
}

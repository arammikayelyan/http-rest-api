package model

import "testing"

// TestUser returns a user with valid data.
// This is needed so that we do not initialize user struct every time.
func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "password",
	}
}

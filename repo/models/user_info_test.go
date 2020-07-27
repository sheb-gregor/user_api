package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPasswordHash(t *testing.T) {
	password := "super-mega-secure-password"
	hash1 := PasswordHash(password)
	assert.True(t, ValidatePassword(password, hash1))
	assert.False(t, ValidatePassword("", hash1))
	assert.False(t, ValidatePassword("password", hash1))

	hash2 := PasswordHash("")
	assert.True(t, ValidatePassword("", hash2))
	assert.False(t, ValidatePassword(password, hash2))
}

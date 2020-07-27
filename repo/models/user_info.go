package models

import (
	"crypto/sha512"
	"encoding/hex"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// nolint:gochecknoglobals
var salt = "caafcf8601f9a6777002dab729ce40b9cd0b6b04c8edf6cd442078353809212e"

type UserInfo struct {
	ID           primitive.ObjectID `bson:"_id"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"password_hash"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
}

func NewUserInfo(email string, password string) *UserInfo {
	return &UserInfo{
		Email:        email,
		PasswordHash: PasswordHash(password),
		CreatedAt:    time.Now().UTC(),
		UpdatedAt:    time.Now().UTC(),
	}

}

func PasswordHash(password string) string {
	p := salt[:len(salt)/2] + password + salt[:len(salt)/2]
	sum := sha512.Sum512([]byte(p))
	return hex.EncodeToString(sum[:])

}

func ValidatePassword(password, hash string) bool {
	return hash == PasswordHash(password)
}

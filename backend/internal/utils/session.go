package utils

import (
	"backend/internal/models"
	"time"

	"github.com/google/uuid"
)

const (
	sessionTokenExpiry = 7 * 24 * time.Hour
)

func CreateSession(userId uuid.UUID) (models.Session, error) {
	newSession := models.Session{
		ID: uuid.New(),
		UserID: userId,
		ExpiresAt: time.Now().Add(sessionTokenExpiry),
		UpdatedAt: time.Now(),
		LastActive: time.Now(),
	}
	// Store session data in databaseconst
	return newSession, nil
}

func GetSession() {
	// Fetches session from database
	// returns user data if session exists and isn't expired
}

func RequireAuth() {
	// Verifies if the session exists and has permission
}

func RefreshSession() {
	// Extends sessions expirtion time
}

func Logout() {
	// Deletes session from database
}

func CleanupExpiredSessions() {
	// Deletes expired sessions from database
}
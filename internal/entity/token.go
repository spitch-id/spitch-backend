package entity

import "time"

type Token struct {
	ID           string     `db:"id"`
	UserID       string     `db:"user_id"`
	UserDeviceID *string    `db:"user_device_id"`
	TokenHash    string     `db:"token_hash"`
	ExpiresAt    time.Time  `db:"expires_at"`
	RevokedAt    *time.Time `db:"revoked_at"`
	CreatedAt    time.Time  `db:"created_at"`
}

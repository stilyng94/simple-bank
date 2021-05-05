package token

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID        uuid.UUID `json:"jti"`
	Username  string    `json:"username"`
	IssuedAt  int64     `json:"iat"`
	ExpiresAt int64     `json:"exp"`
}

func NewPayload(username string, duration time.Duration) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	timer := time.Now()
	issuedAt := timer.Unix()
	expiredAt := timer.Add(duration).Unix()
	payload := &Payload{ID: tokenID, Username: username, IssuedAt: issuedAt, ExpiresAt: expiredAt}
	return payload, nil
}

func (payload *Payload) Valid() error {
	now := time.Now().Unix()
	_ = time.Unix(now, 0).Sub(time.Unix(payload.ExpiresAt, 0))
	if now >= payload.ExpiresAt {
		return ErrExpiredToken
	}
	if payload.IssuedAt > now {
		return ErrInvalidToken
	}
	return nil

}

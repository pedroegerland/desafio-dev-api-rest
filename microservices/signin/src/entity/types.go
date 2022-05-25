package entity

import "time"

var (
	Users    = map[string]string{}
	Sessions = map[string]Session{}
)

type Session struct {
	Cpf    string
	Expiry time.Time
}

type Credentials struct {
	Password string `json:"password"`
	Cpf      string `json:"cpf"`
}

type Response struct {
	Msg          string `json:"message"`
	SessionToken string `json:"session_token"`
	Cpf          string `json:"cpf"`
	ExpiresAt    string `json:"expires_at"`
}

type Token struct {
	Cpf          string `json:"cpf"`
	SessionToken string `json:"session_token"`
	ExpiresAt    string `json:"expires_at"`
}

func (s Session) IsExpired() bool {
	return s.Expiry.Before(time.Now().UTC())
}

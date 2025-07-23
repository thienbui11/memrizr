package model

type TokenPair struct {
	IDTokenID    string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

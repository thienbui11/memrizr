<<<<<<< HEAD
package model

import "github.com/google/uuid"

// RefreshToken stores token properties that
// are accessed in multiple application layers
type RefreshToken struct {
	ID  uuid.UUID `json:"-"`
	UID uuid.UUID `json:"-"`
	SS  string    `json:"refreshToken"`
}

// IDToken stores token properties that
// are accessed in multiple application layers
type IDToken struct {
	SS string `json:"idToken"`
}

// TokenPair used for returning pairs of id and refresh tokens
type TokenPair struct {
	IDToken
	RefreshToken
}
=======
package model

import "github.com/google/uuid"

// RefreshToken stores token properties that
// are accessed in multiple application layers
type RefreshToken struct {
	ID  uuid.UUID `json:"-"`
	UID uuid.UUID `json:"-"`
	SS  string    `json:"refreshToken"`
}

// IDToken stores token properties that
// are accessed in multiple application layers
type IDToken struct {
	SS string `json:"idToken"`
}

// TokenPair used for returning pairs of id and refresh tokens
type TokenPair struct {
	IDToken
	RefreshToken
}
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643

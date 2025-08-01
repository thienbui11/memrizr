<<<<<<< HEAD
package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/thienbui11/MindStack/account/model"
)

// MockTokenService is a mock type for model.TokenService
type MockTokenService struct {
	mock.Mock
}

// NewPairFromUser mocks concrete NewPairFromUser
func (m *MockTokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	ret := m.Called(ctx, u, prevTokenID)

	// first value passed to "Return"
	var r0 *model.TokenPair
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.TokenPair)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// Signout mocks concrete Signout
func (m *MockTokenService) Signout(ctx context.Context, uid uuid.UUID) error {
	ret := m.Called(ctx, uid)
	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

// ValidateIDToken mocks concrete ValidateIDToken
func (m *MockTokenService) ValidateIDToken(tokenString string) (*model.User, error) {
	ret := m.Called(tokenString)

	// first value passed to "Return"
	var r0 *model.User
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// ValidateRefreshToken mocks concrete ValidateRefreshToken
func (m *MockTokenService) ValidateRefreshToken(refreshTokenString string) (*model.RefreshToken, error) {
	ret := m.Called(refreshTokenString)

	var r0 *model.RefreshToken
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.RefreshToken)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
=======
package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/jacobsngoodwin/memrizr/account/model"
	"github.com/stretchr/testify/mock"
)

// MockTokenService is a mock type for model.TokenService
type MockTokenService struct {
	mock.Mock
}

// NewPairFromUser mocks concrete NewPairFromUser
func (m *MockTokenService) NewPairFromUser(ctx context.Context, u *model.User, prevTokenID string) (*model.TokenPair, error) {
	ret := m.Called(ctx, u, prevTokenID)

	// first value passed to "Return"
	var r0 *model.TokenPair
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.TokenPair)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// Signout mocks concrete Signout
func (m *MockTokenService) Signout(ctx context.Context, uid uuid.UUID) error {
	ret := m.Called(ctx, uid)
	var r0 error

	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}

	return r0
}

// ValidateIDToken mocks concrete ValidateIDToken
func (m *MockTokenService) ValidateIDToken(tokenString string) (*model.User, error) {
	ret := m.Called(tokenString)

	// first value passed to "Return"
	var r0 *model.User
	if ret.Get(0) != nil {
		// we can just return this if we know we won't be passing function to "Return"
		r0 = ret.Get(0).(*model.User)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}

// ValidateRefreshToken mocks concrete ValidateRefreshToken
func (m *MockTokenService) ValidateRefreshToken(refreshTokenString string) (*model.RefreshToken, error) {
	ret := m.Called(refreshTokenString)

	var r0 *model.RefreshToken
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(*model.RefreshToken)
	}

	var r1 error

	if ret.Get(1) != nil {
		r1 = ret.Get(1).(error)
	}

	return r0, r1
}
>>>>>>> 3999c5b64d9e82200d4e850d267e0d3ed56f0643
